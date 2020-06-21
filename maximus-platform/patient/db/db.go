package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

const (
	DriverPostgres  = "postgres"  // SELECT version(); "PostgreSQL 9.5.15 on x86_64-pc-linux-gnu, compiled by gcc (GCC) 4.8.3 20140911 (Red Hat 4.8.3-9), 64-bit"
	DriverCockroach = "cockroach" // SELECT version(); "CockroachDB CCL v1.1.1 (linux amd64, built 2017/10/19 15:31:46, go1.8.3)"
	proto           = "postgresql://"
	versionSQL      = `SELECT version();`
)

type DBConfig struct {
	Login, Pass, Host, Database, Name string
	MaxConn, MaxIdle, Port            int
	MaxLife                           time.Duration
	SSL                               bool
}

type Storer interface {
	Type() string
	Name() string
	InnerDB() *sql.DB
	Close() error
}

type database struct {
	*sqlx.DB
	dbType, dbName string
}

func SetupDB(conf DBConfig) (_ Storer, err error) {
	var ssl string
	if !conf.SSL {
		ssl = "&sslmode=disable"
	}
	if conf.Pass != "" {
		conf.Pass = ":"+conf.Pass
	}

	url := fmt.Sprintf(
		"%s%s%s@%s:%d/%s?application_name=%s%s",
		proto,
		conf.Login,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Name,
		ssl,
	)

	db, err := sqlx.Connect(DriverPostgres, url)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(conf.MaxConn)

	if conf.MaxIdle != -1 {
		db.SetMaxIdleConns(conf.MaxIdle)
	}
	db.SetConnMaxLifetime(conf.MaxLife)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	var res string
	err = db.QueryRow(versionSQL).Scan(&res)
	if err != nil {
		return nil, err
	}

	switch {
	case strings.Contains(strings.ToLower(res), DriverCockroach):
		return &database{db, DriverCockroach, conf.Database}, nil
	case strings.Contains(strings.ToLower(res), DriverPostgres):
		return &database{db, DriverPostgres, conf.Database}, nil
	}
	return &database{db, "", conf.Database}, nil

}

func (d *database) Type() string {
	return d.dbType
}

func (d *database) Name() string {
	return d.dbName
}

func (d *database) InnerDB() *sql.DB {
	return d.DB.DB
}

func (d *database) Close() error {
	return d.DB.Close()
}
