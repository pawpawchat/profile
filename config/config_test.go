package config

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	flag.Parse()

	wd, err := os.Getwd()
	assert.NoError(t, err)

	config, err := LoadConfig(wd + "/../config.yaml")
	assert.NoError(t, err)

	assert.NotNil(t, config, config.Env().GRPC_SERVER_ADDR)
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
