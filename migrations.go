package celeritas

import (
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Database is the type for the database package. Members that are exported in this type
// are available to any application that uses it.
func (c *Celeritas) MigrateUp(dsn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err := m.Up(); err != nil {
		log.Println("Error running migration:", err)
		return err
	}

	return nil
}

func (c *Celeritas) MigrateDownAll(dsn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err = m.Down(); err != nil {
		return err
	}
	return nil

}

func (c *Celeritas) Steps(n int, dsn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err = m.Steps(n); err != nil {
		return err
	}
	return nil

}

func (c *Celeritas) MigrateForce(dsn string) error {
	m, err := migrate.New("file://"+c.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err = m.Force(-1); err != nil {
		return err
	}

	return nil

}
