package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var appConfig Config

type Config struct {
	port       int
	appName    string
	appVersion string
	db         DBConfig
}

func Load() {
	viper.SetDefault("APP_PORT", "8080")

	viper.SetConfigName("application")
	if os.Getenv("ENVIRONMENT") == "test" {
		viper.SetConfigName("test")
	}

	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()
	viper.AutomaticEnv()

	appConfig = Config{
		port:       extractIntValue("APP_PORT"),
		appVersion: extractStringValue("APP_VERSION"),
		appName:    extractStringValue("APP_NAME"),
		db: DBConfig{
			host:     extractStringValue("DB_HOST"),
			port:     extractIntValue("DB_PORT"),
			name:     extractStringValue("DB_NAME"),
			user:     extractStringValue("DB_USER"),
			password: extractStringValue("DB_PASSWORD"),
		},
	}
}

func Port() int {
	return appConfig.port
}

func Db() DBConfig {
	return appConfig.db
}

func AppName() string {
	return appConfig.appName
}

func AppVersion() string {
	return appConfig.appVersion
}

func extractStringValue(key string) string {
	checkPresenceOf(key)
	return viper.GetString(key)
}

func extractBoolValue(key string) bool {
	checkPresenceOf(key)
	return viper.GetBool(key)
}

func extractIntValue(key string) int {
	checkPresenceOf(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid Integer value", key))
	}

	return v
}

func checkPresenceOf(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	}
}
