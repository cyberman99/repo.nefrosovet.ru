package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"repo.nefrosovet.ru/libs/mqtt-benchmark/tool"
	"strings"
)

const defaultTargetMPS = 10000

var rootCmd = &cobra.Command{
	Use:   "mqtt-bencmark",
	Short: "Benchmark test tool for mqtt",
	Long:  ``,
	//Uncomment the following line if your bare application
	//has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var (
			message string
			err     error
		)

		broker := viper.GetString("broker")

		action := viper.GetString("action")
		qos := viper.GetInt("qos")
		retain := viper.GetBool("retain")
		topic := viper.GetString("topic")
		username := viper.GetString("username")
		password := viper.GetString("password")
		tls := viper.GetString("tls")
		clients := viper.GetInt("clients")
		count := viper.GetInt("count")
		useDefaultHandler := viper.GetBool("support-unknown-received")
		preTime := viper.GetInt("pretime")
		intervalTime := viper.GetInt("intervaltime")
		debug := viper.GetBool("debug")
		path := viper.GetString("filepath")
		if path != "" {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}
			message = string(data)
		}

		// validate "broker"
		if broker == "" || broker == "tcp://{host}:{port}" {
			log.Fatalf("Invalid argument : --broker -> %s\n", broker)
		}

		// validate "action"
		var method string = ""
		if action == "p" || action == "pub" {
			method = "pub"
		} else if action == "s" || action == "sub" {
			method = "sub"
		}

		if method != "pub" && method != "sub" {
			log.Fatalf("Invalid argument : --action -> %s\n", action)
		}

		// parse TLS mode
		var certConfig tool.CertConfig = nil
		if tls != "" {
			// nil
			if strings.HasPrefix(tls, "server:") {
				var strArray = strings.Split(tls, "server:")
				serverCertFile := strings.TrimSpace(strArray[1])
				if tool.FileExists(serverCertFile) == false {
					log.Fatalf("File is not found. : certFile -> %s\n", serverCertFile)
				}

				certConfig = tool.ServerCertConfig{
					ServerCertFile: serverCertFile}
			} else if strings.HasPrefix(tls, "client:") {
				var strArray = strings.Split(tls, "client:")
				var configArray = strings.Split(strArray[1], ",")
				rootCAFile := strings.TrimSpace(configArray[0])
				clientCertFile := strings.TrimSpace(configArray[1])
				clientKeyFile := strings.TrimSpace(configArray[2])
				if tool.FileExists(rootCAFile) == false {
					log.Fatalf("File is not found. : rootCAFile -> %s\n", rootCAFile)
				}
				if tool.FileExists(clientCertFile) == false {
					log.Fatalf("File is not found. : clientCertFile -> %s\n", clientCertFile)
				}
				if tool.FileExists(clientKeyFile) == false {
					log.Fatalf("File is not found. : clientKeyFile -> %s\n", clientKeyFile)
				}

				certConfig = tool.ClientCertConfig{
					RootCAFile:     rootCAFile,
					ClientCertFile: clientCertFile,
					ClientKeyFile:  clientKeyFile}
			}
		}
		execOpts := tool.ExecOptions{}
		execOpts.Broker = broker
		execOpts.Qos = byte(qos)
		execOpts.Retain = retain
		execOpts.Topic = topic
		execOpts.Username = username
		execOpts.Password = password
		execOpts.CertConfig = certConfig
		execOpts.ClientNum = clients
		execOpts.Count = count
		execOpts.UseDefaultHandler = useDefaultHandler
		execOpts.PreTime = preTime
		execOpts.IntervalTime = intervalTime
		execOpts.TargetMPS = float64(viper.GetInt("mps"))
		execOpts.ReplaceValueWithID = viper.GetString("replace")

		tool.Debug = debug
		switch method {
		case "pub":
			err = tool.Execute(tool.PublishAllClient, execOpts, message)
		case "sub":
			err = tool.Execute(tool.SubscribeAllClient, execOpts, message)
		}
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	cobra.OnInitialize(func() {
		viper.SetEnvPrefix("BENCHMARK")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv() // read in environment variables that match
	})
	rootCmd.PersistentFlags().String("replace", "", "Choose substring that will be replaced with unique id")
	rootCmd.PersistentFlags().String("broker", "tcp://{host}:{port}", "URI of MQTT broker (required)")
	rootCmd.PersistentFlags().String("action", "p|pub or s|sub", "Publish or Subscribe or Subscribe(with publishing) (required)")
	rootCmd.PersistentFlags().Int("qos", 0, "MQTT QoS(0|1|2)")
	rootCmd.PersistentFlags().Bool("retain", false, "MQTT Retain")
	rootCmd.PersistentFlags().String("topic", "/services/1/IN", "Base topic")
	rootCmd.PersistentFlags().String("username", "", "Username for connecting to the MQTT broker")
	rootCmd.PersistentFlags().String("password", "", "Password for connecting to the MQTT broker")
	rootCmd.PersistentFlags().String("tls", "", "TLS mode. 'server:certFile' or 'client:rootCAFile,clientCertFile,clientKeyFile'")
	rootCmd.PersistentFlags().Int("clients", 10, "Number of clients")
	rootCmd.PersistentFlags().Int("count", 100, "Number of loops per client")
	rootCmd.PersistentFlags().Bool("support-unknown-received", false, "Using default messageHandler for a message that does not match any known subscriptions")
	rootCmd.PersistentFlags().Int("pretime", 300, "Pre wait time (ms)")
	rootCmd.PersistentFlags().Int("intervaltime", 0, "Interval time per message (ms)")
	rootCmd.PersistentFlags().Bool("debug", false, "Debug mode")
	rootCmd.PersistentFlags().String("filepath", "./example/payload.example", "Message filepath")
	rootCmd.PersistentFlags().Int("mps", defaultTargetMPS, "target mps")

	viper.BindPFlags(rootCmd.PersistentFlags())

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
