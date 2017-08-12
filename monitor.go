package gossm

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/ssimunic/gossm/config"
	"github.com/ssimunic/gossm/logger"
)

type Monitor struct {
	config    *config.Config
	pipe      chan *Server
	semaphore chan struct{}
	stop      chan struct{}
}

type Server config.Server

func NewMonitor(c *config.Config) *Monitor {
	return &Monitor{
		config:    c,
		pipe:      make(chan *Server),
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
		m.pipe <- s
	}
}

func (m *Monitor) monitor() {
	for {
		select {
		case server := <-m.pipe:
			m.semaphore <- struct{}{}
			go func() {
				server.checkStatus()
				<-m.semaphore
			}()
		case <-m.stop:
			logger.Log("Terminating.")
			os.Exit(0)
		}
	}
}

func (s *Server) checkStatus() {
	logger.Logln("Checking", s)
	formattedAddress := fmt.Sprintf("%s:%d", s.IPAddress, s.Port)
	timeoutSeconds := time.Duration(s.Timeout) * time.Second
	conn, err := net.DialTimeout(s.Protocol, formattedAddress, timeoutSeconds)
	if err != nil {
		logger.Logln(err)
		logger.Logln("ERROR", s)
		go s.handleFailure()
		return
	}
	defer conn.Close()
	logger.Logln("OK", s)
}

func (s *Server) handleFailure() {
	logger.Logln("Sending notification for", s)
	// TODO: Send notifications
}

func (s *Server) String() string {
	return fmt.Sprintf("%s %s:%d", s.Protocol, s.IPAddress, s.Port)
}
