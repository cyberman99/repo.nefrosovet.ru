package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	"repo.nefrosovet.ru/maximus-platform/auth/storage/mongo"
)

type AdminPasswordSuite struct {
	suite.Suite

	stored *storage.AdminPassword

	s storage.AdminPasswordStorage
}

func (s *AdminPasswordSuite) SetupSuite() {
	adminPasswordStorage, err := mongo.NewAdminPasswordStorage(mongoClient)
	s.Require().NoError(err)

	s.s = adminPasswordStorage
}

func (s *AdminPasswordSuite) Test1Store() {
	user, err := s.s.Store(storage.StoreAdminPassword{
		Hash: "hash",
	})
	if !errors.Is(err, storage.ErrAlreadyExists) {
		s.Require().NoError(err)
	}

	s.stored = user
}

func (s *AdminPasswordSuite) Test2Update() {
	_, err := s.s.Update(storage.UpdateAdminPassword{
		Hash: "hash2",
	})
	s.Require().NoError(err)
}

func (s *AdminPasswordSuite) Test3Get() {
	adminPasswords, err := s.s.Get()
	s.Require().NoError(err)
	s.Require().Len(adminPasswords, 1)

	s.Equal("hash2", adminPasswords[0].Hash)
}

func TestAdminPassword(t *testing.T) {
	suite.Run(t, new(AdminPasswordSuite))
}