package broker

import (
	"encoding/json"
	"errors"
	"strconv"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/logger"
)

var (
	ErrMQTTDisconnected = errors.New("MQTT disconnected")
)

type MessageBody struct {
	Payload       interface{} `json:"payload"`
	TransactionID string      `json:"transactionID"`
}

func (m *MessageBody) DoUnmarshal(pl []byte) error {
	if err := json.Unmarshal(pl, m); err != nil {
		return err
	}

	return nil
}

type Messenger interface {
	Publish(topic string, qos byte, payload interface{})
	AsyncPublish(topic string, qos byte, payload interface{})
	Subscribe(subscribe []string, h func(message mqtt.Message)) error
	Close()
}

type mqttCli struct {
	readConnect  mqtt.Client
	writeConnect mqtt.Client

	wg *sync.WaitGroup
	l  logger.Logger
}

func NewMQTTClient(
	l logger.Logger,
	clientId string,
	clientIdW string,
	host string,
	port int,
	login string,
	password string) (Messenger, error) {
	var (
		readConnect  mqtt.Client
		writeConnect mqtt.Client
		err          error
	)

	for i := 1; i < 4; i++ {
		readConnect, err = makeConnect(clientId, host, port, login, password)
		if err == nil {
			break
		}
		l.Debugln("Trying to connect with broker. Try №: ", i, err)
		time.Sleep(7 * time.Second)
	}
	if err != nil {
		return nil, err
	}

	writeConnect, err = makeConnect(clientIdW, host, port, login, password)
	if err != nil {
		return nil, err
	}
	return &mqttCli{
		l:            l,
		wg:           new(sync.WaitGroup),
		readConnect:  readConnect,
		writeConnect: writeConnect,
	}, nil
}

func makeConnect(clientId string, host string, port int, login string, password string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts = opts.AddBroker("tcp://" + host + ":" + strconv.Itoa(port))
	opts.ClientID = clientId
	opts.Username = login
	opts.Password = password
	opts.SetConnectTimeout(5 * time.Second)
	opts.SetMessageChannelDepth(1000)
	mqttClient := mqtt.NewClient(opts)

	token := mqttClient.Connect()
	token.Wait()
	if token.Error() != nil {
		return nil, token.Error()
	}
	if !mqttClient.IsConnectionOpen() {
		return nil, ErrMQTTDisconnected
	}

	return mqttClient, nil
}

func (c *mqttCli) Publish(topic string, qos byte, payload interface{}) {
	token := c.writeConnect.Publish(topic, qos, false, payload)
	token.Wait()
	if token.Error() != nil {
		c.l.Debugf("publish message error: %v", token.Error())
	}
}

func (c *mqttCli) AsyncPublish(topic string, qos byte, payload interface{}) {
	if !c.writeConnect.IsConnected() {
		return
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		token := c.writeConnect.Publish(topic, qos, false, payload)
		token.Wait()
		if token.Error() != nil {
			c.l.Debugf("async publish message error: %v", token.Error())
		}
	}()
}

func (c *mqttCli) Subscribe(subscribe []string, h func(msg mqtt.Message)) error {
	if !c.readConnect.IsConnected() {
		return ErrMQTTDisconnected
	}
	topicMap := make(map[string]byte)
	for _, subChan := range subscribe {
		topicMap[subChan] = 0
	}

	callback := func(_ mqtt.Client, msg mqtt.Message) {
		h(msg)
	}

	token := c.readConnect.SubscribeMultiple(topicMap, callback)
	token.Wait()
	return token.Error()
}

func (c *mqttCli) Close() {
	c.readConnect.Disconnect(0)
	c.l.Debugln("Stop receiving messages. Sending last messages")
	c.wg.Wait()
	c.l.Debugln("All messages sent. Shutting down")
	c.writeConnect.Disconnect(0)
	if c.readConnect.IsConnected() || c.writeConnect.IsConnected() {
		c.l.Fatal("mqtt still connected. Kill all processes manually")
	}
}
