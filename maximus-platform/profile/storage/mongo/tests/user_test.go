package tests

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	dbMongo "repo.nefrosovet.ru/maximus-platform/profile/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/profile/storage"
	"repo.nefrosovet.ru/maximus-platform/profile/storage/mongo"
)

var (
	Client  *dbMongo.Client
	Storage *storage.Storage
)

func TestMain(m *testing.M) {
	mongoClient, err := dbMongo.Connect(&dbMongo.Config{
		Host:     "localhost",
		Port:     27017,
		Database: "profile_test",
	})
	if err != nil {
		panic(err)
	}

	mongoStorage, err := mongo.New(mongoClient)
	if err != nil {
		panic(err)
	}

	Client = mongoClient
	Storage = mongoStorage

	os.Exit(m.Run())
}

type UserStorageSuite struct {
	lastUUID storage.UUID

	suite.Suite
}

func (s *UserStorageSuite) SetupSuite() {
	err := Client.Database.Collection(mongo.UserCollectionName).Drop(context.Background())
	s.Require().NoError(err)
}

func (s *UserStorageSuite) Test1Store() {
	user, err := Storage.StoreUser(storage.StoreUser{
		ID:           uuid.New().String(),
		PasswordHash: "passwordHash",
		FirstName:    "firstName",
		LastName:     "lastName",
		MiddleName:   "middleName",
	})
	s.Require().NoError(err)

	s.lastUUID = user.ID
}

func (s *UserStorageSuite) Test2Update() {
	_, err := Storage.UpdateUser(s.lastUUID, storage.UpdateUser{
		PasswordHash: storage.PtrS("passwordHash2"),
		FirstName:    storage.PtrS("firstName2"),
		LastName:     storage.PtrS("lastName2"),
		MiddleName:   storage.PtrS("middleName2"),
	})
	s.Require().NoError(err)
}

func (s *UserStorageSuite) Test3GetUser() {
	user, err := Storage.GetUser(storage.GetUser{
		ID: &s.lastUUID,
	})
	s.Require().NoError(err)

	s.Equal("passwordHash2", user.PasswordHash)
	s.Equal("firstName2", user.FirstName)
	s.Equal("lastName2", user.LastName)
	s.Equal("middleName2", user.MiddleName)
}

func (s *UserStorageSuite) Test4GetUsers() {
	_, err := Storage.StoreUser(storage.StoreUser{
		ID:           uuid.New().String(),
		PasswordHash: "passwordHash3",
		FirstName:    "firstName3",
		LastName:     "lastName3",
		MiddleName:   "middleName3",
	})
	s.Require().NoError(err)

	users, err := Storage.GetUsers()
	s.Require().NoError(err)

	s.Len(users, 2)
}

func TestUserStorage(t *testing.T) {
	suite.Run(t, new(UserStorageSuite))
}