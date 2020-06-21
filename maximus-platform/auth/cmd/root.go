package cmd

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string

	skipMandatoryParams bool
	startFn             func()
)

var cmd = &cobra.Command{
	Use:     "auth",
	Short:   "auth service",
	Long:    `Just use it`,
	Version: "No Version Provided",
	Run: func(cmd *cobra.Command, args []string) {
		if !skipMandatoryParams {
			checkMandatoryParams()
		}

		if startFn != nil {
			startFn()
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	/*
		Set command flags
	*/

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Config file path")

	cmd.PersistentFlags().String("c.http.host", "", "API host")
	cmd.PersistentFlags().Int("c.http.port", 0, "API port")

	cmd.PersistentFlags().String("c.logging.output", "STDOUT", "Logging output")
	cmd.PersistentFlags().String("c.logging.level", "INFO", "Logging level")
	cmd.PersistentFlags().String("c.logging.format", "TEXT", "Logging format: TEXT or JSON")

	cmd.PersistentFlags().String("c.configDB.host", "", "DB host")
	cmd.PersistentFlags().Int("c.configDB.port", 27017, "DB port")
	cmd.PersistentFlags().String("c.configDB.login", "", "DB login")
	cmd.PersistentFlags().String("c.configDB.database", "auth", "DB database name")
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

	cmd.PersistentFlags().String("c.adminPassword", "", "Admin default password")
	cmd.PersistentFlags().String("c.tokenSecret", "", "JWT secret")
	cmd.PersistentFlags().Int("c.ttl.accessToken", 60*10, "AccessToken expire duration (seconds)")
	cmd.PersistentFlags().Int("c.ttl.refreshToken", 60*60*24, "RefreshToken expire duration (seconds)")

	cmd.PersistentFlags().String("c.index.http.host", "", "Index service host")
	cmd.PersistentFlags().String("c.index.http.path", "", "Index service path")

	cmd.PersistentFlags().String("c.oAuth2.esia.mnemonics", "", "OAuth2 esia mnemonics")
	cmd.PersistentFlags().String("c.oAuth2.esia.privateKeyPath", "", "OAuth2 esia private key path")
	cmd.PersistentFlags().String("c.oAuth2.esia.certPath", "", "OAuth2 esia certificate path")
	cmd.PersistentFlags().String("c.oAuth2.esia.redirectURI", "", "OAuth2 esia redirect uri")

	/*
		Bind command flags to config variables
	*/
	viper.BindPFlag("http.host", cmd.PersistentFlags().Lookup("c.http.host"))
	viper.BindPFlag("http.port", cmd.PersistentFlags().Lookup("c.http.port"))

	viper.BindPFlag("logging.output", cmd.PersistentFlags().Lookup("c.logging.output"))
	viper.BindPFlag("logging.level", cmd.PersistentFlags().Lookup("c.logging.level"))
	viper.BindPFlag("logging.format", cmd.PersistentFlags().Lookup("c.logging.format"))

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

	viper.BindPFlag("sentryDSN", cmd.PersistentFlags().Lookup("c.sentryDSN"))

	viper.BindPFlag("prometheus.port", cmd.PersistentFlags().Lookup("c.prometheus.port"))
	viper.BindPFlag("prometheus.path", cmd.PersistentFlags().Lookup("c.prometheus.path"))

	viper.BindPFlag("adminPassword", cmd.PersistentFlags().Lookup("c.adminPassword"))
	viper.BindPFlag("tokenSecret", cmd.PersistentFlags().Lookup("c.tokenSecret"))
	viper.BindPFlag("ttl.accessToken", cmd.PersistentFlags().Lookup("c.ttl.accessToken"))
	viper.BindPFlag("ttl.refreshToken", cmd.PersistentFlags().Lookup("c.ttl.refreshToken"))

	viper.BindPFlag("index.http.host", cmd.PersistentFlags().Lookup("c.index.http.host"))
	viper.BindPFlag("index.http.path", cmd.PersistentFlags().Lookup("c.index.http.path"))

	// OAuth2
	viper.BindPFlag("oAuth2.google.clientID", cmd.PersistentFlags().Lookup("oAuth2.google.clientID"))
	viper.BindPFlag("oAuth2.google.clientSecret", cmd.PersistentFlags().Lookup("c.oAuth2.google.clientSecret"))
	viper.BindPFlag("oAuth2.google.redirectURI", cmd.PersistentFlags().Lookup("c.oAuth2.google.redirectURI"))

	viper.BindPFlag("oAuth2.esia.mnemonics", cmd.PersistentFlags().Lookup("c.oAuth2.esia.mnemonics"))
	viper.BindPFlag("oAuth2.esia.privateKeyPath", cmd.PersistentFlags().Lookup("c.oAuth2.esia.privateKeyPath"))
	viper.BindPFlag("oAuth2.esia.certPath", cmd.PersistentFlags().Lookup("c.oAuth2.esia.certPath"))
	viper.BindPFlag("oAuth2.esia.redirectURI", cmd.PersistentFlags().Lookup("c.oAuth2.esia.redirectURI"))
}

func checkMandatoryParams() {
	var mandatory = []string{
		"http.host",
		"http.port",
		"configDB.host",
		"eventDB.host",
		"adminPassword",
		"tokenSecret",
		"index.http.host",
	}
	var missing []string

	for _, param := range mandatory {
		if viper.Get(param) == "" || viper.Get(param) == 0 {
			missing = append(missing, param)
		}
	}

	if len(missing) != 0 {
		log.WithField("missed", missing).
			Fatal("Missed mandatory params. Use --help flag or config")
	}
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)

		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).
				Fatal("Config reading error")
		}
	}

	viper.SetEnvPrefix("AUTH")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func SetVersion(version string) {
	cmd.Version = version
	cmd.SetVersionTemplate(fmt.Sprintf("Version: %s\n", version))
}

func SetConfigFile(cfgFile string) {
	configFile = cfgFile
}

func Execute(start func(), skipMandatory ...bool) {
	skipMandatoryParams = len(skipMandatory) > 0 && skipMandatory[0]
	startFn = start

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
