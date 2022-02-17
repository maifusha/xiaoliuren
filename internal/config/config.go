package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Conf *Config

const CONF_FILE = "./config.yaml"

type Config struct {
	Mode string

	Server
	Sqlite3
	Log
}

func NewConfig() *Config {
	return &Config{}
}

type Server struct {
	Host string
	Port string
}

type Sqlite3 struct {
	Path string
}

type Log struct {
	Runtime string
	Request string
	Panic   string
}

func init() {
	data, err := ioutil.ReadFile(CONF_FILE)
	if err != nil {
		log.Fatalln("Please uncomment the config.yaml.example into config.yaml and update it!")
	}

	Conf = NewConfig()
	err = yaml.Unmarshal(data, Conf)
	if err != nil {
		log.Fatalf("Parse config.yaml error：%s\n", err)
	}
}
