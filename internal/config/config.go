package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	Storage    `yaml:"storage"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string `yaml:"address" env-default:"localhost:8082"`
	Timeout     string `yaml:"timeout" env-default:"4s"`
	IdleTimeout string `yaml:"idle_timeout" env-default:"60s"`
}

type Storage struct {
	//StoragePath string `yaml:"storage_path" env-required:"true"`
	Host        string `yaml:"host" env-default:"localhost"`
	Port        string `yaml:"port" env-required:"true"`
	Database    string `yaml:"database" env-required:"true"`
	UserName    string `yaml:"username"`
	Password    string `yaml:"password"`
	MaxAttempts int    `yaml:"max_attempts" env-default:"5"`
}

func MustLoad() *Config {
	configPath := "../../config/config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	fmt.Println("config read")
	return &cfg
}
