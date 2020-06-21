package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/getsentry/raven-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"repo.nefrosovet.ru/maximus-platform/apigw/gateway"
	"repo.nefrosovet.ru/maximus-platform/apigw/influxdb"
	"repo.nefrosovet.ru/maximus-platform/apigw/mongodb"
)

var (
	version = "No Version Provided"
	cfgFile string
)
var cmd = &cobra.Command{
	Use:     "apigw",
	Short:   "Gateway service",
	Long:    `Just use it`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		checkMandatoryParams()
		start()
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

func start() {
	// Init logger
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

	// Connect to MongoDB
	db, err := mongodb.New(
		viper.GetString("configDB.host"),
		viper.GetInt("configDB.port"),
		viper.GetString("configDB.login"),
		viper.GetString("configDB.password"),
		viper.GetString("configDB.database"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// Connect to InfluxDB
	influxdb.New(
		viper.GetString("eventDB.host")+":"+viper.GetString("eventDB.port"),
		viper.GetString("eventDB.username"),
		viper.GetString("eventDB.password"),
		viper.GetString("eventDB.database"),
		viper.GetString("eventDB.retention"),
	)

	// Set sentry definitions
	if viper.GetString("sentryDSN") != "" {
		raven.SetDSN(viper.GetString("sentryDSN"))
	}

	log.WithFields(log.Fields{
		"context": "CORE",
		"version": version,
		"status":  "STARTED",
	}).Info("Application started")

	// Start gateway
	var edgeProxy string
	if viper.GetString("proxy.host") != "" && viper.GetInt("proxy.port") != 0 {
		edgeProxy = fmt.Sprintf("%s:%d", viper.GetString("proxy.host"), viper.GetInt("proxy.port"))
	}
	gateway.New(
		gateway.KrakendParams{
			Host:      viper.GetString("http.host"),
			Port:      viper.GetInt("http.port"),
			EDGEProxy: edgeProxy,
			LogLevel:  viper.GetString("logging.level"),
			JWKURL:    viper.GetString("JWKURL"),
		},
		mongodb.NewPolicyRepo(db))
}

func init() {
	cmd.SetVersionTemplate(fmt.Sprintf("Version: %s\n", version))

	cobra.OnInitialize(initConfig)

	/*
		Set command flags
	*/

	/*

		Common params

	*/
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file path")

	cmd.PersistentFlags().String("c.logging.output", "STDOUT", "Logging output")
	cmd.PersistentFlags().String("c.logging.level", "INFO", "Logging level")
	cmd.PersistentFlags().String("c.logging.format", "TEXT", "Logging format: TEXT or JSON")

	cmd.PersistentFlags().String("c.configDB.host", "", "DB host")
	cmd.PersistentFlags().Int("c.configDB.port", 27017, "DB port")
	cmd.PersistentFlags().String("c.configDB.login", "", "DB login")
	cmd.PersistentFlags().String("c.configDB.database", "apigw", "DB database name")
	cmd.PersistentFlags().String("c.configDB.password", "", "DB password")

	cmd.PersistentFlags().String("c.eventDB.host", "", "EventDB host")
	cmd.PersistentFlags().Int("c.eventDB.port", 8086, "EventDB port")
	cmd.PersistentFlags().String("c.eventDB.login", "", "EventDB login")
	cmd.PersistentFlags().String("c.eventDB.password", "", "EventDB password")
	cmd.PersistentFlags().String("c.eventDB.database", "auth", "EventDB database name")
	cmd.PersistentFlags().String("c.eventDB.retention", "", "EventDB retention policy")

	cmd.PersistentFlags().String("c.sentryDSN", "", "Sentry DSN")

	cmd.PersistentFlags().Int("c.prometheus.port", 0, "Prometheus port")
	cmd.PersistentFlags().String("c.prometheus.path", "/metrics", "Prometheus path")

	/*

		Gateway params

	*/

	cmd.PersistentFlags().String("c.JWKURL", "", "JWT key URL")

	cmd.PersistentFlags().Int("c.http.host", 0, "Gateway host")
	cmd.PersistentFlags().Int("c.http.port", 0, "Gateway port")

	cmd.PersistentFlags().String("c.proxy.host", "", "Proxy host")
	cmd.PersistentFlags().Int("c.proxy.port", 0, "Proxy port")

	/*
		Bind command flags to config variables
	*/

	for _, parameter := range []string{
		"logging.output",
		"logging.level",
		"logging.format",

		"configDB.host",
		"configDB.port",
		"configDB.login",
		"configDB.password",
		"configDB.database",

		"eventDB.host",
		"eventDB.port",
		"eventDB.login",
		"eventDB.password",
		"eventDB.database",
		"eventDB.retention",

		"sentryDSN",

		"prometheus.port",
		"prometheus.path",

		"JWKURL",

		"http.host",
		"http.port",

		"proxy.host",
		"proxy.port",
	} {
		if err := viper.BindPFlag(parameter, cmd.PersistentFlags().Lookup("c."+parameter)); err != nil {
			log.WithFields(log.Fields{"flag": parameter, "error": err}).Fatal("Can't bind command flag")
		}
	}
}

func initConfig() {
	// Use config file from the flag if provided.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Can't read config:", err)
		}
	}

	viper.SetEnvPrefix("APIGW")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func checkMandatoryParams() {
	var mandatory = []string{
		"configDB.host",
		"eventDB.host",

		"JWKURL",
		"http.host",
		"http.port",
	}
	var missing []string

	for _, param := range mandatory {
		if viper.Get(param) == "" || viper.Get(param) == 0 {
			missing = append(missing, param)
		}
	}

	if len(missing) != 0 {
		log.WithField("missed", missing).Fatal("Missed mandatory params. Use --help flag or config")
	}
}
