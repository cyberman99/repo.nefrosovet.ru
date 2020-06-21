package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/influx"

	"github.com/getsentry/raven-go"
	"github.com/go-openapi/loads"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
)

var (
	version = "No Version Provided"
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Use:     "DataRouter",
	Short:   "DataRouter service",
	Long:    `Just use it`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("version") {
			fmt.Println(cmd.Version)
			os.Exit(0)
		}
		var err error
		// Logging init
		lg := logger.New(
			10,
			10,
			viper.GetString("logging.level"),
			viper.GetString("logging.output"),
			viper.GetString("logging.format"),
		)
		lg.Debugln("debug logging is turned on")
		restapi.Lg = lg

		{
			var db mongod.Storer

			dbName := viper.GetString("configdb.database")
			dbhost := viper.GetString("configdb.host")
			dbport := viper.GetInt("configdb.port")

			db, err = mongod.NewCli(
				dbhost,
				dbport,
				viper.GetString("configdb.login"),
				viper.GetString("configdb.password"),
				dbName,
			)
			if err != nil {
				lg.Core().Fatal(logger.CORECONFIGDB, viper.GetString("configdb.host"),
					viper.GetString("configdb.port"), err.Error(), logger.COREFAILED)
			}

			defer db.Close()
			err = db.Connect(context.Background())
			if err != nil {
				lg.Core().Debug(err)
				lg.Core().Fatal(logger.CORECONFIGDB, viper.GetString("configdb.host"),
					viper.GetString("configdb.port"), err.Error(), logger.COREFAILED)
			}
			lg.Core().Info(logger.CORECONFIGDB, viper.GetString("configdb.host"),
				viper.GetString("configdb.port"), "", logger.CORECONNECTED)
		}

		infCli, err := influx.ConnectHTTP(
			viper.GetString("eventdb.host"),
			viper.GetInt("eventdb.port"),
			viper.GetString("eventdb.login"),
			viper.GetString("eventdb.password"),
			viper.GetString("eventdb.database"),
			viper.GetString("eventdb.retention"),
		)
		if err != nil {
			lg.Core().Fatal(logger.COREEVENTDB, viper.GetString("eventdb.host"),
				viper.GetString("eventdb.port"), err.Error(), logger.COREFAILED)
		}
		defer infCli.Close()
		lg.Core().Info(logger.COREEVENTDB, viper.GetString("eventdb.host"),
			viper.GetString("eventdb.port"), "", logger.CORECONNECTED)

		// Start JSON API server
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			lg.Fatalln(err)
		}

		api := operations.NewDataRouterAPI(swaggerSpec)
		restapi.Version = version
		server := restapi.NewServer(api)
		defer func() {
			err = server.Shutdown()
		}()

		// Apply server options
		server.Host = viper.GetString("http.host")
		server.Port = viper.GetInt("http.port")

		server.ConfigureAPI()

		if viper.GetInt("prometheus.port") != 0 {
			go func() {
				http.Handle(viper.GetString("prometheus.path"), promhttp.Handler())

				err := http.ListenAndServe(":"+strconv.Itoa(viper.GetInt("prometheus.port")), nil)
				if err != nil {
					log.Fatal(err)
				}
			}()
		}

		lg.Core().Info("", "", "", version, logger.CORESTARTED)

		if viper.GetString("sentryDSN") != "" {
			err := raven.SetDSN(viper.GetString("sentryDSN"))
			if err != nil {
				log.Fatal(err)
			}
		}

		if err := server.Serve(); err != nil {
			lg.Core().Fatal("", "", "", version, logger.COREFAILED)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	rootCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		var validErrors error
		if !viper.IsSet("http.host") {
			validErrors = errors.Wrap(validErrors, "required: http.host")
		}
		if !viper.IsSet("http.port") {
			validErrors = errors.Wrap(validErrors, "required: http.port")
		}
		if !viper.IsSet("configdb.host") {
			validErrors = errors.Wrap(validErrors, "required: configdb.host")
		}
		if !viper.IsSet("configdb.port") {
			validErrors = errors.Wrap(validErrors, "required: configdb.port")
		}
		if !viper.IsSet("configdb.database") {
			validErrors = errors.Wrap(validErrors, "required: configdb.database")
		}
		if !viper.IsSet("eventdb.host") {
			validErrors = errors.Wrap(validErrors, "required: eventdb.host")
		}
		if !viper.IsSet("eventdb.port") {
			validErrors = errors.Wrap(validErrors, "required: eventdb.port")
		}
		if !viper.IsSet("eventdb.database") {
			validErrors = errors.Wrap(validErrors, "required: eventdb.database")
		}
		if validErrors != nil {
			return validErrors
		}

		return nil
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file path")

	// --- ConfigDb ---
	rootCmd.PersistentFlags().String("configdb.host", "", "MongoDb broker database host")
	rootCmd.PersistentFlags().Int("configdb.port", 27017, "MongoDb broker database port")
	rootCmd.PersistentFlags().String("configdb.database", "", "MongoDb broker database name")

	// --- END ConfigDb ---

	// --- EventDB ---
	rootCmd.PersistentFlags().String("eventdb.host", "", "Event database host")
	rootCmd.PersistentFlags().Int("eventdb.port", 0, "Event database port")
	rootCmd.PersistentFlags().String("eventdb.protocol", "http", "Event protocol")
	rootCmd.PersistentFlags().String("eventdb.database", "", "Event database name")
	rootCmd.PersistentFlags().String("eventdb.login", "", "Event database login")
	rootCmd.PersistentFlags().String("eventdb.password", "", "Event database password")
	rootCmd.PersistentFlags().String("eventdb.retention", "", "Event keeping time")
	// --- END EventDB ---

	// --- Metrics ---
	rootCmd.PersistentFlags().String("prometheus.path", "/metrics", "Prometheus path")
	rootCmd.PersistentFlags().Int("prometheus.port", 0, "Prometheus port")
	// --- END Metrics ---

	// --- Http ---
	rootCmd.PersistentFlags().String("http.host", "", "Webserver binding http host")
	rootCmd.PersistentFlags().Int("http.port", 0, "Webserver binding http port")
	// --- END Http ---

	rootCmd.PersistentFlags().String("logging.level", "", "")
	rootCmd.PersistentFlags().String("logging.output", "", "")
	rootCmd.PersistentFlags().String("logging.format", "", "")

	rootCmd.PersistentFlags().String("sentryDSN", "", "Sentry URL")
	rootCmd.PersistentFlags().String("tracing.host", "", "")
	rootCmd.PersistentFlags().Int("tracing.port", 0, "")
	rootCmd.PersistentFlags().String("tracing.serviceName", "", "")

	rootCmd.PersistentFlags().Bool("version", false, "Show version")

	// Bind command flags to config variables
	viper.BindPFlags(rootCmd.PersistentFlags())

	// Set sentry definitions
	dsn := viper.GetString("sentrydsn")
	if dsn != "" {
		raven.SetDSN(dsn)
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

	viper.SetEnvPrefix("DATAROUTER")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func main() {
	Execute(version)
}
