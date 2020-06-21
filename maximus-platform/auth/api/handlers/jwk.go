package handlers

import (
	"encoding/base64"

	"github.com/go-openapi/runtime/middleware"
	"github.com/spf13/viper"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/j_w_k"
)

var signatureAlg = "HS256"

// GetJwk - GET /jwk
func GetJwk(_ j_w_k.GetJwkParams) middleware.Responder {
	secretKey := base64.RawURLEncoding.EncodeToString([]byte(viper.GetString("tokenSecret")))
	keyID := "default"
	keyType := "oct"

	payload := new(j_w_k.GetJwkOKBody)

	item := new(j_w_k.KeysItems0)
	item.Alg = &signatureAlg
	item.Kid = &keyID
	item.K = &secretKey
	item.Kty = &keyType

	payload.Keys = append(payload.Keys, item)

	return j_w_k.NewGetJwkOK().WithPayload(payload)
}
