package config

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

var env *string

func init() {
	env = flag.String("env", "stage", "application runtime environment")
}

type Config struct {
	Environment map[string]*Environment `yaml:"environment"`
}

type Environment struct {
	GRPC_SERVER_ADDR string `yaml:"grpc_server_addr"`
	LOG_LEVEL        string `yaml:"log_level"`
	DB_URL           string `yaml:"database_url"`
}

func (c *Config) Env() *Environment {
	return c.Environment[*env]
}

func LoadConfig(filePath string) (*Config, error) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	err = yaml.NewDecoder(configFile).Decode(&config)

	if err != nil {
		return nil, err
	}

	for i := range config.Environment {
		config.Environment[i].DB_URL = os.ExpandEnv(config.Environment[i].DB_URL)
	}

	return &config, nil
}

// Set the logging level
func ConfigureLogger(config *Config) error {

	switch config.Env().LOG_LEVEL {
	case "error":
		slog.SetLogLoggerLevel(slog.LevelError)
	case "debug":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "info":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	default:
		return fmt.Errorf("undefined log level: %s", config.Env().LOG_LEVEL)
	}

	return nil
}
