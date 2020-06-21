package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	pops "repo.nefrosovet.ru/libs/Populator"
	"repo.nefrosovet.ru/maximus-platform/patient/db"
	"strings"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Short:   "populator service",
	Long:    `Just use it`,
	Version: "No Version Provided",
	Run: func(cmd *cobra.Command, args []string) {
		store, err := db.SetupDB(
			db.DBConfig{
				Login:    viper.GetString("db.login"),
				Pass:     viper.GetString("db.password"),
				Host:     viper.GetString("db.host"),
				Database: viper.GetString("db.database"),
				Port:     viper.GetInt("db.port"),
				Name:     "populator",
				SSL:      false,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		defer store.Close()
		err = pops.Populate(
			store.InnerDB(),
			viper.GetInt("populator.rows"),
			viper.GetString("populator.path"),
		)
		if err != nil {
			log.Fatal(err)
		}

		var id int32
		err = store.InnerDB().QueryRow(
			`SELECT id FROM appointment WHERE id = $1`,
			int32(rand.Intn(viper.GetInt("populator.rows"))),
		).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file path")

	rootCmd.PersistentFlags().String("db.login", "", "db host")
	rootCmd.PersistentFlags().String("db.password", "", "db password")
	rootCmd.PersistentFlags().String("db.host", "", "db host")
	rootCmd.PersistentFlags().String("db.database", "", "db name")

	rootCmd.PersistentFlags().String("populator.path", "", "path to models")
	rootCmd.PersistentFlags().Int("populator.rows", 1, "rows to insert")

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

	viper.SetEnvPrefix("PATIENT")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
