package config

import (
	"os"

	"github.com/spf13/viper"
)

var appConfig Config

// Config stores the structure for all configuration that the app reads.
type Config struct {
	port int
	db   DatabaseConfig
}

// Port returns the port used by the application.
func Port() int { return appConfig.port }

// Database returns the config for the postgres database.
func Database() DatabaseConfig { return appConfig.db }

// Load prepares the configuration for the app.
func Load() {
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "error")
	viper.SetDefault("DB_QUERY_TIMEOUT_MS", 500)
	viper.SetDefault("LOG_FORMAT", "json")
	viper.SetDefault("DEFAULT_PHONE_REGION", "ID")
	viper.SetDefault("DB_READ_TIMEOUT_MS", "100")
	viper.SetDefault("DB_WRITE_TIMEOUT_MS", "500")

	if os.Getenv("ENVIRONMENT") == "test" {
		viper.SetConfigName("test")
	} else {
		viper.SetConfigName("application")
	}

	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()
	viper.AutomaticEnv()

	appConfig = Config{
		port: mustGetInt("APP_PORT"),
		db:   newDatabaseConfig(),
	}
}
