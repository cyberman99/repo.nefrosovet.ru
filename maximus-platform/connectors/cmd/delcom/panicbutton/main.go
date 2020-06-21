package main

import (
	"github.com/getsentry/raven-go"
	"github.com/gofrs/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"repo.nefrosovet.ru/maximus-platform/connectors/broker"
	"repo.nefrosovet.ru/maximus-platform/connectors/connector/delcom/panicbutton"
	"repo.nefrosovet.ru/maximus-platform/connectors/logger"
	"strings"
	"syscall"
)

var (
	version = "No Version Provided"
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:     "panicbutton",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("version") {
			log.Println(cmd.Version)
			os.Exit(0)
			return
		}
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

		btn, err := panicbutton.NewButton(viper.GetString("device.productID"))
		if err != nil {
			log.Fatal(err)
		}
		defer btn.Close()

		var (
			host    = viper.GetString("mq.host")
			portStr = viper.GetString("mq.port")
			topic   = viper.GetString("mq.topic")
			ID      = viper.GetString("ID")
		)

		l := logger.NewLogger(
			viper.GetString("logging.output"),
			viper.GetString("logging.level"),
			viper.GetString("logging.format"),
		)

		br, err := broker.NewMQTTClient(
			viper.GetString("mq.clientId"),
			viper.GetString("mq.host"),
			viper.GetInt("mq.port"),
			viper.GetString("mq.login"),
			viper.GetString("mq.password"),
			byte(viper.GetInt("mq.qos")),
			viper.GetBool("mq.cleansession"),
		)
		if err != nil {
			l.Core().Debug(err)
			l.Core().MQFailed(host, portStr)
		}
		defer br.Close()

		l.Core().MQConnected(host, portStr)

		msgChan, errChan := btn.Listen()

		l.Core().AppStarted(version)

		go func() {
			for msg := range msgChan {
				txID             := uuid.Must(uuid.NewV1())
				msg.Connector.ID = ID

				if err = br.Publish(topic, txID.String(), msg); err != nil {
					l.Event().EventFail(txID.String(), btn.ConnectorType(), ID, err.Error())
				}
				l.Event().EventSuccess(
					txID.String(),
					btn.ConnectorType(),
					ID,
				)
				l.Event().Debug(msg.Data.StatusCode)
			}
		}()

		go func() {
			for err := range errChan {
				l.Event().Debug("BUTTON SENDS ERROR: ", err)
			}
		}()

		<-stop
		l.Core().Debug("CLOSING")
	},
}

func main() {
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

	rootCmd.PersistentFlags().String("ID", "", "Connector ID")

	// --- Mq ---
	rootCmd.PersistentFlags().String("mq.host", "", "Broker host")
	rootCmd.PersistentFlags().Int("mq.port", 0, "Broker port")
	rootCmd.PersistentFlags().String("mq.сlientId", "", "Broker clientId")
	rootCmd.PersistentFlags().String("mq.login", "", "Broker login")
	rootCmd.PersistentFlags().String("mq.password", "", "Broker password")
	rootCmd.PersistentFlags().String("mq.topic", "", "Broker publishes")
	rootCmd.PersistentFlags().Bool("mq.cleansession", true, "clean undelivered")
	rootCmd.PersistentFlags().String("mq.qos", "", "sendin quality")
	// --- END Mq ---

	rootCmd.PersistentFlags().String("device.productID", "", "device ID")

	rootCmd.PersistentFlags().String("logging.level", "", "")
	rootCmd.PersistentFlags().String("logging.output", "", "")
	rootCmd.PersistentFlags().String("logging.format", "", "")

	rootCmd.PersistentFlags().String("prometheus.path", "", "")
	rootCmd.PersistentFlags().Int("prometheus.port", 0, "")

	rootCmd.PersistentFlags().String("sentrydsn", "", "Sentry URL")

	rootCmd.PersistentFlags().Bool("version", false, "Show version")

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

	viper.SetEnvPrefix("CONNECTOR")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
