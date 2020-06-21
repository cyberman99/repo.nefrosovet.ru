package esia

import (
	"golang.org/x/oauth2"
)

// https://esia.gosuslugi.ru/
// or test:
// https://esia-portal1.test.gosuslugi.ru/
const IssuerName = "https://esia.gosuslugi.ru/"

// Endpoint is ESIA OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:   IssuerName + "/aas/oauth2/ac",
	TokenURL:  IssuerName + "/aas/oauth2/te",
	AuthStyle: oauth2.AuthStyleInParams,
}

// All availables scopes:
// "birthdate", "gender", "snils", "inn", "id_doc", "birthplace", "medical_doc", "email",
//		"contacts", "kid_fullname", "kid_birthdate", "kid_gender", "kid_snils", "kid_inn", "kid_birth_cert_doc",
//		"kid_medical_doc"
