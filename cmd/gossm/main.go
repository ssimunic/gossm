package main

import (
	"flag"
	"io/ioutil"

	"github.com/ssimunic/gossm"
	"github.com/ssimunic/gossm/logger"
)

var configPath = flag.String("config", "config.json", "configuration file")
var logPath = flag.String("log", "log.txt", "log file")
var address = flag.String("http", ":8080", "address for http server")
var nolog = flag.Bool("nolog", false, "disable logging to file")
var logfilter = flag.String("logfilter", "", "text to filter log")

func main() {
	flag.Parse()
	jsonData, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic("error reading from configuration file")
	}

	if *nolog == true {
		logger.Disable()
	}

	if *logfilter != "" {
		logger.Filter(*logfilter)
	}

	config := gossm.NewConfig(jsonData)
	monitor := gossm.NewMonitor(config)
	go gossm.RunHttp(*address)
	monitor.Run()
}
