package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Conf config

const CONF_FILE = "./config.yaml"

type config struct {
	Mode string

	Server
	Sqlite3
	Logfile
}

type Server struct {
	Host        string
	Port        string
	IdleTimeout int
}

type Sqlite3 struct {
	Path string
}

type Logfile struct {
	Runtime string
	Request string
	Panic   string
}

func init() {
	data, err := ioutil.ReadFile(CONF_FILE)
	if err != nil {
		log.Fatalln("Please uncomment the config.yaml.example into config.yaml and update it!")
	}

	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatalf("Parse config.yaml error：%s\n", err)
	}
}
