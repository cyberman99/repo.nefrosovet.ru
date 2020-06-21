package repos

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
)

type ReplySuite struct {
	suite.Suite

	lastUUID strfmt.UUID
	r        ReplyRepository
}

func (s *ReplySuite) SetupSuite() {
	err := mongoClient.Collection(domain.ReplyCollectionName).Drop(context.Background())
	s.Require().NoError(err)

	s.r = NewReplyRepo(mongoClient)
}

func (s *ReplySuite) Test1Set() {
	description := "DESCRIPTION-1"
	reply, err := s.r.Set(domain.Reply{
		Description: &description,
		Regex:       "REGEX-1",
		Replace:     "REPLACE-1",
	})
	s.NoError(err)

	s.lastUUID = reply.ReplyID
}

func (s *ReplySuite) Test2Get() {
	reply, err := s.r.Get(s.lastUUID)
	s.NoError(err)

	s.NotNil(reply.Description)
	s.Equal("DESCRIPTION-1", *reply.Description)
	s.Equal("REGEX-1", reply.Regex)
	s.Equal("REPLACE-1", reply.Replace)
}

func (s *ReplySuite) Test3Update() {
	description := "DESCRIPTION-2"
	_, err := s.r.Update(s.lastUUID, domain.Reply{
		Description: &description,
		Regex:       "REGEX-2",
		Replace:     "REPLACE-2",
	})
	s.NoError(err)

	reply, err := s.r.Get(s.lastUUID)
	s.NoError(err)

	s.NotNil(reply.Description)
	s.Equal("DESCRIPTION-2", *reply.Description)
	s.Equal("REGEX-2", reply.Regex)
	s.Equal("REPLACE-2", reply.Replace)
}

func (s *ReplySuite) Test4Delete() {
	err := s.r.Delete(s.lastUUID)
	s.NoError(err)
}

func (s *ReplySuite) Test5List() {
	uuids := []strfmt.UUID{}

	repliesBefore, err := s.r.List(domain.RepliesFilter{})
	if err != domain.ErrReplyNotFound {
		s.NoError(err)
	}
	repliesLenBefore := len(repliesBefore)

	description := "DESCRIPTION"
	for i := 0; i < 10; i++ {
		reply, err := s.r.Set(domain.Reply{
			Description: &description,
			Regex:       fmt.Sprintf("REGEX-%d", i),
			Replace:     fmt.Sprintf("REPLACE-%d", i),
		})
		s.NoError(err)

		uuids = append(uuids, reply.ReplyID)
	}

	replies, err := s.r.List(domain.RepliesFilter{
		Limit: 100,
	})
	s.NoError(err)

	s.Len(replies, repliesLenBefore+10)

	for _, uuid := range uuids {
		err := s.r.Delete(uuid)
		s.NoError(err)
	}
}

func TestReply(t *testing.T) {
	suite.Run(t, new(ReplySuite))
}
