package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	mgo "go.mongodb.org/mongo-driver/mongo"

	"repo.nefrosovet.ru/maximus-platform/profile/api"
	"repo.nefrosovet.ru/maximus-platform/profile/api/server"
	"repo.nefrosovet.ru/maximus-platform/profile/cmd"
	dbMongo "repo.nefrosovet.ru/maximus-platform/profile/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/profile/pkg/apierrors"
	"repo.nefrosovet.ru/maximus-platform/profile/pkg/middleware"
	"repo.nefrosovet.ru/maximus-platform/profile/storage/mongo"
	"repo.nefrosovet.ru/maximus-platform/profile/logger"
)

var (
	version = "No Version Provided"
	cfgFile string

	sentryHub *sentry.Hub
	lg logger.Logger
)

func init() {
	cmd.CFGFile = cfgFile
	cmd.SetVersion(version)
}

func main() {
	cmd.Execute(func() {
		setupLogging()
		setupSentry()

		start()
	})
}

func start() {
	ctx := context.Background()
	mongoClient, err := dbMongo.Connect(&dbMongo.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.login"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.database"),
	}, ctx)
	if err != nil {
		lg.Core().Debug("Database connection error:", err)
		lg.Core().Fatal(logger.COREDB, viper.GetString("db.host"), viper.GetString("db.port"),
			logger.COREFAILED)
	}


	err = mongoClient.Session(func(mctx mgo.SessionContext) error {
		_, err := mongoClient.Collection("test").InsertOne(mctx, bson.D{{ "1", "test"}})
		if err != nil { return err}
		_, err = mongoClient.Collection("test").InsertOne(mctx, bson.D{{ "2", "test"}})
		if err != nil { return err}
		return errors.New("test error")
	})
	fmt.Println(err, "????????")
	cur, _ := mongoClient.Collection("test").Find(ctx, bson.D{})
	for cur.Next(ctx) {
		fmt.Println(cur.Current.String())
	}
	mongoClient.Collection("test").DeleteMany(ctx, bson.D{})



	storage, err := mongo.New(mongoClient)
	if err != nil {
		lg.Core().Debug("Database ensure error:", err)
		lg.Core().Fatal(logger.COREDB, "", "", logger.COREFAILED)
	}

	lg.Core().Info(logger.COREDB, viper.GetString("db.host"),
		viper.GetString("db.port"), "", logger.CORECONNECTED)

	// Get swagger specification
	swagger, err := api.GetSwagger()
	if err != nil {
		lg.Core().Debug("Swagger specification loading error", err)
		lg.Core().Fatal("", "", "", logger.COREFAILED)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Set up echo router
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.HTTPErrorHandler = apierrors.NewHandler(version).WithSentry(sentryHub).Handle

	e.Use(echoMiddleware.RecoverWithConfig(echoMiddleware.RecoverConfig{
		DisablePrintStack: true,
	}))
	// Use validation middleware to check all requests against the OpenAPI schema.
	e.Use(middleware.OAPIRequestValidator(swagger))

	srv := server.New(version, storage, lg, mongoClient)
	api.RegisterHandlers(e, srv)

	errCh := make(chan error)
	go func() {
		if err = e.Start(fmt.Sprintf(
			"%s:%d",
			viper.GetString("http.host"),
			viper.GetInt("http.port"),
		)); err != nil {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		lg.Core().Debug("Server listening error:", err)
		lg.Core().Fatal("", viper.GetString("http.host"), viper.GetString("http.port"),
			logger.COREFAILED)
	case <-time.After(time.Second * 1):
		lg.Core().Info("", "", "", version, logger.CORESTARTED)

		if err := <-errCh; err != nil {
			lg.Core().Debug("Server error:", err)
			lg.Core().Fatal("", viper.GetString("http.host"), viper.GetString("http.port"),
				logger.COREFAILED)
		}
	}
}

func setupSentry() {
	if dsn := viper.GetString("sentryDSN"); dsn != "" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: dsn,
		})
		if err != nil {
			lg.Core().Debug("Sentry initialization error:", err)
			lg.Core().Fatal("", "", "", logger.COREFAILED)
		}

		sentryHub = sentry.CurrentHub()
	}
}

func setupLogging() {
	var file *os.File

	if viper.GetString("logging.output") != "STDOUT" {
		file, err := os.OpenFile(viper.GetString("logging.output"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.WithError(err).
				Fatalf("Opening log file error")
		}

		defer func() {
			if err := file.Close(); err != nil {
				log.WithError(err).
					Error("Logging file closing error")
			}
		}()
	} else {
		file = os.Stdout
	}

	lg = logger.NewLogger(file, viper.GetString("logging.level"), viper.GetString("logging.format"))
	lg.Core().Info("", "", "", version, logger.CORESTARTED)
}