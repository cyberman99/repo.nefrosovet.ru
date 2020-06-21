package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

type TelegramContactSuite struct {
	s storage.TelegramContactStorage

	suite.Suite
}

func (s *TelegramContactSuite) SetupSuite() {
	s.s = TelegramContactStorage
}

func (s *TelegramContactSuite) Test1StoreTelegramContact() {
	_, err := s.s.StoreTelegramContact(storage.StoreTelegramContact{
		TelegramContact: storage.TelegramContact{
			Phone:    "phone",
			ChatID:   1,
			Username: "username",
		},
	})
	if err != storage.ErrTelegramContactAlreadyExists {
		s.NoError(err)
	}
}

func (s *TelegramContactSuite) Test2GetTelegramContact() {
	phone := "phone"
	chatID := int64(1)

	_, err := s.s.GetTelegramContact(storage.GetTelegramContact{
		Phone:  &phone,
		ChatID: &chatID,
	})
	s.NoError(err)
}

func (s *TelegramContactSuite) Test3UpdateTelegramContact() {
	phone := "phone"
	chatID := int64(1)

	telegramContact, err := s.s.GetTelegramContact(storage.GetTelegramContact{
		Phone:  &phone,
		ChatID: &chatID,
	})
	s.NoError(err)

	telegramContact.Username = "username2"

	_, err = s.s.UpdateTelegramContact(storage.UpdateTelegramContact{
		Phone:           &phone,
		ChatID:          &chatID,
		TelegramContact: telegramContact,
	})
	s.NoError(err)
}

func TestTelegramContacts(t *testing.T) {
	suite.Run(t, new(TelegramContactSuite))
}
