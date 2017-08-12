package main

import (
	"flag"
	"io/ioutil"

	"github.com/ssimunic/gossm"
	"github.com/ssimunic/gossm/config"
)

var configPath = flag.String("config", "config.json", "configuration file")
var logPath = flag.String("log", "log.txt", "log file")

func main() {
	flag.Parse()
	jsonData, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic("error reading from configuration file")
	}

	config := config.New(jsonData)
	monitor := gossm.NewMonitor(config)
	monitor.Run()
}
