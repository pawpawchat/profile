package config

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	flag.Set("env", "testing")
	flag.Parse()

	wd, err := os.Getwd()
	assert.NoError(t, err)

	config, err := LoadConfig(wd + "/../config.yaml")
	assert.NoError(t, err)

	assert.NotNil(t, config)
	assert.NotEmpty(t, config.Env().GRPC_SERVER_ADDR)
	assert.NotEmpty(t, config.Env().DB_URL)

	fmt.Println("url=", config.Env().DB_URL)
}

func TestConfigureLogger(t *testing.T) {
	flag.Parse()

	wd, err := os.Getwd()
	assert.NoError(t, err)

	config, err := LoadConfig(wd + "/../config.yaml")
	assert.NoError(t, err)

	assert.NotNil(t, config, config.Env().GRPC_SERVER_ADDR)

	assert.NoError(t, ConfigureLogger(config))
}
