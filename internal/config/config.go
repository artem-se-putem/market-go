package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const configPath = ".env"

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	// StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer HTTPServer  `yaml:"http_server"`
	Postgres Postgres `yaml:"postgres"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

type Postgres struct {
	Host string `yaml:"address" env-default:"localhost:8080"`
	Port string `yaml:"port" env-default:"5432"`
	Dbname string `yaml:"dbname" env-default:"postgres"`
	User string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"password"`
	Sslmode string `yaml:"sslmode" env-default:"disable"`
}

func MustLoad() *Config {

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
