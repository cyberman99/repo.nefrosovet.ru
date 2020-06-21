package repos

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
)

type RouteSuite struct {
	suite.Suite

	lastUUID strfmt.UUID
	r        RouteRepository
}

func (s *RouteSuite) SetupSuite() {
	err := mongoClient.Collection(domain.RouteCollectionName).Drop(context.Background())
	s.Require().NoError(err)

	s.r = NewRouteRepo(mongoClient)
}

func (s *RouteSuite) Test1Set() {
	replyID := strfmt.UUID("REPLY_ID-1")
	route, err := s.r.Set(domain.Route{
		ReplyID: &replyID,
		Src:     domain.Source{},
	})
	s.NoError(err)

	s.lastUUID = route.RouteID
}

func (s *RouteSuite) Test2Get() {
	route, err := s.r.Get(s.lastUUID)
	s.NoError(err)

	s.NotNil(route.ReplyID)
	s.Equal(strfmt.UUID("REPLY_ID-1"), *route.ReplyID)
}

func (s *RouteSuite) Test3Update() {
	replyID := strfmt.UUID("REPLY_ID-2")
	_, err := s.r.Update(s.lastUUID, domain.Route{
		ReplyID: &replyID,
		Src:     domain.Source{},
	})
	s.NoError(err)

	route, err := s.r.Get(s.lastUUID)
	s.NoError(err)

	s.NotNil(route.ReplyID)
	s.Equal(strfmt.UUID("REPLY_ID-2"), *route.ReplyID)
}

func (s *RouteSuite) Test4UnsetReplyID() {
	err := s.r.UnsetReplyID(s.lastUUID)
	s.NoError(err)

	route, err := s.r.Get(s.lastUUID)
	s.NoError(err)

	s.Nil(route.ReplyID)
}

func (s *RouteSuite) Test5Delete() {
	err := s.r.Delete(s.lastUUID)
	s.NoError(err)
}

func (s *RouteSuite) Test6List() {
	uuids := []strfmt.UUID{}

	routesBefore, err := s.r.List(domain.RoutesFilter{})
	if err != domain.ErrRouteNotFound {
		s.NoError(err)
	}
	routesLenBefore := len(routesBefore)

	for i := 0; i < 10; i++ {
		replyID := strfmt.UUID(fmt.Sprintf("REPLY_ID-%d", i))
		route, err := s.r.Set(domain.Route{
			ReplyID: &replyID,
			Src:     domain.Source{},
		})
		s.NoError(err)

		uuids = append(uuids, route.RouteID)
	}

	routes, err := s.r.List(domain.RoutesFilter{
		Limit: 100,
	})
	s.NoError(err)

	s.Len(routes, routesLenBefore+10)

	for _, uuid := range uuids {
		err := s.r.Delete(uuid)
		s.NoError(err)
	}
}

func TestRoute(t *testing.T) {
	suite.Run(t, new(RouteSuite))
}
