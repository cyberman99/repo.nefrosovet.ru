package mq

import (
    "encoding/json"
    "sync"

    log "github.com/Sirupsen/logrus"
    "github.com/eclipse/paho.mqtt.golang"

    "repo.nefrosovet.ru/maximus-platform/mailer/sender"
    "repo.nefrosovet.ru/maximus-platform/mailer/storage/default"
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

func GetStorage() *_default.Storage {
	return _default.GetStorage()
}

// GetSubClient returns SUB MQTT connection
func (m *MQ) GetSubClient() mqtt.Client {
	m.onceIn.Do(func() {
		mqtt.ERROR = log.New()
		opts := mqtt.NewClientOptions()
		opts = opts.AddBroker("tcp://" + m.Addr)
		opts.ClientID = m.SubClientID
		opts.Username = m.Login
		opts.Password = m.Password
		opts.SetOnConnectHandler(func(client mqtt.Client) {
			log.WithFields(log.Fields{
				"context":  "CORE",
				"version":  sender.Version,
				"resource": "MQ",
				"addr":     m.Addr,
				"clientID": m.SubClientID,
				"status":   "CONNECTED",
			}).Info("Connection to MQ established (Subscription)")
			token := client.Subscribe(
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
		})

		m.subClient = mqtt.NewClient(opts)
		token := m.subClient.Connect()
		token.Wait()

		if token.Error() != nil {
			log.WithFields(log.Fields{
				"context":  "CORE",
				"resource": "MQ",
				"addr":     m.Addr,
				"clientID": m.SubClientID,
				"status":   "FAILED",
				"error":    token.Error(),
			}).Fatal("Can't connect to MQ")
		}

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
				"resource": "MQ",
				"addr":     m.Addr,
				"clientID": m.PubClientID,
				"status":   "FAILED",
				"error":    token.Error(),
			}).Fatal("Can't connect to MQ")
		}

		log.WithFields(log.Fields{
			"context":  "CORE",
			"resource": "MQ",
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

// MessageSendResult is response struct
type MessageSendResult struct {
	TransactionID string `json:"transactionID"`
	Payload       struct {
		Body struct {
			Version string   `json:"version,omitempty"`
			Errors  []string `json:"errors"`

			Status struct {
				Code    int    `json:"code,omitempty"`
				Message string `json:"message,omitempty"`
			} `json:"status,omitempty"`
		} `json:"body,omitempty"`

		Data []sender.Result `json:"data,omitempty"`
	} `json:"payload,omitempty"`
}


// Listen gets input messages
func (m *MQ) Listen() {
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
	case "sendMessage":
	    m.handleSendMessage(&message)
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

// PublishSendErrorMessage sends queue message with error
func (m *MQ) PublishSendErrorMessage(transactionID, errMessage string) {
	var response MessageSendResult
	response.TransactionID = transactionID

	response.Payload.Body.Version = sender.Version
	response.Payload.Body.Status.Code = 400
	response.Payload.Body.Status.Message = "ERROR"
	response.Payload.Body.Errors = append(response.Payload.Body.Errors, errMessage)

	message, err := json.Marshal(&response)
	if err != nil {
		log.WithFields(log.Fields{
			"context": "MQ",
			"error":   err,
		}).Error("Can't marshal sending message responce")

		return
	}

	log.WithFields(log.Fields{
		"context":       "MQ",
		"action":        "REQUEST",
		"transactionID": transactionID,
	}).Info("Sending MQ error answer")
	log.WithFields(log.Fields{
		"message": string(message),
	}).Debug()

	m.PublishMessage(message)

	return
}
