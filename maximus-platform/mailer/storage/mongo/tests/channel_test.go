package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

type ChannelSuite struct {
	s storage.ChannelStorage

	suite.Suite
}

func (s *ChannelSuite) SetupSuite() {
	s.s = ChannelStorage
}

func (s *ChannelSuite) Test1StoreChannel() {
	sc := storage.StoreChannel{
		Channel: storage.Channel{
			ID:          "id",
			Type:        "type",
			AccessToken: "accessToken",
		},
	}

	sc.Channel.Params.Token = "token"

	_, err := s.s.StoreChannel(sc)
	if err != storage.ErrChannelAlreadyExists {
		s.NoError(err)
	}
}

func (s *ChannelSuite) Test2GetChannel() {
	id := "id"
	token := "token"

	gc := storage.GetChannel{
		ID: &id,
	}

	gc.Params.Token = &token

	channel, err := s.s.GetChannel(gc)
	s.NoError(err)

	s.Equal("id", channel.ID)
	s.Equal(storage.ChannelType("type"), channel.Type)
	s.Equal("accessToken", channel.AccessToken)
	s.Equal("token", channel.Params.Token)
}

func (s *ChannelSuite) Test3GetChannels() {
	typ := storage.ChannelType("type")

	channels, err := s.s.GetChannels(storage.GetChannels{
		Type: &typ,
	})
	s.NoError(err)
	s.Len(channels, 1)
}

func (s *ChannelSuite) Test4UpdateChannel() {
	id := "id"
	token := "token"

	gc := storage.GetChannel{
		ID: &id,
	}

	gc.Params.Token = &token

	channel, err := s.s.GetChannel(gc)
	s.NoError(err)

	_, err = s.s.UpdateChannel(storage.UpdateChannel{
		ID:      id,
		Channel: channel,
	})
	s.NoError(err)
}

func (s *ChannelSuite) Test5DeleteChannel() {
	_, err := s.s.DeleteChannel(storage.DeleteChannel{
		ID: "id",
	})
	s.NoError(err)
}

func TestChannels(t *testing.T) {
	suite.Run(t, new(ChannelSuite))
}
