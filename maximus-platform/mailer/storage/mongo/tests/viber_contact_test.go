package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

type ViberContactSuite struct {
	s storage.ViberContactStorage

	suite.Suite
}

func (s *ViberContactSuite) SetupSuite() {
	s.s = ViberContactStorage
}

func (s *ViberContactSuite) Test1StoreViberContact() {
	_, err := s.s.StoreViberContact(storage.StoreViberContact{
		ViberContact: storage.ViberContact{
			Phone:  "phone",
			UserID: "userID",
			Name:   "name",
			Token:  "token",
		},
	})
	if err != storage.ErrViberContactAlreadyExists {
		s.NoError(err)
	}
}

func (s *ViberContactSuite) Test2GetViberContact() {
	phone := "phone"
	userID := "userID"

	_, err := s.s.GetViberContact(storage.GetViberContact{
		Phone:  &phone,
		UserID: &userID,
	})
	s.NoError(err)
}

func (s *ViberContactSuite) Test3UpdateViberContact() {
	phone := "phone"
	userID := "userID"

	viberContact, err := s.s.GetViberContact(storage.GetViberContact{
		Phone:  &phone,
		UserID: &userID,
	})
	s.NoError(err)

	_, err = s.s.UpdateViberContact(storage.UpdateViberContact{
		Phone:        &phone,
		UserID:       &userID,
		ViberContact: viberContact,
	})
	s.NoError(err)
}

func TestViberContacts(t *testing.T) {
	suite.Run(t, new(ViberContactSuite))
}
