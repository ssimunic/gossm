package gossm

import (
	"fmt"

	"github.com/ssimunic/gossm/logger"
)

func (c *Config) validate() {
	c.validateSettings()
	c.validateServers()
}

func (c *Config) validateSettings() {
	// validate monitor settings
	if c.Settings.Monitor.CheckInterval <= 0 || c.Settings.Monitor.MaxConnections <= 0 || c.Settings.Monitor.Timeout <= 0 {
		panic("monitor settings missing")
	}

	// validate email notifications
	for _, email := range c.Settings.Notifications.Email {
		if email == (EmailSettings{}) {
			logger.Logln("No email notification settings found.")
		} else {
			if !email.isValid() {
				panic("invalid email settings")
			}
		}
	}

	// validate sms notifications
	for _, sms := range c.Settings.Notifications.Sms {
		if sms == (SmsSettings{}) {
			logger.Logln("No SMS notification settings found.")
		} else {
			// TODO: add sms validation
		}
	}
}

func (c *Config) validateServers() {
	if len(c.Servers) == 0 {
		panic("no servers found in config")
	}

	for _, server := range c.Servers {
		if server.Name == "" || server.IPAddress == "" || server.Port == 0 || server.Protocol == "" {
			panic(fmt.Sprintf("invalid data for server: %#v", server))
		}
		if server.CheckInterval == 0 {
			server.CheckInterval = c.Settings.Monitor.CheckInterval
		}
		if server.Timeout == 0 {
			server.Timeout = c.Settings.Monitor.Timeout
		}
	}
}

func (e *EmailSettings) isValid() bool {
	if e.Password == "" || e.Username == "" || e.SMTP == "" || e.Port == 0 {
		return false
	}
	return true
}

func (s *SmsSettings) isValid() bool {
	// TODO: add sms
	return false
}
