package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldLoadConfigfile(t *testing.T) {
	Load()
	assert.NotNil(t, Port())
	assert.NotEmpty(t, Log().LogLevel())
	assert.NotEmpty(t, Db())
}

func TestShouldLoadConfigFromEnvironment(t *testing.T) {
	configEnv := map[string]string{
		"APP_PORT":          "8888",
		"LOG_LEVEL":         "debug",
		"DATABASE_HOST":     "host",
		"DATABASE_PORT":     "3000",
		"DATABASE_USER":     "user",
		"DATABASE_PASSWORD": "123",
		"DATABASE_NAME":     "name",
	}

	for key, val := range configEnv {
		err := os.Setenv(key, val)
		assert.NoError(t, err, "Unable to set os env for"+key)
	}

	expectedDbConfig := DbConfig{
		host:     "host",
		port:     3000,
		name:     "name",
		password: "123",
		user:     "user",
	}

	Load()
	assert.Equal(t, 8888, Port())
	assert.Equal(t, "debug", Log().LogLevel())
	assert.Equal(t, expectedDbConfig, Db())
}
