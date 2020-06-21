// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/proxy/broker"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/influx"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod/repos"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
)

var (
	version    = "No Version Provided"
	logQueue   = 10
	logWorkers = 10
	cfgFile    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("version") {
			fmt.Println(cmd.Version)
			os.Exit(0)
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		l := logger.New(
			logWorkers,
			logQueue,
			viper.GetString("logging.level"),
			viper.GetString("logging.output"),
			viper.GetString("logging.format"),
		)
		l.Core().Debug("debug logging is turned on")

		db := mongoConnect(l)
		defer db.Close()

		repos.NewClientRepo(db, map[string]interface{}{ // TODO idk why viper.GetMapString is not working
			"subClientID": viper.GetString("mq.subClientID"),
			"pubClientID": viper.GetString("mq.pubClientID"),
			"host":        viper.GetString("mq.host"),
			"port":        viper.GetInt("mq.port"),
			"login":       viper.GetString("mq.login"),
			"password":    viper.GetString("mq.password"),
			"subscribe":   viper.GetStringSlice("mq.subscribe"),
			"publish":     viper.GetStringSlice("mq.publish"),
		})

		infCli, err := influx.ConnectHTTP(
			viper.GetString("eventdb.host"),
			viper.GetInt("eventdb.port"),
			viper.GetString("eventdb.login"),
			viper.GetString("eventdb.password"),
			viper.GetString("eventdb.database"),
			viper.GetString("eventdb.retention"),
		)
		if err != nil {
			l.Core().Fatal(logger.COREEVENTDB, viper.GetString("eventdb.host"),
				viper.GetString("eventdb.port"), err.Error(), logger.COREFAILED)
		}
		defer infCli.Close()

		l.Core().Info(logger.COREEVENTDB, viper.GetString("eventdb.host"), viper.GetString("eventdb.port"),
			"", logger.CORECONNECTED)

		mqtt.ERROR = l

		br, err := broker.NewMQTTClient(
			l,
			viper.GetString("mq.subClientID"),
			viper.GetString("mq.pubClientID"),
			viper.GetString("mq.host"),
			viper.GetInt("mq.port"),
			viper.GetString("mq.login"),
			viper.GetString("mq.password"),
		)
		if err != nil {
			l.Core().Fatal(logger.COREMQ, viper.GetString("mq.host"), viper.GetString("mq.port"), err.Error(),
				logger.COREFAILED)
		}
		defer br.Close()
		l.Core().Info(logger.COREMQ, viper.GetString("mq.host"),
			viper.GetString("mq.port"),
			version, logger.CORECONNECTED)

		h := broker.NewHandler(l, db, infCli, br)
		err = br.Subscribe(viper.GetStringSlice("mq.subscribe"), h.RouteMessage)
		if err != nil {
			l.Core().Fatal(logger.COREMQ, viper.GetString("mq.host"), viper.GetString("mq.port"),
				err.Error(), logger.COREFAILED)
		}
		l.Core().Info("", "", "", version, logger.CORESTARTED)

		<-c
	},
}

func mongoConnect(lg logger.Logger) (db mongod.Storer) {
	var err error

	db, err = mongod.NewCli(
		viper.GetString("configdb.host"),
		viper.GetInt("configdb.port"),
		viper.GetString("configdb.login"),
		viper.GetString("configdb.password"),
		viper.GetString("configdb.database"),
	)

	if err != nil {
		lg.Core().Fatal(logger.CORECONFIGDB, viper.GetString("configdb.host"),
			viper.GetString("configdb.port"), err.Error(), logger.COREFAILED)
	}

	if err = db.Connect(context.Background()); err != nil {
		lg.Core().Debug(err)
		lg.Core().Fatal(logger.CORECONFIGDB, viper.GetString("configdb.host"),
			viper.GetString("configdb.port"), err.Error(), logger.COREFAILED)
	}
	lg.Core().Info(logger.CORECONFIGDB, viper.GetString("configdb.host"), viper.GetString("configdb.port"),
		"", logger.CORECONNECTED)

	return
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	rootCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		var validErrors error
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
		if !viper.IsSet("mq.host") {
			validErrors = errors.Wrap(validErrors, "required: mq.host")
		}
		if !viper.IsSet("mq.port") {
			validErrors = errors.Wrap(validErrors, "required: mq.port")
		}
		if !viper.IsSet("mq.pubClientID") {
			validErrors = errors.Wrap(validErrors, "required: mq.pubClientID")
		}
		if !viper.IsSet("mq.subClientID") {
			validErrors = errors.Wrap(validErrors, "required: mq.subClientID")
		}
		if !viper.IsSet("mq.login") {
			validErrors = errors.Wrap(validErrors, "required: mq.login")
		}
		if !viper.IsSet("mq.password") {
			validErrors = errors.Wrap(validErrors, "required: mq.password")
		}
		if !viper.IsSet("mq.subscribe") {
			validErrors = errors.Wrap(validErrors, "required: mq.subscribe")
		}
		if !viper.IsSet("mq.publish") {
			validErrors = errors.Wrap(validErrors, "required: mq.publish")
		}

		if validErrors != nil {
			return validErrors
		}

		return nil
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
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
	rootCmd.PersistentFlags().Int("configdb.port", 0, "MongoDb broker database port")
	rootCmd.PersistentFlags().String("configdb.database", "", "MongoDb broker database name")
	// --- END ConfigDb ---

	// --- EventDB ---
	rootCmd.PersistentFlags().String("eventdb.host", "", "Event database host")
	rootCmd.PersistentFlags().Int("eventdb.port", 0, "Event database port")
	rootCmd.PersistentFlags().String("eventdb.protocol", "http", "Event protocol")
	rootCmd.PersistentFlags().String("eventdb.login", "", "Event database login")
	rootCmd.PersistentFlags().String("eventdb.password", "", "Event database password")
	rootCmd.PersistentFlags().String("eventdb.database", "", "Event database name")
	rootCmd.PersistentFlags().String("eventdb.retention", "", "Event keeping time")
	// --- END EventDB ---

	// --- Mq ---
	rootCmd.PersistentFlags().String("mq.host", "", "Broker host")
	rootCmd.PersistentFlags().Int("mq.port", 0, "Broker port")
	rootCmd.PersistentFlags().String("mq.pubClientID", "", "Broker clientId")
	rootCmd.PersistentFlags().String("mq.subClientID", "", "Broker clientId")
	rootCmd.PersistentFlags().String("mq.login", "", "Broker login")
	rootCmd.PersistentFlags().String("mq.password", "", "Broker password")
	rootCmd.PersistentFlags().StringSlice("mq.subscribe", []string{}, "Broker subscribes")
	rootCmd.PersistentFlags().StringSlice("mq.publish", []string{}, "Broker publishes")
	// --- END Mq ---

	rootCmd.PersistentFlags().String("logging.level", "", "")
	rootCmd.PersistentFlags().String("logging.output", "", "")
	rootCmd.PersistentFlags().String("logging.format", "", "")

	rootCmd.PersistentFlags().String("prometheus.path", "", "")
	rootCmd.PersistentFlags().Int("prometheus.port", 0, "")

	rootCmd.PersistentFlags().String("sentryDSN", "", "Sentry URL")
	rootCmd.PersistentFlags().String("tracing.host", "", "")
	rootCmd.PersistentFlags().Int("tracing.port", 0, "")
	rootCmd.PersistentFlags().String("tracing.serviceName", "", "")

	rootCmd.PersistentFlags().Bool("version", false, "Show version")

	viper.BindPFlags(rootCmd.PersistentFlags())
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
