package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/token"
)

// GetWhoami - GET /whoami
func GetWhoami(_ token.GetWhoamiParams, principal interface{}) middleware.Responder {
	claims := principal.(jwt.MapClaims)

	payload := new(token.GetWhoamiOKBody)
	payload.Version = &Version
	payload.Errors = make([]interface{}, 0)

	item := new(token.DataItems0)
	item.ID = claims["ID"].(string)

	if claims["roles"] != nil {
		for _, role := range claims["roles"].([]interface{}) {
			item.Roles = append(item.Roles, role.(string))
		}
	}

	payload.Data = append(payload.Data, item)
	payload.Message = &PayloadSuccessMessage
	return token.NewGetWhoamiOK().WithPayload(payload)
}
