package mq

import (
	"encoding/json"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

// MQ is MQTT client struct
type MQ struct {
	Addr        string
	PubClientID string
	SubClientID string
	Login       string
	Password    string
	TopicIn     string
	TopicOut    string
	QoS         byte

	onceIn  sync.Once
	onceOut sync.Once

	pubClient mqtt.Client
	subClient mqtt.Client
}

var mqttClients map[string]mqtt.Client

// GetSubClient returns SUB MQTT connection
func (m *MQ) GetSubClient() mqtt.Client {
	m.onceIn.Do(func() {
		mqtt.ERROR = log.New()
		opts := mqtt.NewClientOptions()
		opts = opts.AddBroker("tcp://" + m.Addr)
		opts.ClientID = m.SubClientID
		opts.Username = m.Login
		opts.Password = m.Password

		m.subClient = mqtt.NewClient(opts)
		token := m.subClient.Connect()
		token.Wait()

		if token.Error() != nil {
			log.WithFields(log.Fields{
				"context":  "CORE",
				"resourse": "MQ",
				"addr":     m.Addr,
				"clientID": m.SubClientID,
				"status":   "FAILED",
				"error":    token.Error(),
			}).Fatal("Can't connect to MQ")
		}

		log.WithFields(log.Fields{
			"context": "CORE",
			// "version":  sender.Version,
			"resourse": "MQ",
			"addr":     m.Addr,
			"clientID": m.SubClientID,
			"status":   "CONNECTED",
		}).Info("Connection to MQ established (Subscription)")
	})

	return m.subClient
}

// GetPubClient returns PUB MQTT connection
func (m *MQ) GetPubClient() mqtt.Client {
	m.onceOut.Do(func() {
		// mqtt.ERROR = log.New()
		opts := mqtt.NewClientOptions()
		opts = opts.AddBroker("tcp://" + m.Addr)
		opts.ClientID = m.PubClientID
		opts.Username = m.Login
		opts.Password = m.Password

		m.pubClient = mqtt.NewClient(opts)
		token := m.pubClient.Connect()
		token.Wait()

		if token.Error() != nil {
			log.WithFields(log.Fields{
				"context":  "CORE",
				"resourse": "MQ",
				"addr":     m.Addr,
				"clientID": m.PubClientID,
				"status":   "FAILED",
				"error":    token.Error(),
			}).Fatal("Can't connect to MQ")
		}

		log.WithFields(log.Fields{
			"context":  "CORE",
			"resourse": "MQ",
			"addr":     m.Addr,
			"clientID": m.PubClientID,
			"status":   "CONNECTED",
		}).Info("Connection to MQ established (Publication)")
	})

	return m.pubClient
}

// IncomingMessage is MQTT message struct
type IncomingMessage struct {
	TransactionID string `json:"transactionID"`
	Payload       struct {
		Query  string `json:"query,omitempty"`
		Method string `json:"method,omitempty"`
		Params struct {
			AccessToken string `json:"accessToken,omitempty"`
		} `json:"params,omitempty"`
		Body struct {
			ChannelID   string `json:"channelID,omitempty"`
			Destination string `json:"destination,omitempty"`
			Data        string `json:"data,omitempty"`
		} `json:"body,omitempty"`
	} `json:"payload,omitempty"`
}

// Listen gets input messages
func (m *MQ) Listen() error {
	token := m.GetSubClient().Subscribe(
		m.TopicIn,
		0,
		func(client mqtt.Client, msg mqtt.Message) {
			m.ParseIncomingMessage(msg.Payload())
		},
	)

	token.Wait()

	if token.Error() != nil {
		log.WithFields(log.Fields{
			"context":  "MQ",
			"QUEUE":    m.TopicIn,
			"clientID": m.SubClientID,
			"error":    token.Error(),
		}).Fatal("Listen MQ messages error")
	}

	return token.Error()
}

// ParseIncomingMessage parses incoming queue message
func (m *MQ) ParseIncomingMessage(body []byte) {
	var message IncomingMessage
	if err := json.Unmarshal(body, &message); err != nil {
		log.WithFields(log.Fields{
			"context": "MQ",
			"error":   err,
			"body":    string(body),
		}).Error("Can't unmarshal incoming message body")

		return
	}

	log.WithFields(log.Fields{
		"context":       "MQ",
		"action":        "REQUEST",
		"transactionID": message.TransactionID,
	}).Info("Got message")
	log.WithFields(log.Fields{
		"message": string(body),
	}).Debug()

	switch message.Payload.Query {
	default:
		// nothing
	}

}

// PublishMessage sends queue message
func (m *MQ) PublishMessage(message []byte) {
	t := m.GetPubClient().Publish(m.TopicOut, m.QoS, false, message)
	t.Wait()

	if t.Error() != nil {
		log.WithFields(log.Fields{
			"error":   t.Error(),
			"message": string(message),
		}).Error("Can't send message to queue")
	}
}
