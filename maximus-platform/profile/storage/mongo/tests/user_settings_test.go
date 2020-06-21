package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/profile/storage"
	"repo.nefrosovet.ru/maximus-platform/profile/storage/mongo"
)

type UserSettingsStorageSuite struct {
	suite.Suite
}

func (s *UserSettingsStorageSuite) SetupSuite() {
	err := Client.Database.Collection(mongo.UserSettingsCollectionName).Drop(context.Background())
	s.Require().NoError(err)
}

func (s *UserSettingsStorageSuite) Test1Store() {
	_, err := Storage.StoreUserSettings(storage.StoreUserSettings{
		UserID:           "userID",
		TwoFAChannelType: "SMS",
		Locale:           "RUS",
	})
	s.Require().NoError(err)
}

func (s *UserSettingsStorageSuite) Test2Update() {
	_, err := Storage.UpdateUserSettings("userID", storage.UpdateUserSettings{
		TwoFAChannelType: storage.PtrS("EMAIL"),
		Locale:           storage.PtrS("FRA"),
	})
	s.Require().NoError(err)
}

func (s *UserSettingsStorageSuite) Test3GetUserSettings() {
	user, err := Storage.GetUserSettings(storage.GetUserSettings{
		UserID: "userID",
	})
	s.Require().NoError(err)

	s.Equal("EMAIL", user.TwoFAChannelType)
	s.Equal("FRA", user.Locale)
}

func TestUserSettingsStorage(t *testing.T) {
	suite.Run(t, new(UserSettingsStorageSuite))
}
