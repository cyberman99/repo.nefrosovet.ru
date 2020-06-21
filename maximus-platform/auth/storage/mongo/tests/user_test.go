package tests

import (
	"os"
	"testing"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/auth/cmd"
	dbMongo "repo.nefrosovet.ru/maximus-platform/auth/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	"repo.nefrosovet.ru/maximus-platform/auth/storage/mongo"
)

var (
	mongoClient *dbMongo.Client
)

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)

	cmd.Execute(func() {
		client, err := dbMongo.Connect(
			viper.GetString("configDB.host"),
			viper.GetInt("configDB.port"),
			viper.GetString("configDB.login"),
			viper.GetString("configDB.password"),
			viper.GetString("configDB.database"),
		)
		if err != nil {
			log.WithError(err).
				Fatal("Mongo connection error")
		}

		mongoClient = client

		os.Exit(m.Run())
	}, true)
}

type UserSuite struct {
	suite.Suite

	stored *storage.User

	s storage.UserStorage
}

func (s *UserSuite) SetupSuite() {
	userStorage, err := mongo.NewUserStorage(mongoClient)
	s.Require().NoError(err)

	s.s = userStorage
}

func (s *UserSuite) Test1Store() {
	user, err := s.s.Store(storage.StoreUser{
		ID: uuid.New().String(),
		Roles: map[string]bool{
			"test":   true,
			"test_2": true,
		},
		BackendEntryIDs: nil,
	})
	s.Require().NoError(err)

	s.stored = user
}

func (s *UserSuite) Test2Update() {
	_, err := s.s.Update(s.stored.ID, storage.UpdateUser{
		Roles: map[string]bool{
			"test_2": false,
			"test_3": true,
			"test_4": true,
		},
	})
	s.Require().NoError(err)
}

func (s *UserSuite) Test3Get() {
	users, err := s.s.Get(storage.GetUser{
		ID: &s.stored.ID,
	})
	s.Require().NoError(err)
	s.Require().Len(users, 1)

	s.Equal(s.stored.ID, users[0].ID)
	s.Equal(map[string]bool{
		"test": true,
		"test_3": true,
		"test_4": true,
	}, users[0].Roles)
	s.Equal(s.stored.BackendEntryIDs, users[0].BackendEntryIDs)
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
