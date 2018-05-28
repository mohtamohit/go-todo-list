package config

import (
	"fmt"
	"time"
)

// DatabaseConfig is the structure for the database config.
type DatabaseConfig struct {
	databaseName        string
	databaseHost        string
	databaseUser        string
	databasePassword    string
	databasePort        int
	databaseMaxPoolSize int
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
}

func newDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		databaseName:        mustGetString("DB_NAME"),
		databaseHost:        mustGetString("DB_HOST"),
		databaseUser:        mustGetString("DB_USER"),
		databasePassword:    mustGetString("DB_PASSWORD"),
		databasePort:        mustGetInt("DB_PORT"),
		databaseMaxPoolSize: mustGetInt("DB_POOL"),
		ReadTimeout:         time.Millisecond * time.Duration(mustGetInt("DB_READ_TIMEOUT_MS")),
		WriteTimeout:        time.Millisecond * time.Duration(mustGetInt("DB_WRITE_TIMEOUT_MS")),
	}
}

// DatabaseMaxPoolSize returns the maximum pool size for postgres.
func (dc DatabaseConfig) DatabaseMaxPoolSize() int {
	return dc.databaseMaxPoolSize

}

// ConnectionString returns the postgres connection string to be used.
func (dc DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable",
		dc.databaseName,
		dc.databaseUser,
		dc.databasePassword,
		dc.databaseHost,
		dc.databasePort,
	)
}

// ConnectionURL returns the connection url to be used for postgres.
func (dc DatabaseConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dc.databaseUser,
		dc.databasePassword,
		dc.databaseHost,
		dc.databasePort,
		dc.databaseName,
	)
}
