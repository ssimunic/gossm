package main

import (
	"flag"
	"io/ioutil"
	"time"

	"github.com/ssimunic/gossm"
	"github.com/ssimunic/gossm/logger"
)

var configPath = flag.String("config", "configs/default.json", "configuration file")
var logPath = flag.String("log", "logs/from-"+time.Now().Format("2006-01-02")+".log", "log file")
var address = flag.String("http", ":8080", "address for http server")
var nolog = flag.Bool("nolog", false, "disable logging to file only")
var logfilter = flag.String("logfilter", "", "text to filter log by (both console and file)")

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

	logger.SetFilename(*logPath)

	config := gossm.NewConfig(jsonData)
	monitor := gossm.NewMonitor(config)
	go gossm.RunHttp(*address, monitor)
	monitor.Run()
}
