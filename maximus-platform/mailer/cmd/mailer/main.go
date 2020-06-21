package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/getsentry/raven-go"
	"github.com/go-openapi/loads"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"repo.nefrosovet.ru/maximus-platform/mailer/api/restapi"
	"repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations"
	dbInflux "repo.nefrosovet.ru/maximus-platform/mailer/db/influx"
	dbMongo "repo.nefrosovet.ru/maximus-platform/mailer/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/mailer/mq"
	"repo.nefrosovet.ru/maximus-platform/mailer/sender"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage/default"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage/influx"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage/mongo"
)

var (
	version = "No Version Provided"
	cfgFile string
)

func init() {
	sender.Version = version
	cmd.SetVersionTemplate(fmt.Sprintf("Version: %s\n", sender.Version))

	cobra.OnInitialize(initConfig)

	/*
		Set command flags
	*/

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file path")

	cmd.PersistentFlags().String("c.masterToken", "", "Master access token")

	cmd.PersistentFlags().String("c.http.host", "", "API host")
	cmd.PersistentFlags().Int("c.http.port", 0, "API port")

	cmd.PersistentFlags().String("c.configDB.host", "", "ConfigDB host")
	cmd.PersistentFlags().Int("c.configDB.port", 27017, "ConfigDB port")
	cmd.PersistentFlags().String("c.configDB.login", "", "ConfigDB login")
	cmd.PersistentFlags().String("c.configDB.database", "mailer_config", "ConfigDB database name")
	cmd.PersistentFlags().String("c.configDB.password", "", "ConfigDB password")

	cmd.PersistentFlags().String("c.eventDB.host", "", "EventDB host")
	cmd.PersistentFlags().Int("c.eventDB.port", 8086, "EventDB port")
	cmd.PersistentFlags().String("c.eventDB.login", "", "EventDB login")
	cmd.PersistentFlags().String("c.eventDB.password", "", "EventDB password")
	cmd.PersistentFlags().String("c.eventDB.database", "mailer", "EventDB database name")
	cmd.PersistentFlags().String("c.eventDB.retention", "", "EventDB retention policy")

	cmd.PersistentFlags().String("c.logging.output", "STDOUT", "Logging output")
	cmd.PersistentFlags().String("c.logging.level", "INFO", "Logging level")
	cmd.PersistentFlags().String("c.logging.format", "TEXT", "Logging format: TEXT or JSON")

	cmd.PersistentFlags().String("c.sentryDSN", "", "Sentry DSN")

	cmd.PersistentFlags().String("c.tracing.host", "", "Tracing host")
	cmd.PersistentFlags().Int("c.tracing.port", 0, "Tracing port")
	cmd.PersistentFlags().String("c.tracing.serviceName", "", "Tracing service name")

	cmd.PersistentFlags().String("c.mq.host", "", "Message broker host")
	cmd.PersistentFlags().Int("c.mq.port", 0, "Message broker port")
	cmd.PersistentFlags().String("c.mq.subClientID", "", "Message broker subscribe client ID")
	cmd.PersistentFlags().String("c.mq.pubClientID", "", "Message broker publication client ID")
	cmd.PersistentFlags().String("c.mq.login", "", "Message broker login")
	cmd.PersistentFlags().String("c.mq.password", "", "Message broker password")
	cmd.PersistentFlags().String("c.mq.topicIN", "", "Message broker topic IN")
	cmd.PersistentFlags().String("c.mq.topicOUT", "", "Message broker topic OUT")
	cmd.PersistentFlags().String("c.mq.qos", "", "Message broker quality of service level")

	cmd.PersistentFlags().String("c.botProxy.http.host", "", "BotProxy HTTP host")
	cmd.PersistentFlags().String("c.botProxy.http.path", "/", "BotProxy HTTP base path")
	cmd.PersistentFlags().String("c.botProxy.mq.host", "", "BotProxy MQ host")
	cmd.PersistentFlags().Int("c.botProxy.mq.port", 0, "BotProxy MQ port")

	cmd.PersistentFlags().Int("c.prometheus.port", 0, "Prometheus port")
	cmd.PersistentFlags().String("c.prometheus.path", "/metrics", "Prometheus path")

	/*
		Bind command flags to config variables
	*/

	viper.BindPFlag("masterToken", cmd.PersistentFlags().Lookup("c.masterToken"))

	viper.BindPFlag("http.host", cmd.PersistentFlags().Lookup("c.http.host"))
	viper.BindPFlag("http.port", cmd.PersistentFlags().Lookup("c.http.port"))

	viper.BindPFlag("configDB.host", cmd.PersistentFlags().Lookup("c.configDB.host"))
	viper.BindPFlag("configDB.port", cmd.PersistentFlags().Lookup("c.configDB.port"))
	viper.BindPFlag("configDB.login", cmd.PersistentFlags().Lookup("c.configDB.login"))
	viper.BindPFlag("configDB.password", cmd.PersistentFlags().Lookup("c.configDB.password"))
	viper.BindPFlag("configDB.database", cmd.PersistentFlags().Lookup("c.configDB.database"))

	viper.BindPFlag("eventDB.host", cmd.PersistentFlags().Lookup("c.eventDB.host"))
	viper.BindPFlag("eventDB.port", cmd.PersistentFlags().Lookup("c.eventDB.port"))
	viper.BindPFlag("eventDB.login", cmd.PersistentFlags().Lookup("c.eventDB.login"))
	viper.BindPFlag("eventDB.password", cmd.PersistentFlags().Lookup("c.eventDB.password"))
	viper.BindPFlag("eventDB.database", cmd.PersistentFlags().Lookup("c.eventDB.database"))
	viper.BindPFlag("eventDB.retention", cmd.PersistentFlags().Lookup("c.eventDB.retention"))

	viper.BindPFlag("logging.output", cmd.PersistentFlags().Lookup("c.logging.output"))
	viper.BindPFlag("logging.level", cmd.PersistentFlags().Lookup("c.logging.level"))
	viper.BindPFlag("logging.format", cmd.PersistentFlags().Lookup("c.logging.format"))

	viper.BindPFlag("sentryDSN", cmd.PersistentFlags().Lookup("c.sentryDSN"))

	viper.BindPFlag("tracing.host", cmd.PersistentFlags().Lookup("c.tracing.host"))
	viper.BindPFlag("tracing.port", cmd.PersistentFlags().Lookup("c.tracing.port"))
	viper.BindPFlag("tracing.serviceName", cmd.PersistentFlags().Lookup("c.tracing.serviceName"))

	viper.BindPFlag("mq.host", cmd.PersistentFlags().Lookup("c.mq.host"))
	viper.BindPFlag("mq.port", cmd.PersistentFlags().Lookup("c.mq.port"))
	viper.BindPFlag("mq.subClientID", cmd.PersistentFlags().Lookup("c.mq.subClientID"))
	viper.BindPFlag("mq.pubClientID", cmd.PersistentFlags().Lookup("c.mq.pubClientID"))
	viper.BindPFlag("mq.login", cmd.PersistentFlags().Lookup("c.mq.login"))
	viper.BindPFlag("mq.password", cmd.PersistentFlags().Lookup("c.mq.password"))
	viper.BindPFlag("mq.topicIN", cmd.PersistentFlags().Lookup("c.mq.topicIN"))
	viper.BindPFlag("mq.topicOUT", cmd.PersistentFlags().Lookup("c.mq.topicOUT"))
	viper.BindPFlag("mq.qos", cmd.PersistentFlags().Lookup("c.mq.qos"))

	viper.BindPFlag("botProxy.http.host", cmd.PersistentFlags().Lookup("c.botProxy.http.host"))
	viper.BindPFlag("botProxy.http.path", cmd.PersistentFlags().Lookup("c.botProxy.http.path"))
	viper.BindPFlag("botProxy.mq.host", cmd.PersistentFlags().Lookup("c.botProxy.mq.host"))
	viper.BindPFlag("botProxy.mq.port", cmd.PersistentFlags().Lookup("c.botProxy.mq.port"))

	viper.BindPFlag("prometheus.port", cmd.PersistentFlags().Lookup("c.prometheus.port"))
	viper.BindPFlag("prometheus.path", cmd.PersistentFlags().Lookup("c.prometheus.path"))
}

func initConfig() {
	// Use config file from the flag if provided.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Can't read config:", err)
		}
	}

	viper.SetEnvPrefix("MAILER")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func checkMandatoryParams() {
	var missing []string

	if value := viper.GetString("http.host"); value == "" {
		missing = append(missing, "http.host")
	}
	if value := viper.GetInt("http.port"); value == 0 {
		missing = append(missing, "http.port")
	}
	if value := viper.GetString("configDB.host"); value == "" {
		missing = append(missing, "configDB.host")
	}
	if value := viper.GetInt("configDB.port"); value == 0 {
		missing = append(missing, "configDB.port")
	}
	if value := viper.GetString("eventDB.host"); value == "" {
		missing = append(missing, "eventDB.host")
	}
	if value := viper.GetInt("eventDB.port"); value == 0 {
		missing = append(missing, "eventDB.port")
	}
	if value := viper.GetString("masterToken"); value == "" {
		missing = append(missing, "masterToken")
	}
	if len(missing) != 0 {
		log.WithField("missed", missing).Fatal("Missed mandatory params. Use --help flag or config")
	}
}

func applyAPIServerOptions(server *restapi.Server) {
	server.Host = viper.GetString("http.host")
	server.Port = viper.GetInt("http.port")
}

var cmd = &cobra.Command{
	Use:     "mailer",
	Short:   "Mailer is a message sender",
	Long:    `Just use it`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		checkMandatoryParams()

		start()
	},
}

func start() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewMailerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	applyAPIServerOptions(server)

	logLevel, err := log.ParseLevel(viper.GetString("logging.level"))
	if err != nil {
		log.WithError(err).Fatal("Can't set logging level")
	}

	log.SetLevel(logLevel)
	if viper.GetString("logging.output") != "STDOUT" {
		logFile, err := os.OpenFile(viper.GetString("logging.output"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		log.SetOutput(logFile)
		defer logFile.Close()
		defer log.SetOutput(os.Stdout)
	}
	if viper.GetString("logging.format") == "JSON" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// Set sentry definitions
	if viper.GetString("sentryDSN") != "" {
		raven.SetDSN(viper.GetString("sentryDSN"))
	}

	mongoClient, influxClient := initDB()

	if err := influx.Ensure(influxClient); err != nil {
		log.Fatal(err)
	}

	if err := mongo.Ensure(mongoClient); err != nil {
		log.Fatal(err)
	}

	_default.InitStorage(mongoClient, influxClient)

	server.ConfigureAPI()

	// start prometheus endpoint
	if viper.GetInt("prometheus.port") != 0 {
		path := viper.GetString("prometheus.path")
		port := viper.GetString("prometheus.port")

		go func() {
			http.Handle(path, promhttp.Handler())

			log.Fatal(http.ListenAndServe(":"+port, nil))
		}()
	}

	// Start MQ message sender
	if viper.GetString("mq.host") != "" {
		go initMQ()
	}

	log.WithFields(log.Fields{
		"context": "CORE",
		"version": sender.Version,
		"status":  "STARTED",
	}).Info("Application started")

	if err := server.Serve(); err != nil {
		log.WithFields(log.Fields{
			"context":  "CORE",
			"resource": "http",
			"addr":     viper.GetString("http.host") + ":" + strconv.Itoa(viper.GetInt("http.port")),
			"status":   "FAILED",
			"error":    err,
		}).Fatal("Starting HTTP server error")
	}
}

func initDB() (*dbMongo.Client, *dbInflux.Client) {
	// Connect to Mongo
	mongoClient, err := dbMongo.Connect(
		viper.GetString("configDB.host"),
		viper.GetInt("configDB.port"),
		viper.GetString("configDB.login"),
		viper.GetString("configDB.password"),
		viper.GetString("configDB.database"),
	)
	if err != nil {
		log.WithError(err).Debug("Can't connect to configDB")

		log.WithFields(log.Fields{
			"context":  "CORE",
			"resource": "configDB",
			"addr":     viper.GetString("configDB.host") + ":" + strconv.Itoa(viper.GetInt("configDB.port")),
			"status":   "FAILED",
		}).Fatal("Connection to configDB failed")
	}
	log.WithFields(log.Fields{
		"context":  "CORE",
		"resource": "configDB",
		"addr":     viper.GetString("configDB.host") + ":" + strconv.Itoa(viper.GetInt("configDB.port")),
		"status":   "CONNECTED",
	}).Info("Connection to configDB established")

	// Connect to Influx

	influxClient, err := dbInflux.Connect(
		viper.GetString("eventDB.host"),
		viper.GetInt("eventDB.port"),
		viper.GetString("eventDB.username"),
		viper.GetString("eventDB.password"),
		viper.GetString("eventDB.database"),
		viper.GetString("eventDB.retention"),
	)
	if err != nil {
		log.WithError(err).Debug("Can't connect to eventDB")

		log.WithFields(log.Fields{
			"context":  "CORE",
			"resource": "eventDB",
			"addr":     viper.GetString("eventDB.host") + ":" + strconv.Itoa(viper.GetInt("eventDB.port")),
			"status":   "FAILED",
		}).Fatal("Connection to eventDB failed")
	}

	log.WithFields(log.Fields{
		"context":  "CORE",
		"resource": "eventDB",
		"addr":     viper.GetString("eventDB.host") + ":" + strconv.Itoa(viper.GetInt("eventDB.port")),
		"status":   "CONNECTED",
	}).Info("Connection to eventDB established")

	return mongoClient, influxClient
}

func initMQ() {
	mqClient := mq.MQ{
		Addr: fmt.Sprintf("%s:%d", viper.GetString("mq.host"), viper.GetInt("mq.port")),

		PubClientID: viper.GetString("mq.pubClientID"),
		SubClientID: viper.GetString("mq.subClientID"),
		Login:       viper.GetString("mq.login"),
		Password:    viper.GetString("mq.password"),
		TopicIn:     viper.GetString("mq.topicIN"),
		TopicOut:    viper.GetString("mq.topicOUT"),
		QoS:         byte(viper.GetInt("mq.qos")),
	}

	mqClient.Listen()
	mqClient.GetPubClient()
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
