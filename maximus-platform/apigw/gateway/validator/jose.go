package validator

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/auth0-community/go-auth0"
	krakendjose "github.com/devopsfaith/krakend-jose"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	muxkrakend "github.com/devopsfaith/krakend/router/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v2/jwt"
	"repo.nefrosovet.ru/maximus-platform/apigw/influxdb"
)

const (
	authStatusPass = "PASS"
	authStatusFail = "FAIL"
)

func HandlerFactory(hf muxkrakend.HandlerFactory, paramExtractor muxkrakend.ParamExtractor, logger logging.Logger, rejecter krakendjose.Rejecter) muxkrakend.HandlerFactory {
	return TokenSigner(TokenSignatureValidator(hf, logger, rejecter), paramExtractor, logger)
}

func TokenSigner(hf muxkrakend.HandlerFactory, paramExtractor muxkrakend.ParamExtractor, logger logging.Logger) muxkrakend.HandlerFactory {
	return func(cfg *config.EndpointConfig, prxy proxy.Proxy) http.HandlerFunc {
		signerCfg, signer, err := krakendjose.NewSigner(cfg, nil)
		if err == krakendjose.ErrNoSignerCfg {
			logger.Debug("JOSE: singer disabled for the endpoint", cfg.Endpoint)
			return hf(cfg, prxy)
		}
		if err != nil {
			logger.Error(err.Error(), cfg.Endpoint)
			return hf(cfg, prxy)
		}

		logger.Debug("JOSE: singer enabled for the endpoint", cfg.Endpoint)

		return func(w http.ResponseWriter, r *http.Request) {
			proxyReq := muxkrakend.NewRequestBuilder(paramExtractor)(r, cfg.QueryString, cfg.HeadersToPass)
			ctx, cancel := context.WithTimeout(r.Context(), cfg.Timeout)
			defer cancel()

			response, err := prxy(ctx, proxyReq)
			if err != nil {
				logger.Error("proxy response error:", err.Error())
				http.Error(w, "", http.StatusBadRequest)
				return
			}

			if response == nil {
				http.Error(w, "", http.StatusBadRequest)
				return
			}

			if err := krakendjose.SignFields(signerCfg.KeysToSign, signer, response); err != nil {
				logger.Error(err.Error())
				http.Error(w, "", http.StatusBadRequest)
				return
			}

			for k, v := range response.Metadata.Headers {
				w.Header().Set(k, v[0])
			}

			err = jsonRender(w, response)
			if err != nil {
				logger.Error("render answer error:", err.Error())
			}
		}
	}
}

var emptyResponse = []byte("{}")

func jsonRender(w http.ResponseWriter, response *proxy.Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Metadata.StatusCode)

	if response == nil {
		_, err := w.Write(emptyResponse)
		return err
	}

	js, err := json.Marshal(response.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	_, err = w.Write(js)
	return err
}

func TokenSignatureValidator(hf muxkrakend.HandlerFactory, logger logging.Logger, rejecter krakendjose.Rejecter) muxkrakend.HandlerFactory {
	if rejecter == nil {
		rejecter = krakendjose.FixedRejecter(false)
	}
	return func(cfg *config.EndpointConfig, prxy proxy.Proxy) http.HandlerFunc {
		policyID := ""
		if cfg.ExtraConfig["policyID"] != nil {
			policyID = cfg.ExtraConfig["policyID"].(string)
		}
		handler := hf(cfg, prxy)
		signatureConfig, err := krakendjose.GetSignatureConfig(cfg)
		if err == krakendjose.ErrNoValidatorCfg {
			logger.Debug("JOSE: validator disabled for the endpoint", cfg.Endpoint)
			return handler
		}
		if err != nil {
			logger.Warning(fmt.Sprintf("JOSE: validator for %s: %s", cfg.Endpoint, err.Error()))
			return handler
		}

		validator, err := krakendjose.NewValidator(signatureConfig, FromCookie)
		if err != nil {
			logrus.Fatalf("%s: %s", cfg.Endpoint, err.Error())
		}

		logger.Debug("JOSE: validator enabled for the endpoint", cfg.Endpoint)

		return func(w http.ResponseWriter, r *http.Request) {
			token, err := validator.ValidateRequest(r)
			if err != nil {
				influxdb.LogEvent(influxdb.EventsParams{
					Status:   authStatusFail,
					IP:       getSourceIP(r),
					PolicyID: policyID,
					Endpoint: cfg.Endpoint,
					Method:   r.Method,
					Path:     r.RequestURI,
					Roles:    "",
				})

				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			claims := map[string]interface{}{}
			err = validator.Claims(r, token, &claims)
			if err != nil {
				influxdb.LogEvent(influxdb.EventsParams{
					Status:   authStatusFail,
					IP:       getSourceIP(r),
					PolicyID: policyID,
					Endpoint: cfg.Endpoint,
					Method:   r.Method,
					Path:     r.RequestURI,
					Roles:    "",
				})

				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			if rejecter.Reject(claims) {
				influxdb.LogEvent(influxdb.EventsParams{
					Status:   authStatusFail,
					IP:       getSourceIP(r),
					PolicyID: policyID,
					Endpoint: cfg.Endpoint,
					Method:   r.Method,
					Path:     r.RequestURI,
					Roles:    strings.Join(getRolesFromClaims(claims), " "),
				})

				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			if !krakendjose.CanAccess(signatureConfig.RolesKey, claims, signatureConfig.Roles) {
				influxdb.LogEvent(influxdb.EventsParams{
					Status:   authStatusFail,
					IP:       getSourceIP(r),
					PolicyID: policyID,
					Endpoint: cfg.Endpoint,
					Method:   r.Method,
					Path:     r.RequestURI,
					Roles:    strings.Join(getRolesFromClaims(claims), " "),
				})

				http.Error(w, "", http.StatusUnauthorized)
				return
			}

			influxdb.LogEvent(influxdb.EventsParams{
				Status:   authStatusPass,
				IP:       getSourceIP(r),
				PolicyID: policyID,
				Endpoint: cfg.Endpoint,
				Method:   r.Method,
				Path:     r.RequestURI,
				Roles:    strings.Join(getRolesFromClaims(claims), " "),
			})

			handler(w, r)
		}
	}
}

func FromCookie(key string) func(r *http.Request) (*jwt.JSONWebToken, error) {
	if key == "" {
		key = "access_token"
	}
	return func(r *http.Request) (*jwt.JSONWebToken, error) {
		cookie, err := r.Cookie(key)
		if err != nil {
			return nil, auth0.ErrTokenNotFound
		}
		return jwt.ParseSigned(cookie.Value)
	}
}

func getSourceIP(req *http.Request) string {
	var err error
	var addr string

	switch {
	case req.Header.Get("X-REAL-IP") != "":
		addr = req.Header.Get("X-REAL-IP")
	case req.Header.Get("X-FORWARDED-FOR") != "":
		addr = req.Header.Get("X-FORWARDED-FOR")
	default:
		addr, _, err = net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			logrus.WithField("address", req.RemoteAddr).Error("address is not IP:port")
		}
	}

	return addr
}

func getRolesFromClaims(claims map[string]interface{}) []string {
	var roles []string
	if claims["roles"] != nil {
		for _, role := range claims["roles"].([]interface{}) {
			roles = append(roles, role.(string))
		}
	}

	return roles
}
