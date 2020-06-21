package influx

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

type Influxer interface {
	DBName() string

	client.Client
}

type InfluxHTTPCli struct {
	dbName string

	client.Client
}

func (cli *InfluxHTTPCli) DBName() string {
	return cli.dbName
}

var influxer Influxer

func GetInfluxer() Influxer {
	if influxer == nil {
		log.Fatal("Influx session is not found")
	}

	return influxer
}

func ConnectHTTP(host string, port int, username, password, database, retention string) (Influxer, error) {
	ic, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:%d", host, port),
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	_, _, err = ic.Ping(time.Second * 10)
	if err != nil {
		return nil, err
	}

	if err = ensureDatabase(database, retention, ic); err != nil {
		return nil, err
	}

	influxer = &InfluxHTTPCli{
		database,
		ic,
	}

	return influxer, nil
}

func ensureDatabase(database, retention string, ic client.Client) error {
	err := createDatabase(database, ic)
	if err != nil {
		return err
	}

	err = createRetentionPolicy(false, database, retention, ic)
	if err != nil {
		return err
	}

	return nil
}

func createDatabase(database string, ic client.Client) error {
	q := fmt.Sprintf(`
		CREATE DATABASE %s
	`, database)

	resp, err := ic.Query(client.NewQuery(q, "", ""))

	if err == nil {
		err = resp.Error()
	}
	return err
}

func createRetentionPolicy(alter bool, database, retention string, ic client.Client) error {
	if retention == "" {
		return nil
	}

	operation := "CREATE"
	if alter {
		operation = "ALTER"
	}

	q := fmt.Sprintf(`
		%s RETENTION POLICY "custom"
		ON "%s"
		DURATION %s
		REPLICATION 1 DEFAULT
	`, operation, database, retention)

	response, err := ic.Query(client.NewQuery(q, "", ""))
	if err != nil || response.Error() != nil {
		if !alter {
			return createRetentionPolicy(true, database, retention, ic)
		}

		if err == nil {
			err = response.Error()
		}

		return err
	}

	return nil
}
