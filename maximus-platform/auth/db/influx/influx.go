package influx

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

type Client struct {
	Address   string
	Username  string
	Password  string
	Database  string
	Retention string

	Cln client.Client
}

// NewClient returns InfluxDB client instance
func NewClient(addr, username, password, database, retention string) (*Client, error) {
	var err error

	dbConf := client.HTTPConfig{
		Addr: "http://" + addr,
	}
	if username != "" {
		dbConf.Username = username
		dbConf.Password = password
	}

	c := Client{
		Address: addr,
		Username: username,
		Password: password,
		Database: database,
		Retention: retention,
	}

	// Create a new HTTPClient
	cln, err := client.NewHTTPClient(dbConf)
	if err != nil {
		return nil, err
	}

	_, _, err = cln.Ping(20 * time.Second)
	if err != nil {
		return nil, err
	}

	c.Cln = cln

	return &c, nil
}




