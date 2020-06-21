package gateway

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/devopsfaith/krakend/encoding"

	"github.com/devopsfaith/krakend/config"
	"github.com/sirupsen/logrus"
	"repo.nefrosovet.ru/maximus-platform/apigw/mongodb"
)

type KrakendParams struct {
	Host      string
	Port      int
	EDGEProxy string
	LogLevel  string
	JWKURL    string
}

// New starts KrakenD instance and reloads it on every policies, rules and filters change.
func New(params KrakendParams, repo mongodb.PolicyRepository) {
	limiter := time.Tick(time.Second)

	for {
		revision, err := repo.GetLastPolicyTime()
		if err != nil && err != mongo.ErrNoDocuments {
			logrus.WithError(err).Debug()
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "configDB",
				"status":   "FAILED",
			}).Error("cant't get KrakenD config revision")

			<-limiter

			continue
		}

		ctx, cancel := context.WithCancel(context.Background())
		go Run(ctx, constructKrakenDConfig(params, repo))

		newRevision := revision
		for revision == newRevision {
			<-limiter

			newRevision, err = repo.GetLastPolicyTime()
			if err != nil && err != mongo.ErrNoDocuments {
				logrus.WithError(err).Debug()
				logrus.WithFields(logrus.Fields{
					"context":  "CORE",
					"resource": "configDB",
					"status":   "FAILED",
				}).Error("cant't get KrakenD config revision")

				continue
			}
		}

		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "Gateway",
			"status":   "RELOAD",
		}).Info("reloading gateway")

		cancel()
	}
}

func constructKrakenDConfig(params KrakendParams, repo mongodb.PolicyRepository) config.ServiceConfig {
	serviceConfig := config.ServiceConfig{}
	serviceConfig.Name = "API gateway"
	serviceConfig.Version = config.ConfigVersion
	serviceConfig.Port = params.Port
	serviceConfig.DialerTimeout = time.Minute
	if strings.ToLower(params.LogLevel) == "debug" {
		serviceConfig.Debug = true
	}

	policies, err := repo.GetPolicies()
	if err != nil && err != mongo.ErrNoDocuments {
		logrus.Fatal(err)
	}
	for _, policy := range policies {
		backend := new(config.Backend)
		backend.Method = policy.Method
		backend.Encoding = encoding.NOOP
		if params.EDGEProxy != "" {
			backend.Host = append(backend.Host, fmt.Sprintf("http://%s", params.EDGEProxy))
			backend.URLPattern = fmt.Sprintf("%s/%s", policy.Resource, policy.BackendPath)
		} else {
			backend.Host = append(backend.Host, fmt.Sprintf("http://%s", policy.BackendHost))
			backend.URLPattern = policy.BackendPath
		}

		endpoint := new(config.EndpointConfig)

		if params.EDGEProxy != "" {
			endpoint.Endpoint = fmt.Sprintf("%s/%s", policy.Resource, policy.Path)
		} else {
			endpoint.Endpoint = policy.Path
		}

		endpoint.Method = policy.Method
		endpoint.Backend = append(endpoint.Backend, backend)
		endpoint.OutputEncoding = encoding.NOOP
		endpoint.HeadersToPass = policy.HeadersToPass
		endpoint.QueryString = policy.QueryStringParams
		endpoint.Timeout = time.Minute

		endpoint.ExtraConfig = make(map[string]interface{})
		endpoint.ExtraConfig["policyID"] = policy.ID

		serviceConfig.Endpoints = append(serviceConfig.Endpoints, endpoint)

		if policy.Roles != nil && len(policy.Roles) != 0 {
			validator := make(map[string]interface{})
			validator["jwk-url"] = params.JWKURL
			validator["alg"] = "HS256"
			validator["roles_key"] = "roles"
			validator["roles"] = policy.Roles
			validator["disable_jwk_security"] = true
			if policy.KeyCache != 0 {
				validator["cache"] = true
				validator["key_cache"] = policy.KeyCache
			}

			endpoint.ExtraConfig["github.com/devopsfaith/krakend-jose/validator"] = validator
		}

		if policy.Cache {
			endpoint.ExtraConfig["github.com/devopsfaith/krakend-httpcache"] = make(map[string]interface{})
		}
	}

	err = serviceConfig.Init()
	if err != nil {
		log.Fatal(err)
	}

	//spew.Dump(serviceConfig)
	return serviceConfig
}
