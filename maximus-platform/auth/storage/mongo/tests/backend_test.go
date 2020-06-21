package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/auth/api/models"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	"repo.nefrosovet.ru/maximus-platform/auth/storage/mongo"
)

type BackendSuite struct {
	suite.Suite

	stored *storage.Backend

	s storage.BackendStorage
}

func (s *BackendSuite) SetupSuite() {
	backendStorage, err := mongo.NewBackendStorage(mongoClient)
	s.Require().NoError(err)

	s.s = backendStorage
}

func (s *BackendSuite) Test1Store() {
	backend, err := s.s.Store(storage.StoreBackend{
		Backend: storage.Backend{
			ID:          "id",
			Type:        "type",
			Description: "description",
			Sync:        "sync",
			Attributes:  models.BackendAttributeParams{},
			LDAPParams: storage.LDAPParams{
				Groups: map[string]string{
					"1": "2",
					"2": "3",
				},
			},
			OAuth2Params: storage.OAuth2Params{},
		},
	})
	s.Require().NoError(err)

	s.stored = backend
}

func (s *BackendSuite) Test2Update() {
	_, err := s.s.Update(s.stored.ID, storage.UpdateBackend{
		Type:        ptrS("type2"),
		Description: ptrS("description2"),
		Sync:        ptrS("sync2"),
		Groups: map[string]string{
			"2": "",
			"3": "4",
		},
	})
	s.Require().NoError(err)
}

func (s *BackendSuite) Test3Get() {
	backends, err := s.s.Get(storage.GetBackend{
		ID: &s.stored.ID,
	})
	s.Require().NoError(err)
	s.Require().Len(backends, 1)

	s.Equal(s.stored.ID, backends[0].ID)
	s.Equal("type2", backends[0].Type)
	s.Equal("description2", backends[0].Description)
	s.Equal("sync2", backends[0].Sync)
	s.Equal(map[string]string{
		"1": "2",
		"3": "4",
	}, backends[0].Groups)
}

func (s *BackendSuite) Test4Delete() {
	err := s.s.Delete(s.stored.ID)
	s.Require().NoError(err)
}

func ptrS(s string) *string {
	return &s
}

func TestBackend(t *testing.T) {
	suite.Run(t, new(BackendSuite))
}
