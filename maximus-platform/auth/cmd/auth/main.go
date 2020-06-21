package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/getsentry/raven-go"
	"github.com/go-openapi/loads"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"repo.nefrosovet.ru/maximus-platform/auth/cmd"
	"repo.nefrosovet.ru/maximus-platform/auth/db/influx"
	dbMongo "repo.nefrosovet.ru/maximus-platform/auth/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/auth/jwt"
	"repo.nefrosovet.ru/maximus-platform/auth/storage"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"

	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations"
)

var (
	version = "No Version Provided"
	cfgFile string
)

func init() {
	restapi.Version = version

	cmd.SetVersion(version)
	cmd.SetConfigFile(cfgFile)
}

func main() {
	cmd.Execute(start)
}

func start() {
	// Init logger
	logLevel, err := log.ParseLevel(viper.GetString("logging.level"))
	if err != nil {
		log.WithError(err).
			Fatal("Parse logging level error")
	}

	log.SetLevel(logLevel)

	if viper.GetString("logging.output") != "STDOUT" {
		logFile, err := os.OpenFile(viper.GetString("logging.output"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.WithError(err).
				Fatal("Logging output opening error")
		}

		log.SetOutput(logFile)

		defer logFile.Close()
		defer log.SetOutput(os.Stdout)
	}

	if viper.GetString("logging.format") == "JSON" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// Connect to MongoDB
	mongoClient, err := dbMongo.Connect(
		viper.GetString("configDB.host"),
		viper.GetInt("configDB.port"),
		viper.GetString("configDB.login"),
		viper.GetString("configDB.password"),
		viper.GetString("configDB.database"),
	)
	if err != nil {
		log.WithError(err).Debug()
		log.WithFields(log.Fields{
			"context":  "CORE",
			"recurse":  "mongoDB",
			"addr":     viper.GetString("configDB.host") + ":" + strconv.Itoa(viper.GetInt("configDB.port")),
			"database": viper.GetString("configDB.database"),
			"status":   "FAILED",
		}).Info("Connection to influxDB")
	} else {
		log.WithFields(log.Fields{
			"context":  "CORE",
			"recurse":  "mongoDB",
			"addr":     viper.GetString("configDB.host") + ":" + strconv.Itoa(viper.GetInt("configDB.port")),
			"database": viper.GetString("configDB.database"),
			"status":   "CONNECTED",
		}).Info("Connection to mongoDB")
	}

	addr := viper.GetString("eventDB.host") + ":" + viper.GetString("eventDB.port")
	influxClient, err := influx.NewClient(
		addr,
		viper.GetString("eventDB.username"),
		viper.GetString("eventDB.password"),
		viper.GetString("eventDB.database"),
		viper.GetString("eventDB.retention"),
	)
	if err != nil {
		log.WithError(err).Debug()
		log.WithFields(log.Fields{
			"context":  "CORE",
			"recurse":  "eventDB",
			"addr":     addr,
			"database": viper.GetString("eventDB.database"),
			"status":   "FAILED",
		}).Fatal("Connection to influxDB")
	} else {
		log.WithFields(log.Fields{
			"context":  "CORE",
			"resource": "eventDB",
			"addr":     addr,
			"database": viper.GetString("eventDB.database"),
			"status":   "CONNECTED",
		}).Info("Connection to influxDB")
	}

	newAdminPassword := viper.GetString("adminPassword")
	if newAdminPassword == "" {
		panic("Can't start without admin password or default admin password")
	}

	hash, err := jwt.HashPassword(newAdminPassword)
	if err != nil {
		panic(fmt.Sprintf("Can't generate admin password hash, err: %s", err.Error()))
	}

	if err := st.InitDefaultStorage(mongoClient, influxClient); err != nil {
		log.WithFields(log.Fields{
			"function": "start",
			"error":    err,
		}).Fatal("Init default storage failed")
	}

	// Store default data

	_, err = st.GetStorage().AdminPasswordStorage.Store(storage.StoreAdminPassword{
		Hash: hash,
	})
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrAlreadyExists):
			_, err = st.GetStorage().AdminPasswordStorage.Update(storage.UpdateAdminPassword{
				Hash: hash,
			})
			if err != nil {
				log.WithError(err).
					Fatal("Admin password updating error")
			}
		default:
			log.WithError(err).
				Fatal("Admin password storing error")
		}
	}

	in := storage.StoreUser{
		User: storage.User{
			ID: "admin",
			Roles: map[string]bool{
				storage.RoleDefaultAdmin: true,
			},
		},
	}

	if _, err := st.GetStorage().UserStorage.Store(in); err != nil && !errors.Is(err, storage.ErrAlreadyExists) {
		log.WithError(err).
			Fatal("Admin user storing error")
	}

	if _, err = st.GetStorage().BackendsOrderStorage.Get(); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			if _, err = st.GetStorage().BackendsOrderStorage.Store(storage.StoreBackendsOrder{
				Order: storage.DefaultBackendsOrder,
			}); err != nil {
				log.WithError(err).
					Fatal("Default backends order storing error")
			}
		} else {
			log.WithError(err).
				Fatal("Backends order storing error")
		}
	}

	// Set sentry definitions
	if viper.GetString("sentryDSN") != "" {
		raven.SetDSN(viper.GetString("sentryDSN"))
	}

	log.WithFields(log.Fields{
		"context": "CORE",
		"version": restapi.Version,
		"status":  "STARTED",
	}).Info("Application started")

	// Start JSON API server
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewAuthAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// Apply server options
	server.Host = viper.GetString("http.host")
	server.Port = viper.GetInt("http.port")

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
