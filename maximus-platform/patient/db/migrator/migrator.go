package migrator

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/cockroachdb"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/pkg/errors"
	"log"
	"repo.nefrosovet.ru/maximus-platform/patient/db"
	"strings"
)

const (
	stepUp   = "up"
	stepDown = "down"

	migrationsTable     = "migrations"
	migrationsLockTable = "migrations_lock"
)

func ApplyMigrations(d db.Storer, up, down, flush bool, path string) error {
	var (
		step string
		err  error
		dr   database.Driver
	)

	switch {
	case up && !down:
		step = stepUp
	case !up && down:
		step = stepDown
	}

	switch d.Type() {
	case db.DriverCockroach:
		migrCfg := new(cockroachdb.Config)
		migrCfg.DatabaseName = d.Name()
		migrCfg.MigrationsTable = migrationsTable
		migrCfg.LockTable = migrationsLockTable
		migrCfg.ForceLock = true

		dr, err = cockroachdb.WithInstance(d.InnerDB(), migrCfg)
		if err != nil {
			return errors.Errorf("Connecting migrations with cockroachdb driver failed: %v", err)
		}
	default: // Postgres
		migrCfg := new(postgres.Config)
		migrCfg.DatabaseName = d.Name()
		migrCfg.MigrationsTable = migrationsTable

		dr, err = postgres.WithInstance(d.InnerDB(), migrCfg)
		if err != nil {
			return errors.Errorf("Connecting migrations with postgres driver failed: %v", err)
		}
	}
	defer func() {
		er := dr.Close()
		if er != nil {
			err = errors.Wrap(err, er.Error())
		}

	}()

	return applyMigrations(dr, step, path, flush)
}

func applyMigrations(driver database.Driver, step, path string, flush bool) error {
	m, err := migrate.NewWithDatabaseInstance(path, "", driver)
	if err != nil {
		return errors.Errorf("Migrations init failed: %v", err)
	}

	defer m.Close()

	if flush {
		err = m.Force(-1)
		if err != nil {
			return errors.Errorf("Migrations flushing failed: %v", err)
		}
		log.Println("flushed")
		return nil
	}

	var ver uint
	ver, _, err = m.Version()

	switch step {
	case stepUp:
		err = m.Up()
		if err != nil {
			if strings.Contains(err.Error(), "file does not exist") {
				log.Println(err, "No migrations left. Up to date")
				return nil
			}
			return errors.Errorf("Migrations applying failed: %v", err)
		}

		ver, _, err = m.Version()
		log.Printf("Migrations applied. Current version: %v, err: %v", ver, err)
	case stepDown:
		if ver == 0 {
			log.Println("Migrations zeroed")
			return nil
		}

		err = m.Steps(-1)
		if err == migrate.ErrNoChange {
			log.Println(err)
			return nil
		}
		if err != nil {
			return errors.Errorf("Migrations applying failed: %v", err)
		}

		ver, _, err = m.Version()
		log.Printf("Migrations rollbacked. Current version: %v, err: %v", ver, err)
	default:
		return errors.Errorf("Wrong directive given: up or down?")
	}
	return nil
}
