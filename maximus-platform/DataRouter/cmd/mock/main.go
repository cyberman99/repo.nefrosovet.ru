package main

import (
	"context"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/datarouter/domain"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/cmd/proxy/broker"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/influx"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/db/mongod/repos"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:               "mock proxy",
	PersistentPreRunE: preRun,
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		l := logger.New(
			10,
			1000,
			viper.GetString("logging.level"),
			viper.GetString("logging.output"),
			viper.GetString("logging.format"),
		)
		l.Core().Debug("debug logging is turned on")

		db := mongoConnect(l)
		defer db.Close()

		src := domain.Source{}
		_ = json.Unmarshal([]byte(`{ "and" : [
				  {"<" : [ { "var" : "temp" }, 110 ]},
				  {"==" : [ { "var" : "pie.filling" }, "apple" ] }
				] }`), &src.Payload)
		_ = json.Unmarshal(
			[]byte(`{"and": [{"==": [{"var": "name"}, "services/mock/OUT"]}]}`),
			&src.Topic,
		)

		repo := repos.NewRouteRepo(db)
		rt, err := repo.Set(domain.Route{
			ReplyID: nil,
			Dst: []domain.Destinations{
				{0, "services/test1/IN"},
				{1, "services/test2/IN"},
				{2, "services/test3/IN"},
			},
			Src: src,
		})
		if err != nil {
			log.Fatal(err)
		}
		defer repo.Delete(rt.RouteID)

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
		l.Core().Info(logger.COREMQ, viper.GetString("mq.host"), viper.GetString("mq.port"),
			"", logger.CORECONNECTED)

		h := broker.NewHandler(l, db, infCli, br)
		err = br.Subscribe(viper.GetStringSlice("mq.subscribe"), h.RouteMessage)
		if err != nil {
			l.Core().Fatal(logger.COREMQ, viper.GetString("mq.host"), viper.GetString("mq.port"),
				err.Error(), logger.COREFAILED)
		}
		l.Core().Info("", "", "", "", logger.CORESTARTED)

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

func preRun(cmd *cobra.Command, args []string) error {
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

func init() {
	rootCmd.Version = ""
	cobra.OnInitialize(func() {
		viper.SetEnvPrefix("DATAROUTER")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
	})

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

	viper.BindPFlags(rootCmd.PersistentFlags())
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
