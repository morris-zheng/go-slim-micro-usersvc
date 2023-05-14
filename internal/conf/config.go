package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Logger struct {
	Level string `yaml:"level"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Etcd struct {
	Endpoints         []string `yaml:"endpoints"`
	Prefix            string   `yaml:"prefix"`
	KeepAliveInterval int64    `yaml:"keep_alive_interval"`
}

type Config struct {
	Name   string `yaml:"name"`
	Host   string `yaml:"host"`
	Env    string `yaml:"env"`
	Debug  bool   `yaml:"debug"`
	Port   int    `yaml:"port"`
	Logger Logger `yaml:"logger"`
	Mysql  Mysql  `yaml:"mysql"`
	Etcd   Etcd   `yaml:"etcd"`
}

func Load(path string) *Config {
	cb, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("read config err", err)
	}

	var c Config

	err = yaml.Unmarshal(cb, &c)
	if err != nil {
		log.Fatal("unmarshal config err", err)
	}

	return &c
}
