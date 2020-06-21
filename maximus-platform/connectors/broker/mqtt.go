package broker

import (
	"encoding/json"
	"errors"
	"strconv"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	proto                      = "tcp://"
	defaultTimeout             = 2 * time.Second
	defaultMessageChannelDepth = 100
	defaultConnectTimeout      = 5 * time.Second
)

var (
	ErrInvalidMessage   = errors.New("mqttpubsub: invalid or empty message")
	ErrConnRequired     = errors.New("mqttpubsub: mqtt connection is required")
	ErrStillConnected   = errors.New("mqttpubsub: still connected. Kill all processes manually")
	ErrMQTTDisconnected = errors.New("mqttpubsub: disconnected")
)

type Message struct {
	Payload       interface{} `json:"payload"`
	TransactionID string      `json:"transactionID"`
}

type Broker interface {
	Publish(topic string, txID string, payload interface{}) error
	Close() error
}

type mqttCli struct {
	cli     mqtt.Client
	timeout time.Duration
	qos     byte

	isStopped bool
	wg        *sync.WaitGroup
}

func NewMQTTClient(
	clientId string,
	host string,
	port int,
	login string,
	password string,
	qos byte,
	cleanSession bool,
) (Broker, error) {
	var (
		cli mqtt.Client
		err error
	)

	cli, err = makeConnect(clientId, host, port, login, password, cleanSession)
	if err != nil {
		return nil, err
	}
	return &mqttCli{
		qos:       qos,
		timeout:   defaultTimeout,
		isStopped: false,
		wg:        new(sync.WaitGroup),
		cli:       cli,
	}, nil
}

func makeConnect(
	clientId string,
	host string,
	port int,
	login string,
	password string,
	cleanSession bool,
) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts = opts.AddBroker(proto + host + ":" + strconv.Itoa(port))
	opts.ClientID = clientId
	opts.Username = login
	opts.Password = password

	opts.SetCleanSession(cleanSession)
	opts.SetConnectTimeout(defaultConnectTimeout)
	opts.SetMessageChannelDepth(defaultMessageChannelDepth)
	mqttClient := mqtt.NewClient(opts)

	token := mqttClient.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	if !mqttClient.IsConnectionOpen() {
		return nil, ErrMQTTDisconnected
	}

	return mqttClient, nil
}

func (c *mqttCli) Publish(topic string, txID string, payload interface{}) error {
	if c.cli == nil {
		return ErrConnRequired
	}
	if payload == nil {
		return ErrInvalidMessage
	}
	bt, err := c.wrapPayloadToMsg(txID, payload)
	if err != nil {
		return err
	}

	token := c.cli.Publish(topic, c.qos, false, bt)
	if token.WaitTimeout(c.timeout) && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *mqttCli) wrapPayloadToMsg(txID string, payload interface{}) ([]byte, error) {
	msg := Message{
		Payload:       payload,
		TransactionID: txID,
	}
	return json.Marshal(msg)
}

func (c *mqttCli) Close() error {
	if c.cli == nil {
		return ErrConnRequired
	}
	if c.isStopped {
		return nil
	}
	c.isStopped = true
	c.wg.Wait()
	c.cli.Disconnect(0)
	if c.cli.IsConnected() {
		return ErrStillConnected
	}
	return nil
}
