package storage

import (
	"fmt"

	"repo.nefrosovet.ru/maximus-platform/auth/api/models"
)

var (
	ErrBackendNotFound = fmt.Errorf("backend: %w", ErrNotFound)
)

type BackendStorage interface {
	Store(in StoreBackend) (*Backend, error)
	Update(id string, in UpdateBackend) (*Backend, error)
	Get(in GetBackend) ([]*Backend, error)
	Delete(id string) error
}

const (
	BackendTypeLDAP             = "LDAP"
	BackendTypeOAuth2           = "OAuth2"
	BackendSyncPatient          = "PATIENT"
	BackendSyncEmployee         = "EMPLOYEE"
	BackendAttrSyncAlways       = "ALWAYS"
	BackendOAuth2ProviderGoogle = "GOOGLE"
	BackendOAuth2ProviderESIA   = "ESIA"
)

type Backend struct {
	ID          string `json:"ID" bson:"id"`
	Type        string `json:"type" bson:"type"`
	Description string `json:"description" bson:"description"`

	Sync       string                        `json:"sync" bson:"sync"`
	Attributes models.BackendAttributeParams `json:"attributes" bson:"attributes"`

	LDAPParams   `bson:",inline"`
	OAuth2Params `bson:",inline"`
}

type LDAPParams struct {
	Host     string            `json:"host" bson:"host"`
	Port     int64             `json:"port" bson:"port"`
	Cipher   string            `json:"cipher" bson:"cipher"`
	SSL      bool              `json:"ssl" bson:"ssl"`
	BindDN   string            `json:"bindDN" bson:"bindDN"`
	BaseDN   string            `json:"baseDN" bson:"baseDN"`
	Filter   string            `json:"filter" bson:"filter"`
	Password string            `json:"-" bson:"password"`
	Groups   map[string]string `json:"-" bson:"groups"`
}

type OAuth2Params struct {
	ClientID     string `json:"clientID" bson:"clientID"`
	ClientSecret string `json:"clientSecret" bson:"clientSecret"`
	// Enum: [GITHUB GOOGLE EMPLOYEE YANDEX ESIA]
	Provider string `json:"provider" bson:"provider"`
}

type StoreBackend struct {
	Backend `bson:",inline"`
}

type UpdateBackend struct {
	ID           *string                        `json:"id" bson:"id"`
	Type         *string                        `json:"type" bson:"type"`
	Description  *string                        `json:"description" bson:"description"`
	Sync         *string                        `json:"sync" bson:"sync"`
	Attributes   *models.BackendAttributeParams `json:"attributes" bson:"attributes"`
	Host         *string                        `json:"host" bson:"host"`
	Port         *int64                         `json:"port" bson:"port"`
	Cipher       *string                        `json:"cipher" bson:"cipher"`
	SSL          *bool                          `json:"ssl" bson:"ssl"`
	BindDN       *string                        `json:"bindDN" bson:"bindDN"`
	BaseDN       *string                        `json:"baseDN" bson:"baseDN"`
	Filter       *string                        `json:"filter" bson:"filter"`
	Password     *string                        `json:"-" bson:"password"`
	Groups       map[string]string              `json:"-" bson:"groups"`
	ClientID     *string                        `json:"clientID" bson:"clientID"`
	ClientSecret *string                        `json:"clientSecret" bson:"clientSecret"`
	Provider     *string                        `json:"provider" bson:"provider"`
}

type GetBackend struct {
	ID *string `json:"id" bson:"id"`
}
