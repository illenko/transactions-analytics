package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"strings"
)

type DatabaseConfig struct {
	Url string `yaml:"url"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type AppConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func Get(log *slog.Logger) (AppConfig, error) {
	conf := viper.New()
	conf.SetConfigFile("config.yaml")
	replacer := strings.NewReplacer(".", "_")
	conf.SetEnvKeyReplacer(replacer)
	conf.AutomaticEnv()

	err := conf.ReadInConfig()
	if err != nil {
		log.Error(fmt.Sprintf("error reading config file: %v", err))
	}
	var cfg AppConfig

	err = conf.Unmarshal(&cfg)
	if err != nil {
		log.Error(fmt.Sprintf("configuration unmarshalling failed!. Error: %v", err))
		return cfg, err
	}
	return cfg, nil
}
