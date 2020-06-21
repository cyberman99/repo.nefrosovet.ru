package influx

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/influxdata/influxdb/client/v2"

	"repo.nefrosovet.ru/maximus-platform/mailer/db/influx"
	"repo.nefrosovet.ru/maximus-platform/mailer/storage"
)

func GetMessageEventStorage(influxClient *influx.Client) storage.MessageEventStorage {
	return &MessageEventStorage{client: influxClient}
}

func Ensure(influxClient *influx.Client) error {
	if err := ensureDatabase(influxClient); err != nil {
		return err
	}

	return ensurePolicy(influxClient, false)
}

func ensureDatabase(influxClient *influx.Client) error {
	query := client.NewQuery(
		fmt.Sprintf(`CREATE DATABASE "%s"`, influxClient.Database),
		"",
		"",
	)

	response, err := influxClient.Query(query)
	if err != nil || response.Error() != nil {
		if err == nil {
			err = response.Error()
		}

		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "eventDB",
			"addr":     influxClient.Address,
			"database": influxClient.Database,
			"status":   "FAILED",
			"error":    err,
		}).Fatal("Ensuring influxDB database")

		return err
	}

	return nil
}

func ensurePolicy(influxClient *influx.Client, alter bool) error {
	if influxClient.Retention == "" {
		return nil
	}

	operation := "CREATE"
	if alter {
		operation = "ALTER"
	}

	query := client.NewQuery(
		fmt.Sprintf(`
				%s RETENTION POLICY "custom"
				ON "%s"
				DURATION %s
				REPLICATION 1 DEFAULT
			`,
			operation,
			influxClient.Database,
			influxClient.Retention,
		),
		"",
		"",
	)

	response, err := influxClient.Query(query)
	if err != nil || response.Error() != nil {
		if !alter {
			logrus.Debug("ALTER custom database retention policy")

			return ensurePolicy(influxClient, true)
		}

		if err == nil {
			err = response.Error()
		}

		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "eventDB",
			"addr":     influxClient.Address,
			"database": influxClient.Database,
			"status":   "FAILED",
			"error":    err,
		}).Fatal("Ensuring influxDB database retention policy")

		return err
	}

	return nil
}
