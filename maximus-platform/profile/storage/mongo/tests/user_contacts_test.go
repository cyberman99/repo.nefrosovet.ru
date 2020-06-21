package tests

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/profile/storage"
	"repo.nefrosovet.ru/maximus-platform/profile/storage/mongo"
)

type UserContactsStorageSuite struct {
	suite.Suite
}

func (s *UserContactsStorageSuite) SetupSuite() {
	err := Client.Database.Collection(mongo.UserContactsCollectionName).Drop(context.Background())
	s.Require().NoError(err)
}

func (s *UserContactsStorageSuite) Test1Store() {
	_, err := Storage.StoreUserContact(storage.StoreUserContact{
		ID:       uuid.New().String(),
		UserID:   "userID",
		TypeCODE: "EMAIL",
		Value:    "name@company.com",
	})
	s.Require().NoError(err)

	_, err = Storage.StoreUserContact(storage.StoreUserContact{
		ID:       uuid.New().String(),
		UserID:   "userID",
		TypeCODE: "MOBILE",
		Value:    "89991234567",
	})
	s.Require().NoError(err)
}

func (s *UserContactsStorageSuite) Test2Update() {
	_, err := Storage.UpdateUserContact("userID", storage.UpdateUserContact{
		TypeCODE: storage.PtrS("EMAIL"),
		Value:    storage.PtrS("name@company.com"),
		Verified: storage.PtrB(true),
	})
	s.Require().NoError(err)
}

func (s *UserContactsStorageSuite) Test3GetUserContacts() {
	contacts, err := Storage.GetUserContacts(storage.GetUserContacts{
		UserID: "userID",
	})
	s.Require().NoError(err)

	//s.Equal("userID", string(contacts[0].UserID))
	s.Equal("EMAIL", contacts[0].TypeCODE)
	s.Equal("name@company.com", contacts[0].Value)
	s.Equal(true, contacts[0].Verified != nil)

	//s.Equal("userID", string(contacts[1].UserID))
	s.Equal("MOBILE", contacts[1].TypeCODE)
	s.Equal("89991234567", contacts[1].Value)
	s.Equal(false, contacts[1].Verified != nil)
}

func TestUserContactsStorage(t *testing.T) {
	suite.Run(t, new(UserContactsStorageSuite))
}