package migrator

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate"
	migrateDatabase "github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/cockroachdb"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	DriverPostgres  = "postgres"  // SELECT version(); "PostgreSQL 9.5.15 on x86_64-pc-linux-gnu, compiled by gcc (GCC) 4.8.3 20140911 (Red Hat 4.8.3-9), 64-bit"
	DriverCockroach = "cockroach" // SELECT version(); "CockroachDB CCL v1.1.1 (linux amd64, built 2017/10/19 15:31:46, go1.8.3)"
	DriverMySQL     = "mysql"

	stepUp   = "up"
	stepDown = "down"

	migrationsTable     = "migrations"
	migrationsLockTable = "migrations_lock"
)

type Driver interface {
	Type() string
	Name() string
	InnerDB() *sql.DB
}


type Loggerf func(format string, args ...interface{}) (int, error)

func ApplyMigrations(d Driver, logger Loggerf, up, down, flush bool, path string) (err error) {
	var (
		driver migrateDatabase.Driver
		step   string
	)
	if logger == nil {
		logger = func(format string, args ...interface{}) (i int, err error) {
			return 0, nil
		}
	}

	switch {
	case up && !down:
		step = stepUp
	case !up && down:
		step = stepDown
	}

	switch d.Type() {
	case DriverCockroach:
		cfg := new(cockroachdb.Config)
		cfg.DatabaseName = d.Name()
		cfg.MigrationsTable = migrationsTable
		cfg.LockTable = migrationsLockTable
		cfg.ForceLock = true

		driver, err = cockroachdb.WithInstance(d.InnerDB(), cfg)
		if err != nil {
			return fmt.Errorf("connecting with cockroachdb driver failed: %w", err)
		}
	case DriverMySQL:
		cfg := new(mysql.Config)
		cfg.DatabaseName = d.Name()
		cfg.MigrationsTable = migrationsTable

		driver, err = mysql.WithInstance(d.InnerDB(), cfg)
		if err != nil {
			return fmt.Errorf("connecting with mysql driver failed: %w", err)
		}
	default: // Postgres
		cfg := new(postgres.Config)
		cfg.DatabaseName = d.Name()
		cfg.MigrationsTable = migrationsTable

		driver, err = postgres.WithInstance(d.InnerDB(), cfg)
		if err != nil {
			return fmt.Errorf("connecting with postgres driver failed: %w", err)
		}
	}

	defer func() {
		closeErr := driver.Close()
		if closeErr != nil {
			err = fmt.Errorf("%s: %w", err.Error(), closeErr)
		}
	}()

	return applyMigrations(driver, logger, step, path, flush)
}

func applyMigrations(driver migrateDatabase.Driver, logf Loggerf, step, path string, flush bool) error {
	m, err := migrate.NewWithDatabaseInstance(path, "", driver)
	if err != nil {
		return fmt.Errorf("migrations init failed: %w", err)
	}
	defer m.Close()

	if flush {
		err = m.Force(-1)
		if err != nil {
			return fmt.Errorf("migrations flushing failed: %w", err)
		}
		logf("flushed")
		return nil
	}

	var ver uint
	ver, _, err = m.Version()

	switch step {
	case stepUp:
		err = m.Up()
		if err != nil {
			if strings.Contains(err.Error(), "file does not exist") {
				logf("No migrations left. Up to date: %v", err)
				return nil
			}
			return fmt.Errorf("migrations applying failed: %w", err)
		}

		ver, _, err = m.Version()
		logf("Migrations applied. Current version: %v, err: %v", ver, err)
	case stepDown:
		if ver == 0 {
			logf("Migrations zeroed")
			return nil
		}

		err = m.Steps(-1)
		if err == migrate.ErrNoChange {
			logf(err.Error())
			return nil
		}
		if err != nil {
			return fmt.Errorf("migrations applying failed: %w", err)
		}

		ver, _, err = m.Version()
		logf("Migrations rollbacked. Current version: %v, err: %v", ver, err)
	default:
		return fmt.Errorf("wrong directive given: allowed up or down")
	}

	return nil
}
