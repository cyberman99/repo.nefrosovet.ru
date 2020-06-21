package esia

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/fullsailor/pkcs7"
	"github.com/google/uuid"
	"golang.org/x/net/context/ctxhttp"
)

var certExp = regexp.MustCompile(`(?m)\-+?BEGIN.+\-+?$[\w\W]+?\-+?END.+\-+?$`)

type ExchangeResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`

	AuthToken struct {
		Nbf          string `json:"-"`
		Scope        string `json:"scope"`
		Iss          string `json:"iss"`
		UrnEsiaSid   string `json:"urn:esia:sid"`
		UrnEsiaSbjId int32  `json:"urn:esia:sbj_id"`
		Iat          int32  `json:"iat"`
		ClientId     string `json:"client_id"`
		Exp          int32  `json:"exp"`
	} `json:"auth_code"`

	State     string `json:"state"`
	TokenType string `json:"token_type"`
	ExpiresIn int64  `json:"expires_in"`
}

func Secret(clientID, clientSecret string, scopes []string, timestamp, state string) (string, error) {
	secretString := fmt.Sprintf("%s%s%s%s",
		strings.Join(scopes, " "),
		timestamp,
		clientID,
		state,
	)

	return Sign(clientSecret, []byte(secretString))
}

func State() string {
	return uuid.New().String()
}

func Timestamp() string {
	return time.Now().Format("2006.01.02 15:04:05 -0700")
}

func Sign(secret string, content []byte) (string, error) {
	ss := certExp.FindAllStringSubmatch(secret, -1)
	if len(ss) != 2 {
		return "", errors.New("can't parse esia secret")
	}

	// Parse private key
	pkBuffer := []byte(ss[1][0])
	pkBlock, _ := pem.Decode(pkBuffer)
	pkBufferParseResult, err := x509.ParsePKCS8PrivateKey(pkBlock.Bytes)
	if err != nil {
		return "", err
	}

	// Parse certificate
	certBuffer := []byte(ss[0][0])
	certBlock, _ := pem.Decode(certBuffer)
	certBufferParseResult, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return "", err
	}

	// Signing
	toBeSigned, err := pkcs7.NewSignedData(content)
	if err != nil {
		return "", err
	}

	if err := toBeSigned.AddSigner(certBufferParseResult, pkBufferParseResult, pkcs7.SignerInfoConfig{}); err != nil {
		return "", err
	}
	toBeSigned.Detach()

	signBuffer, err := toBeSigned.Finish()
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(signBuffer), nil
}

func Exchange(ctx context.Context, clientID, clientSecret string, scopes []string, code, redirectURI string) (*ExchangeResult, error) {
	timestamp := Timestamp()
	state := State()

	clientSecret, err := Secret(clientID, clientSecret, scopes, timestamp, state)
	if err != nil {
		return nil, err
	}

	v := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"timestamp":     {timestamp},
		"state":         {state},
		"redirect_uri":  {redirectURI},
		"token_type":    {"Bearer"},
		"scope":         {strings.Join(scopes, " ")},
	}

	req, err := http.NewRequest("POST", Endpoint.TokenURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := ctxhttp.Do(ctx, http.DefaultClient, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var exchangeResult ExchangeResult
	if err := json.NewDecoder(resp.Body).Decode(&exchangeResult); err != nil {
		return &exchangeResult, err
	}

	chunks := strings.Split(exchangeResult.AccessToken, ".")
	if len(chunks) < 2 {
		return nil, fmt.Errorf("can't get access token")
	}

	data, err := base64.URLEncoding.DecodeString(chunks[1])
	if err != nil {
		data, err = base64.URLEncoding.DecodeString(chunks[1] + "==")
		if err != nil {
			return &exchangeResult, err
		}
	}

	if err = json.Unmarshal(data, &exchangeResult.AuthToken); err != nil {
		return &exchangeResult, err
	}

	return &exchangeResult, nil
}

func GetInfo(ctx context.Context, userID int32, accessToken, path string, item interface{}) error {
	var infoURL string

	if userID == -1 {
		infoURL = path
	} else {
		infoURL = fmt.Sprintf("%srs/prns/%d%s", IssuerName, userID, path)
	}

	req, err := http.NewRequest("GET", infoURL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := ctxhttp.Do(ctx, http.DefaultClient, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return err
	}

	return nil
}
