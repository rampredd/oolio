package app

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type AppConfig struct {
	ListenHost string `json:"listen_host"`
	Port       string `json:"port"`
	ServerName string `json:"server_name"`
	ApiKey     string `json:"api_key"`
}

var (
	Config      *AppConfig
	WgTerminate sync.WaitGroup
)

func LoadConfig(location string) error {
	log.Printf("Loading config file %s", location)
	var config *AppConfig
	f, e := os.Open(location)
	if e != nil {
		return e
	}
	defer f.Close()

	if e := json.NewDecoder(f).Decode(&config); e != nil {
		return e
	}

	Config = config

	return nil
}
