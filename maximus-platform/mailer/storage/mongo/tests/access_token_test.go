package tests

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	dbMongo "repo.nefrosovet.ru/maximus-platform/mailer/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage/mongo"
)

var (
	AccessTokenStorage     storage.AccessTokenStorage
	ChannelStorage         storage.ChannelStorage
	TelegramContactStorage storage.TelegramContactStorage
	ViberContactStorage    storage.ViberContactStorage
)

func TestMain(m *testing.M) {
	host := flag.String("c.configDB.host", "127.0.0.1", "ConfigDB host")
	port := flag.Int("c.configDB.port", 27017, "ConfigDB port")
	username := flag.String("c.configDB.login", "", "ConfigDB login")
	password := flag.String("c.configDB.password", "", "ConfigDB password")
	database := flag.String("c.configDB.database", "mailer_config", "ConfigDB database name")
	flag.Parse()

	client, err := dbMongo.Connect(*host, *port, *username, *password, *database)
	if err != nil {
		panic(err)
	}

	err = mongo.Ensure(client)
	if err != nil {
		panic(err)
	}

	AccessTokenStorage = mongo.GetAccessTokenStorage(client)
	ChannelStorage = mongo.GetChannelStorage(client)
	TelegramContactStorage = mongo.GetTelegramContactStorage(client)
	ViberContactStorage = mongo.GetViberContactStorage(client)

	os.Exit(m.Run())
}

type AccessTokenSuite struct {
	s storage.AccessTokenStorage

	suite.Suite
}

func (s *AccessTokenSuite) SetupSuite() {
	s.s = AccessTokenStorage
}

func (s *AccessTokenSuite) Test1StoreAccessToken() {
	_, err := s.s.StoreAccessToken(storage.StoreAccessToken{
		AccessToken: storage.AccessToken{
			Token:       "token",
			Description: "description",
		},
	})
	if err != storage.ErrAccessTokenAlreadyExists {
		s.NoError(err)
	}
}

func (s *AccessTokenSuite) Test2GetAccessToken() {
	accessToken, err := s.s.GetAccessToken(storage.GetAccessToken{
		Token: "token",
	})
	s.NoError(err)

	s.Equal("description", accessToken.Description)
}

func (s *AccessTokenSuite) Test3GetAccessTokens() {
	limit := 1

	accessTokens, err := s.s.GetAccessTokens(storage.GetAccessTokens{
		Limit:  &limit,
		Offset: nil,
	})
	s.NoError(err)

	s.Len(accessTokens, 1)
}

func (s *AccessTokenSuite) Test4UpdateAccessToken() {
	accessToken, err := s.s.GetAccessToken(storage.GetAccessToken{
		Token: "token",
	})
	s.NoError(err)

	accessToken.Description = "description2"

	_, err = s.s.UpdateAccessToken(storage.UpdateAccessToken{
		Token:       "token",
		AccessToken: accessToken,
	})
	s.NoError(err)
}

func (s *AccessTokenSuite) Test5DeleteAccessToken() {
	_, err := s.s.DeleteAccessToken(storage.DeleteAccessToken{
		Token: "token",
	})
	s.NoError(err)
}

func TestAccessTokens(t *testing.T) {
	suite.Run(t, new(AccessTokenSuite))
}
