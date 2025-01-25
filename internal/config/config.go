package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env          string `yaml:"env"`
	ServerConfig `yaml:"server_config"`
	MongoConfig  `yaml:"mongo_config"`
}

type ServerConfig struct {
	Addr         string        `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type MongoConfig struct {
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	DBName     string `yaml:"db_name"`
	CollName   string `yaml:"coll_name"`
	AuthSource string `yaml:"auth_source"`
}

func Load() *Config {
	configPath := flag.String("config", "", "path to the config file")
	flag.Parse()

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatal("config path is not set")
	}

	cfg := new(Config)

	if err := cleanenv.ReadConfig(*configPath, cfg); err != nil {
		log.Fatal("failed to read config")
	}

	return cfg
}
