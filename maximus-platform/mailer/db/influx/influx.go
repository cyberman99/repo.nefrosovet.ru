package influx

import (
	"fmt"
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
)

type Client struct {
	Address   string
	Database  string
	Retention string

	influx.Client
}

func Connect(host string, port int, username, password, database, retention string) (*Client, error) {
	client := &Client{
		Database:  database,
		Retention: retention,
	}

	conf := influx.HTTPConfig{
		Addr: "http://" + host,
	}

	if port != 0 {
		conf.Addr += fmt.Sprintf(":%d", port)
	}

	client.Address = conf.Addr

	if len(username) > 0 && len(password) > 0 {
		conf.Username = username
		conf.Password = password
	}

	influxClient, err := influx.NewHTTPClient(conf)
	if err != nil {
		return nil, err
	}

	_, _, err = influxClient.Ping(10 * time.Second)
	if err != nil {
		return nil, err
	}

	client.Client = influxClient

	return client, nil
}
