// Package config ...
package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App  AppConfig   `toml:"app"`
	Bots []BotConfig `toml:"bots"`
}

type AppConfig struct {
	LogLevel string `toml:"log_level"`
}

type BotConfig struct {
	ID      string        `toml:"id"`
	Token   string        `toml:"token"`
	Enable  bool          `toml:"enable"`
	Storage StorageConfig `toml:"storage"`
}

type StorageConfig struct {
	Driver string `toml:"driver"`
	DSN    string `toml:"dsn"`
}

func MustLoad() *Config {
	var cfg Config
	if err := cleanenv.ReadConfig("./config.toml", cfg); err != nil {
		log.Fatal(fmt.Printf("Could not read config file\nНе возможно прочитать файл конфигураци: %v", err))
	}

	return &cfg
}
