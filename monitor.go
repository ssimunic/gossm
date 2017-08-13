package gossm

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/ssimunic/gossm/logger"
)

type Monitor struct {
	config    *Config
	checker   chan *Server
	notifiers Notifiers
	notifier  chan *Server
	// Used to regulate number of concurrent connections
	semaphore chan struct{}
	// Exit on receive
	stop chan struct{}
}

func NewMonitor(c *Config) *Monitor {
	return &Monitor{
		config:    c,
		checker:   make(chan *Server),
		notifiers: c.Settings.Notifications.GetNotifiers(),
		notifier:  make(chan *Server),
		semaphore: make(chan struct{}, c.Settings.Monitor.MaxConnections),
		stop:      make(chan struct{}),
	}
}

func (m *Monitor) Run() {
	m.RunForSeconds(0)
}

func (m *Monitor) RunForSeconds(runningSeconds int) {
	if runningSeconds != 0 {
		go func() {
			runningSecondsTime := time.Duration(runningSeconds) * time.Second
			<-time.After(runningSecondsTime)
			m.stop <- struct{}{}
		}()
	}

	logger.Logln("Starting monitor.")
	for _, server := range m.config.Servers {
		server := Server(server)
		go m.handleServer(&server)
	}
	m.monitor()
}

func (m *Monitor) handleServer(s *Server) {
	tickerSeconds := time.NewTicker(time.Duration(s.CheckInterval) * time.Second)

	for range tickerSeconds.C {
		m.checker <- s
	}
}

func (m *Monitor) monitor() {
	go m.listenServers()
	go m.listenNotifiers()
	<-m.stop
	logger.Log("Terminating.")
	os.Exit(0)
}

func (m *Monitor) listenServers() {
	for {
		server := <-m.checker
		go func() {
			m.semaphore <- struct{}{}
			m.checkServerStatus(server)
			<-m.semaphore
		}()
	}
}

func (m *Monitor) listenNotifiers() {
	for {
		server := <-m.notifier
		go func() {
			m.notifiers.NotifyAll(server.String())
		}()
	}
}

func (m *Monitor) checkServerStatus(server *Server) {
	logger.Logln("Checking", server)
	formattedAddress := fmt.Sprintf("%s:%d", server.IPAddress, server.Port)
	timeoutSeconds := time.Duration(server.Timeout) * time.Second
	conn, err := net.DialTimeout(server.Protocol, formattedAddress, timeoutSeconds)
	if err != nil {
		logger.Logln(err)
		logger.Logln("ERROR", server)
		go func() {
			m.notifier <- server
		}()
		return
	}
	defer conn.Close()
	logger.Logln("OK", server)
}
