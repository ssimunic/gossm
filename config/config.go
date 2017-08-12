package config

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Servers  Servers  `json:"servers"`
	Settings Settings `json:"settings"`
}

type Servers []Server

type Server struct {
	Name          string `json:"name"`
	IPAddress     string `json:"ipAddress"`
	Port          int    `json:"port"`
	Protocol      string `json:"protocol"`
	CheckInterval int    `json:"checkInterval"`
	Timeout       int    `json:"timeout"`
}

type Settings struct {
	Monitor       Monitor
	Notifications Notifications
}

type Monitor struct {
	CheckInterval  int `json:"checkInterval"`
	Timeout        int `json:"timeout"`
	MaxConnections int `json:"maxConnections"`
}

type Notifications struct {
	Email Email `json:"email"`
	Sms   Sms   `json:"sms"`
}

type Email struct {
	Smtp     string `json:"smtp"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Sms struct{}

func New(jsonData []byte) *Config {
	config := &Config{}
	err := json.Unmarshal(jsonData, config)
	if err != nil {
		panic("error parsing json configuration data")
	}
	config.validate()
	return config
}

func (c *Config) validate() {
	if len(c.Servers) == 0 {
		panic("no servers found in config")
	}

	if c.Settings.Monitor.CheckInterval == 0 || c.Settings.Monitor.MaxConnections == 0 || c.Settings.Monitor.Timeout == 0 {
		panic("monitor settings missing")
	}

	for _, server := range c.Servers {
		if server.Name == "" || server.IPAddress == "" || server.Port == 0 || server.Protocol == "" {
			panic(fmt.Sprintf("invalid data for server: %#v", server))
		}
	}
}
