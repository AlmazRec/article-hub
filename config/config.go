package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
	}

	Server struct {
		Port string `yaml:"port"`
	}

	Database struct {
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Port      string `yaml:"port"`
		Host      string `yaml:"host"`
		DBName    string `yaml:"dbname"`
		Sslmode   string `yaml:"sslmode"`
		ParseTime bool   `yaml:"parseTime"`
	}

	JWT struct {
		Secret     string `yaml:"secret"`
		Expiration string `yaml:"expiration"`
	}
}

func MustLoad(cfgPath string) *Config {
	var cfg Config

	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
