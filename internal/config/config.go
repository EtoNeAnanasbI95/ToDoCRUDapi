package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env              string `yaml:"env" env-default:"local" json:"env"`
	ConnectionString string `yaml:"connection_string" json:"connection-string" env-required:"true"`
	Api              Api    `yaml:"api" json:"api"`
}

type Api struct {
	Port    int           `yaml:"port" env-default:"8080" json:"port"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s" json:"timeout"`
}

func MustLoadConfig() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("CONFIG_PATH environment variable not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("CONFIG_PATH file does not exist")
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Failed to parse config with: " + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var configPath string

	// --config="path/to/config.yaml"
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()
	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}
	return configPath
}
