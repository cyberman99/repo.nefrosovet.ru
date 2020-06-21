package tool

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	Debug                 bool = false
	DefaultHandlerResults []*SubscribeResult
)

type ExecOptions struct {
	Broker               string
	Qos                  byte
	Retain               bool
	Topic                string
	Username             string
	Password             string
	CertConfig           CertConfig
	ClientNum            int
	Count                int
	UseDefaultHandler    bool
	PreTime              int
	IntervalTime         int
	TargetMPS            float64
	ReplaceValueWithID string
}

type CertConfig interface{}

type ServerCertConfig struct {
	CertConfig
	ServerCertFile string
}

type ClientCertConfig struct {
	CertConfig
	RootCAFile     string
	ClientCertFile string
	ClientKeyFile  string
}

func CreateServerTlsConfig(serverCertFile string) *tls.Config {
	certpool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(serverCertFile)
	if err == nil {
		certpool.AppendCertsFromPEM(pem)
	}

	return &tls.Config{
		RootCAs: certpool,
	}
}

func CreateClientTlsConfig(rootCAFile string, clientCertFile string, clientKeyFile string) *tls.Config {
	certpool := x509.NewCertPool()
	rootCA, err := ioutil.ReadFile(rootCAFile)
	if err == nil {
		certpool.AppendCertsFromPEM(rootCA)
	}

	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		panic(err)
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}

	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
}

func Execute(
	exec func(clients []MQTT.Client, opts ExecOptions, param ...string) int,
	opts ExecOptions,
	message string) error {

	DefaultHandlerResults = make([]*SubscribeResult, opts.ClientNum)

	clients := make([]MQTT.Client, opts.ClientNum)
	hasErr := false
	for i := 0; i < opts.ClientNum; i++ {
		client := Connect(i, opts)
		if client == nil {
			hasErr = true
			break
		}
		clients[i] = client
	}

	if hasErr {
		for i := 0; i < len(clients); i++ {
			client := clients[i]
			if client != nil {
				Disconnect(client)
			}
		}
		return errors.New("can't connect to mqtt")
	}

	time.Sleep(time.Duration(opts.PreTime) * time.Millisecond)

	fmt.Printf("%s Start benchmark\n", time.Now())

	startTime := time.Now()
	totalCount := exec(clients, opts, message)
	endTime := time.Now()

	fmt.Printf("%s End benchmark\n", time.Now())

	AsyncDisconnect(clients)

	duration := (endTime.Sub(startTime)).Nanoseconds() / int64(1000000) // nanosecond -> millisecond
	throughput := float64(totalCount) / float64(duration) * 1000        // messages/sec
	fmt.Printf("\nResult : broker=%s, clients=%d, totalCount=%d, duration=%dms, throughput=%.2fmessages/sec\n",
		opts.Broker, opts.ClientNum, totalCount, duration, throughput)
	if throughput < opts.TargetMPS {
		return fmt.Errorf("low throughput: %f", throughput)

	}
	return nil
}

func PublishAllClient(clients []MQTT.Client, opts ExecOptions, param ...string) int {
	message := param[0]

	wg := new(sync.WaitGroup)
	//mx := new(sync.Mutex)

	totalCount := 0
	for id := 0; id < len(clients); id++ {
		wg.Add(1)
		client := clients[id]

		go func(clientId int) {
			defer wg.Done()

			for index := 0; index < opts.Count; index++ {
				topic := opts.Topic

				if Debug {
					fmt.Printf("Publish : id=%d, count=%d, topic=%s\n", clientId, index, topic)
				}

				if strings.Contains(message, opts.ReplaceValueWithID) && opts.ReplaceValueWithID != "" {
					txID := strconv.Itoa(id)+ "-"+strconv.Itoa(index)+"-"+strconv.Itoa(rand.Int())
					msgWithUUID := strings.Replace(message, opts.ReplaceValueWithID, txID, 1)
					fmt.Println(msgWithUUID)
					Publish(client, topic, opts.Qos, opts.Retain, msgWithUUID)
				} else {
					Publish(client, topic, opts.Qos, opts.Retain, message)
				}
				totalCount++

				if opts.IntervalTime > 0 {
					time.Sleep(time.Duration(opts.IntervalTime) * time.Millisecond)
				}
			}
		}(id)
	}

	wg.Wait()

	return totalCount
}

func Publish(client MQTT.Client, topic string, qos byte, retain bool, message string) {
	token := client.Publish(topic, qos, retain, message)

	if token.Wait() && token.Error() != nil {
		fmt.Printf("Publish error: %s\n", token.Error())
	}
}

func SubscribeAllClient(clients []MQTT.Client, opts ExecOptions, param ...string) int {
	wg := new(sync.WaitGroup)

	results := make([]*SubscribeResult, len(clients))
	for id := 0; id < len(clients); id++ {
		wg.Add(1)

		client := clients[id]
		topic := fmt.Sprintf(opts.Topic+"/%d", id)

		results[id] = Subscribe(client, topic, opts.Qos)

		if opts.UseDefaultHandler == true {
			results[id] = DefaultHandlerResults[id]
		}

		go func(clientId int) {
			defer wg.Done()

			var loop int = 0
			for results[clientId].Count < opts.Count {
				loop++

				if Debug {
					fmt.Printf("Subscribe : id=%d, count=%d, topic=%s\n", clientId, results[clientId].Count, topic)
				}

				if opts.IntervalTime > 0 {
					time.Sleep(time.Duration(opts.IntervalTime) * time.Millisecond)
				} else {
					time.Sleep(1000 * time.Nanosecond)
				}

				if loop >= opts.Count*100 {
					panic("Subscribe error : Not finished in the max count. It may not be received the message.")
				}
			}
		}(id)
	}

	wg.Wait()

	totalCount := 0
	for id := 0; id < len(results); id++ {
		totalCount += results[id].Count
	}

	return totalCount
}

type SubscribeResult struct {
	Count int
}

func Subscribe(client MQTT.Client, topic string, qos byte) *SubscribeResult {
	var result *SubscribeResult = &SubscribeResult{}
	result.Count = 0

	var handler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		result.Count++
		if Debug {
			fmt.Printf("Received message : topic=%s, message=%s\n", msg.Topic(), msg.Payload())
		}
	}

	token := client.Subscribe(topic, qos, handler)

	if token.Wait() && token.Error() != nil {
		fmt.Printf("Subscribe error: %s\n", token.Error())
	}

	return result
}

func Connect(id int, execOpts ExecOptions) MQTT.Client {

	pid := strconv.FormatInt(int64(os.Getpid()), 16)
	clientId := fmt.Sprintf("mqttbench%s-%d", pid, id)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(execOpts.Broker)
	opts.SetClientID(clientId)

	if execOpts.Username != "" {
		opts.SetUsername(execOpts.Username)
	}
	if execOpts.Password != "" {
		opts.SetPassword(execOpts.Password)
	}

	certConfig := execOpts.CertConfig
	switch c := certConfig.(type) {
	case ServerCertConfig:
		tlsConfig := CreateServerTlsConfig(c.ServerCertFile)
		opts.SetTLSConfig(tlsConfig)
	case ClientCertConfig:
		tlsConfig := CreateClientTlsConfig(c.RootCAFile, c.ClientCertFile, c.ClientKeyFile)
		opts.SetTLSConfig(tlsConfig)
	default:
		// do nothing.
	}

	if execOpts.UseDefaultHandler == true {
		var result *SubscribeResult = &SubscribeResult{}
		result.Count = 0

		var handler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
			result.Count++
			if Debug {
				fmt.Printf("Received at defaultHandler : topic=%s, message=%s\n", msg.Topic(), msg.Payload())
			}
		}
		opts.SetDefaultPublishHandler(handler)

		DefaultHandlerResults[id] = result
	}

	client := MQTT.NewClient(opts)
	token := client.Connect()

	if token.Wait() && token.Error() != nil {
		fmt.Printf("Connect error: %s\n", token.Error())
		return nil
	}

	return client
}

func AsyncDisconnect(clients []MQTT.Client) {
	wg := new(sync.WaitGroup)

	for _, client := range clients {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Disconnect(client)
		}()
	}

	wg.Wait()
}

func Disconnect(client MQTT.Client) {
	client.Disconnect(10)
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}
