package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	entsql "github.com/facebookincubator/ent/dialect/sql"
	"github.com/getsentry/sentry-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	apierrors "repo.nefrosovet.ru/libs/oapi-errors"

	"repo.nefrosovet.ru/go-lms/api-video/api"
	"repo.nefrosovet.ru/go-lms/api-video/api/server"
	"repo.nefrosovet.ru/go-lms/api-video/ent"
	"repo.nefrosovet.ru/go-lms/api-video/logger"
)

var (
	version = "No Version Provided"
	cfgFile string

	mandatoryParams = []string{
		"http.host",
		"http.port",
	}

	lg logger.Logger

	sentryHub *sentry.Hub
)

var rootCmd = &cobra.Command{
	Short:   "api-video service",
	Long:    `Just use it`,
	Version: version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		lg = logger.NewLogger(
			viper.GetString("logging.output"),
			viper.GetString("logging.level"),
			viper.GetString("logging.format"),
		)
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
		setupSentry()
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

		db, err := sql.Open(
			"mysql",
			fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?parseTime=true",
				viper.GetString("db.login"),
				viper.GetString("db.password"),
				viper.GetString("db.host"),
				viper.GetString("db.port"),
				viper.GetString("db.database"),
			))
		if err != nil {
			lg.Core().Debug("Database client connection error:", err)
			lg.Core().Fatal("", viper.GetString("http.host"), viper.GetString("http.port"),
				logger.COREFAILED)
		}

		db.SetMaxOpenConns(viper.GetInt("db.maxopenconns"))
		db.SetMaxIdleConns(viper.GetInt("db.maxidleconns"))
		db.SetConnMaxLifetime(viper.GetDuration("db.maxlife"))

		client := ent.NewClient(ent.Driver(entsql.OpenDB("mysql", db)))

		if err := client.Schema.Create(context.Background()); err != nil {
			lg.Core().Debug("Database migration error:", err)
			lg.Core().Fatal("", viper.GetString("http.host"), viper.GetString("http.port"),
				logger.COREFAILED)
		}

		srv := server.New(version, client, lg)
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
	viper.BindPFlags(rootCmd.PersistentFlags())
}

func initConfig() {
	// Use config file from the flag if provided.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
	}

	viper.SetEnvPrefix("API-VIDEO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
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
