package main

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/go-openapi/loads"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/recognition/logger"
	"repo.nefrosovet.ru/maximus-platform/recognition/services"
	"strings"
)

var (
	version = "No Version Provided"
	cfgFile string
)

var cmd = &cobra.Command{
	Use:     "rekognition",
	Short:   "rekognition service",
	Long:    `Just use it`,
	Version: version,
	PreRun: func(cmd *cobra.Command, args []string) {
		var mandatoryParams = []string{
			"aws.bucket",
			"aws.accessID",
			"aws.accessSecret",
			"aws.region",
			"http.host",
			"http.port",
		}
		var missing []string

		for _, param := range mandatoryParams {
			if viper.Get(param) == "" || viper.Get(param) == 0 {
				missing = append(missing, param)
			}
		}

		if len(missing) != 0 {
			println("Missed mandatory params: ", missing, ". Use --help flag or config")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.NewLogger(
			setOutput(),
			viper.GetString("logging.level"),
			viper.GetString("logging.format"),
		)
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			l.Core().Debug(err, "fail swagger embedding")
			os.Exit(1)
		}

		api := operations.NewRecognitionAPI(swaggerSpec)
		server := restapi.NewServer(api)
		defer server.Shutdown()

		sim := viper.GetFloat64("similarity")
		var conf = services.Config{
			Bucket:       viper.GetString("aws.bucket"),
			Similarity:   sim,
			AccessID:     viper.GetString("aws.accessID"),
			AccessSecret: viper.GetString("aws.accessSecret"),
			Region:       viper.GetString("aws.region"),
		}

		var servicer services.CloudServicer
		servicer, err = services.NewService(services.AmazonService, &conf, l.Core())
		if err != nil {
			l.Core().Debug(err)
			l.Core().Fatal("service", "", "", logger.COREFAILED)
		}
		l.Core().Info("service", "", "", version, logger.CORECONNECTED)

		if viper.GetString("sentryDSN") != "" {
			err = raven.SetDSN(viper.GetString("sentryDSN"))
			if err != nil {
				l.Core().Debug(err, "failed to start sentry dsn")
				os.Exit(1)
			}
		}

		server.SetHandler(restapi.ConfigureAPI(api, l.Api(sim), servicer, version))

		server.Host = viper.GetString("http.host")
		server.Port = viper.GetInt("http.port")

		l.Core().Info("", "", "", version, logger.CORESTARTED)
		if err := server.Serve(); err != nil {
			l.Core().Debug(err, "failed to start server")
			os.Exit(1)
		}
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	cmd.SetVersionTemplate(fmt.Sprintf("Version: %s\n", version))

	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file path")

	cmd.PersistentFlags().String("http.host", "", "API host")
	cmd.PersistentFlags().Int("http.port", 0, "API port")

	cmd.PersistentFlags().String("aws.bucket", "", "bucket name")
	cmd.PersistentFlags().Float64("aws.similarity", 0, "photo similarity")
	cmd.PersistentFlags().String("aws.accessID", "", "access ID for AWS")
	cmd.PersistentFlags().String("aws.accessSecret", "", "access secret for AWS")
	cmd.PersistentFlags().String("aws.region", "", "region for AWS")

	cmd.PersistentFlags().String("logging.output", "STDOUT", "Logging output")
	cmd.PersistentFlags().String("logging.level", "INFO", "Logging level")
	cmd.PersistentFlags().String("logging.format", "TEXT", "Logging format: TEXT or JSON")

	cmd.PersistentFlags().String("sentryDSN", "", "Sentry DSN")

	// Bind command flags to config variables
	viper.BindPFlags(cmd.PersistentFlags())
}

func initConfig() {
	// Use config file from the flag if provided.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Can't read config:", err)
		}
	}

	viper.SetEnvPrefix("RECOGNITION")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func setOutput() *os.File {
	if viper.GetString("logging.output") != "STDOUT" {
		logFile, err := os.OpenFile(viper.GetString("logging.output"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer func() {
			err = logFile.Close()
			if err != nil {
				panic(err)
			}
		}()
		if err != nil {
			panic(err)
		}
		return logFile

	}
	return os.Stdout
}
