package config

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

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

// LoadConfig загружает конфиг из указанного файла
func LoadConfig(filePath string) (*Config, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	godotenv.Load(filePath + ".env")

	dir := filepath.Dir(filePath)

	godotenv.Load(filepath.Join(dir, ".env"))

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

func LoadDefaultConfig() (*Config, error) {
	configFile, err := FindDefaultConfigFiles()
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

func FindDefaultConfigFiles() (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var configFile *os.File
	for range 5 {
		configFile, err = os.Open(filepath.Join(wd, "config.yaml"))
		if err != nil {
			wd = filepath.Join(wd, "..")
			continue
		}
		break
	}

	if configFile == nil {
		return nil, fmt.Errorf("config file isn't exist")
	}

	godotenv.Load(filepath.Join(wd, ".env"))

	return configFile, nil
}

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
