package services

import (
	"errors"
	"github.com/go-openapi/strfmt"
	"repo.nefrosovet.ru/maximus-platform/recognition/logger"
	"repo.nefrosovet.ru/maximus-platform/recognition/services/AWS"
)

type ServiceType int

const (
	AmazonService ServiceType = iota
	GoogleService
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrUnknownService = errors.New("unknown service")
)

type CloudServicer interface {
	Set(patientID string, format string, body []byte) (imgID *strfmt.UUID, err error)
	Get(photoID string) (patientID string, err error)
	Delete(photoID string) (err error)
	List(limit *int64) (keys []*strfmt.UUID, err error)
	Rekognize(sourceImage []byte) (_ AWS.MatchedFaces, err error)
	ServiceName() string
}

type Config struct {
	Bucket       string
	Similarity   float64
	AccessID     string
	AccessSecret string
	Region       string
}

func NewService(s ServiceType, conf *Config, l logger.CoreEntrier) (CloudServicer, error) {
	switch s {
	case AmazonService:
		return AWS.NewAWSClient(
			conf.Similarity,
			conf.Bucket,
			conf.AccessID,
			conf.AccessSecret,
			conf.Region,
			l)
	case GoogleService:
		// TODO
		return nil, ErrNotImplemented
	default:
		return nil, ErrUnknownService
	}
}
