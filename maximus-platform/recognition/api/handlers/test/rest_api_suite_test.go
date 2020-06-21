package test

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/recognition/logger"
	"strconv"
	"time"
)

type RestApiTestingSuit struct {
	suite.Suite
	Host     string
	Port     int

	Keys     []string
	ImageID  *strfmt.UUID
	Image    *os.File
	PersonID []byte
}

func (suite *RestApiTestingSuit) InitServer() *restapi.Server{
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewRecognitionAPI(swaggerSpec)
	server := restapi.NewServer(api)

	l := logger.NewLogger(os.Stdout, "debug", "JSON")

	mock := NewMockService()
	server.SetHandler(restapi.ConfigureAPI(api, l.Api(90.00), &mock, "999.999.999"))

	server.Host = suite.Host
	server.Port = suite.Port

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalln(err)
		}
	}()
	time.Sleep(3 * time.Second)

	return server
}

func (suite *RestApiTestingSuit) GetBaseUrl() string {
	return "http://"+suite.Host+":"+strconv.Itoa(suite.Port)
}

