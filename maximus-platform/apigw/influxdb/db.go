package influxdb

import (
	"fmt"
	"time"

	"sync"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
)

// InfluxClient is InfluxDB client instance
var (
	Addr     string
	Username string
	Password string
	Database string

	cln  client.Client
	once sync.Once
)

// New returns InfluxDB client instance
func New(addr, username, password, database, retention string) {
	// over-protection from double-calling
	once.Do(func() {

		Addr = addr
		Username = username
		Password = password
		Database = database

		var err error

		dbConf := client.HTTPConfig{
			Addr: "http://" + Addr,
		}
		if Username != "" {
			dbConf.Username = Username
			dbConf.Password = Password
		}

		// Create a new HTTPClient
		cln, err = client.NewHTTPClient(dbConf)

		if err != nil {
			logrus.WithError(err).Debug()
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"recurse":  "eventDB",
				"addr":     Addr,
				"database": database,
				"status":   "FAILED",
			}).Fatal("Connection to influxDB")
		}

		_, _, err = cln.Ping(10 * time.Second)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "eventDB",
				"addr":     Addr,
				"database": database,
				"status":   "FAILED",
			}).Fatal("Connection to influxDB")
		}

		EnsureDatabase(database)
		EnsureDatabaseRetentionPolicy(database, retention)

		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "eventDB",
			"addr":     Addr,
			"database": database,
			"status":   "CONNECTED",
		}).Info("Connection to influxDB")

		return
	})
}

// EnsureDatabase - creates DB with data retention policy
func EnsureDatabase(database string) {
	queryString := fmt.Sprintf(`CREATE DATABASE "%s"`, database)
	query := client.NewQuery(queryString, "", "")
	response, err := cln.Query(query)
	if err != nil || response.Error() != nil {
		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "eventDB",
			"addr":     Addr,
			"database": database,
			"status":   "FAILED",
			"error":    err,
		}).Fatal("Ensuring influxDB database")
	}
}

// alter changes INSERT query to UPDATE on EnsureDatabaseRetentionPolicy function
var alter bool

// EnsureDatabaseRetentionPolicy - creates DB retention policy
func EnsureDatabaseRetentionPolicy(database, retention string) {
	if retention == "" {
		return
	}

	operation := "CREATE"
	if alter {
		operation = "ALTER"
	}
	queryString := fmt.Sprintf(`
		%s RETENTION POLICY "custom"
		ON "%s"
		DURATION %s
		REPLICATION 1 DEFAULT
	`,
		operation,
		database,
		retention,
	)
	query := client.NewQuery(queryString, "", "")
	response, err := cln.Query(query)
	if err != nil || response.Error() != nil {
		if !alter {
			alter = true
			logrus.Debug("ALTER custom database retention policy")
			EnsureDatabaseRetentionPolicy(database, retention)

			return
		}

		if err == nil {
			err = response.Error()
		}

		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "eventDB",
			"addr":     Addr,
			"database": database,
			"status":   "FAILED",
			"error":    err,
		}).Fatal("Ensuring influxDB database retention policy")
	}
}

func stringValue(value interface{}) string {
	if value == nil {
		return ""
	}

	return value.(string)
}
