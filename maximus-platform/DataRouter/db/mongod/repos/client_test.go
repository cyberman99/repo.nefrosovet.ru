package repos

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
)

var (
	mongoClient mongod.Storer

	host     = os.Getenv("DATAROUTER_CONFIGDB_HOST")
	strPort  = os.Getenv("DATAROUTER_CONFIGDB_PORT")
	database = os.Getenv("DATAROUTER_CONFIGDB_DATABASE")
)

func TestMain(m *testing.M) {
	port, err := strconv.Atoi(strPort)
	if err != nil {
		panic(err)
	}

	mClient, err := mongod.NewCli(
		host,
		port,
		"",
		"",
		database,
	)
	if err != nil {
		panic(err)
	}

	err = mClient.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	mongoClient = mClient

	os.Exit(m.Run())
}

type ClientSuite struct {
	suite.Suite

	r ClientRepository
}

func (s *ClientSuite) SetupSuite() {
	err := mongoClient.Collection(domain.ClientCollectionName).Drop(context.Background())
	s.Require().NoError(err)

	s.r = NewClientRepo(mongoClient, nil)
}

func (s *ClientSuite) Test1Set() {
	_, err := s.r.Set(domain.Client{
		ClientID:   "CLIENT_ID-1",
		Passhash:   "PASSHASH-1",
		Username:   "USERNAME-1",
		Mountpoint: "MOUNTPOINT-1",
	})
	s.NoError(err)
}

func (s *ClientSuite) Test2Get() {
	client, err := s.r.Get("CLIENT_ID-1")
	s.NoError(err)

	s.Equal("PASSHASH-1", client.Passhash)
	s.Equal("USERNAME-1", client.Username)
	s.Equal("MOUNTPOINT-1", client.Mountpoint)
}

func (s *ClientSuite) Test3Update() {
	_, err := s.r.Update("CLIENT_ID-1", domain.Client{
		ClientID:    "CLIENT_ID-2",
		Passhash:    "PASSHASH-2",
		Username:    "USERNAME-2",
		Created:     time.Time{},
		Mountpoint:  "MOUNTPOINT-2",
		Permissions: domain.Permissions{},
	})
	s.NoError(err)

	client, err := s.r.Get("CLIENT_ID-2")
	s.NoError(err)

	s.Equal("PASSHASH-2", client.Passhash)
	s.Equal("USERNAME-2", client.Username)
	s.Equal("MOUNTPOINT-2", client.Mountpoint)
}

func (s *ClientSuite) Test4SetOrReplace() {
	_, err := s.r.SetOrReplace("CLIENT_ID-2", domain.Permissions{
		Publish: []domain.Acl{{
			Pattern: "1",
		}},
		Subscribe: []domain.Acl{{
			Pattern: "2",
		}},
	})
	s.NoError(err)
}

func (s *ClientSuite) Test5GetPermissions() {
	permissions, err := s.r.GetPermissions("CLIENT_ID-2")
	s.NoError(err)

	s.Equal("1", permissions.Publish[0].Pattern)
	s.Equal("2", permissions.Subscribe[0].Pattern)
}

func (s *ClientSuite) Test6Delete() {
	err := s.r.Delete("CLIENT_ID-2")
	s.NoError(err)
}

func (s *ClientSuite) Test7List() {
	uuids := []strfmt.UUID{}

	for i := 0; i < 10; i++ {
		client, err := s.r.Set(domain.Client{
			ClientID:   fmt.Sprintf("CLIENT_ID-%d", i),
			Passhash:   "PASSHASH",
			Username:   "USERNAME",
			Mountpoint: "MOUNTPOINT",
		})
		s.NoError(err)

		uuids = append(uuids, strfmt.UUID(client.ClientID))
	}

	clients, err := s.r.List(domain.ClientsFilter{
		Limit: 100,
	})
	s.NoError(err)

	s.Len(clients, 10)

	for _, uuid := range uuids {
		err := s.r.Delete(uuid)
		s.NoError(err)
	}
}

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}
