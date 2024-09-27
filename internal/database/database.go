package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prajnasatryass/go-clean-arch-example/config"
)

const (
	DriverPostgres = "postgres"
)

func NewDatabase(dbDriver string, dbConfig config.DatabaseConfig) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user='%s' password='%s' dbname=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

	conn, err := sqlx.Connect(dbDriver, connectionString)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
