package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-default:"/var/run/docker.sock"`
	GRPCServer
}

type GRPCServer struct {
	Port    int    `yaml:"port" env-default:"77077"`
	Timeout string `yaml:"timeout" env-default:"5s"`
}

func MustLoad() *Config {
	return MustLoadPath("D:/go Projects/notes_grpc/notes/config/local.yml")
}

func MustLoadPath(configPath string) *Config {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file not found: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
