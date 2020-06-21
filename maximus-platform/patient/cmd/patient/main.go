package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-openapi/loads"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"repo.nefrosovet.ru/maximus-platform/patient/api/handlers"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi"
	"repo.nefrosovet.ru/maximus-platform/patient/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/patient/client/profile"
	"repo.nefrosovet.ru/maximus-platform/patient/db"
	"repo.nefrosovet.ru/maximus-platform/patient/db/migrator"
	"repo.nefrosovet.ru/maximus-platform/patient/db/sqlc"
	"repo.nefrosovet.ru/maximus-platform/patient/logger"
)

var (
	version = "No Version Provided"
	cfgFile string

	mandatoryParams = []string{
		"http.host",
		"http.port",
	}

	store db.Storer
	l     logger.Logger
)

var rootCmd = &cobra.Command{
	Short:   "patient service",
	Long:    `Just use it`,
	Version: version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		l = logger.NewLogger(
			viper.GetString("logging.output"),
			viper.GetString("logging.level"),
			viper.GetString("logging.format"),
		)

		initSentry(l)

		store, err = db.SetupDB(
			db.DBConfig{
				Login:    viper.GetString("db.login"),
				Pass:     viper.GetString("db.password"),
				Host:     viper.GetString("db.host"),
				Database: viper.GetString("db.database"),
				Name:     viper.GetString("db.app.name"),
				MaxConn:  viper.GetInt("db.maxopenconns"),
				MaxIdle:  viper.GetInt("db.maxidleconns"),
				Port:     viper.GetInt("db.port"),
				MaxLife:  viper.GetDuration("db.maxlife"),
				SSL:      viper.GetBool("db.ssl"),
			},
		)
		if err != nil {
			l.Core().Debug(err)
			l.Core().Fatal("cockroach",
				viper.GetString("db.host"),
				strconv.Itoa(viper.GetInt("db.port")),
				logger.COREFAILED)
		}

	},
	PreRun: func(cmd *cobra.Command, args []string) {
		var isFailed bool
		for _, param := range mandatoryParams {
			if viper.Get(param) == "" || viper.Get(param) == 0 {
				println("missed mandatory param: ", param)
				isFailed = true
			}
		}
		if isFailed {
			println("Use --help flag or config")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			l.Core().Debug(err)
			l.Core().Fatal("", "", "", logger.COREFAILED)
		}

		api := operations.NewPatientWPAPI(swaggerSpec)
		server := restapi.NewServer(api)
		defer server.Shutdown()

		profileUri := viper.GetString("profile.host") + ":" + viper.GetString("profile.port")
		profileClient, err := profile.NewClientWithResponses(profileUri)
		if err != nil {
			l.Core().Debug(err)
			l.Core().Fatal("", "", "", logger.COREFAILED)
			return
		}

		server.SetHandler(handlers.ConfigureAPI(api, l.Api(), store, profileClient, version, sqlc.New(store.InnerDB())))

		server.Host = viper.GetString("http.host")
		server.Port = viper.GetInt("http.port")

		l.Core().Info("", "", "", version, logger.CORESTARTED)
		if err := server.Serve(); err != nil {
			l.Core().Debug(err)
			l.Core().Fatal("", "", "", logger.COREFAILED)
		}
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		l.Core().Debug("database closed")
		return store.Close()
	},
}

var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "migration service",
	Long:    `Just use it`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		err := migrator.ApplyMigrations(
			store,
			viper.GetBool("up"),
			viper.GetBool("down"),
			viper.GetBool("flush"),
			viper.GetString("migrations.path"),
		)
		if err != nil {
			l.Core().Debug(err)
			l.Core().Fatal("cockroach",
				viper.GetString("db.host"),
				strconv.Itoa(viper.GetInt("db.port")),
				logger.COREFAILED)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.SetVersionTemplate(`{{printf "Version: %s" .Version}}`)
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file path")

	rootCmd.PersistentFlags().String("db.login", "", "db host")
	rootCmd.PersistentFlags().String("db.password", "", "db password")
	rootCmd.PersistentFlags().String("db.host", "", "db host")
	rootCmd.PersistentFlags().String("db.database", "", "db name")
	rootCmd.PersistentFlags().String("db.app.name", "", "db app name")
	rootCmd.PersistentFlags().Int("db.maxopenconns", 0, "db open connections maximum")
	rootCmd.PersistentFlags().Int("db.maxidleconns", -1, "db idle connections maximum")
	rootCmd.PersistentFlags().Int("db.port", 5432, "db port")
	rootCmd.PersistentFlags().Duration("db.maxlife", time.Duration(0), "db connection life maximum")
	rootCmd.PersistentFlags().Bool("db.ssl", false, "db security")

	rootCmd.PersistentFlags().String("logging.output", "STDOUT", "Logging output")
	rootCmd.PersistentFlags().String("logging.level", "INFO", "Logging level")
	rootCmd.PersistentFlags().String("logging.format", "TEXT", "Logging format: TEXT or JSON")

	rootCmd.PersistentFlags().String("sentryDSN", "", "Sentry DSN")

	rootCmd.PersistentFlags().String("http.host", "", "API host")
	rootCmd.PersistentFlags().Int("http.port", 0, "API port")

	rootCmd.PersistentFlags().String("profile.host", "", "Profile service host")
	rootCmd.PersistentFlags().String("profile.port", "80", "Profile service port")

	migrateCmd.PersistentFlags().Bool("up", false, "to latest migration")
	migrateCmd.PersistentFlags().Bool("down", false, "to previous migration")
	migrateCmd.PersistentFlags().Bool("flush", false, "reset migrations counter")
	migrateCmd.PersistentFlags().String("migrations.path", "", "url to migrations")

	viper.BindPFlags(rootCmd.PersistentFlags())
	viper.BindPFlags(migrateCmd.PersistentFlags())

	rootCmd.AddCommand(migrateCmd)
}

func initConfig() {
	// Use config file from the flag if provided.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
	}

	viper.SetEnvPrefix("PATIENT")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func initSentry(l logger.Logger) {
	if dsn := viper.GetString("sentryDSN"); dsn != "" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: dsn,
		})
		if err != nil {
			l.Core().Debug("Sentry initialization error:", err)
			l.Core().Fatal("", "", "", logger.COREFAILED)
		}
	}
}
