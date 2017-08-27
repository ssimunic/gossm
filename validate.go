package gossm

import (
	"fmt"
)

func (c *Config) Validate() error {
	if err := c.Settings.Validate(); err != nil {
		return fmt.Errorf("invalid settings: %v", err)
	}
	if err := c.Servers.Validate(); err != nil {
		return fmt.Errorf("invalid servers: %v", err)
	}
	return nil
}

func (s *Settings) Validate() error {
	if err := s.Monitor.Validate(); err != nil {
		return fmt.Errorf("invalid monitor settings: %v", err)
	}
	if err := s.Notifications.Validate(); err != nil {
		return fmt.Errorf("invalid notification settings: %v", err)
	}
	return nil
}

func (ms *MonitorSettings) Validate() error {
	// ExponentialBackoffSeconds can be 0, which means when calculated,
	// delay for notifications will always be 1 second
	if ms.CheckInterval <= 0 || ms.MaxConnections <= 0 || ms.Timeout <= 0 || ms.ExponentialBackoffSeconds < 0 {
		return fmt.Errorf("monitor settings missing")
	}
	return nil
}

func (ns *NotificationSettings) Validate() error {
	for _, email := range ns.Email {
		if err := email.Validate(); err != nil {
			return fmt.Errorf("invalid email settings: %v", err)
		}
	}
	for _, sms := range ns.Sms {
		if err := sms.Validate(); err != nil {
			return fmt.Errorf("invalid sms settings: %v", err)
		}
	}
	return nil
}

func (servers Servers) Validate() error {
	if len(servers) == 0 {
		return fmt.Errorf("no servers found in config")
	}

	for _, server := range servers {
		if err := server.Validate(); err != nil {
			return fmt.Errorf("invalid server settings: %s", err)
		}

	}
	return nil
}

func (s *Server) Validate() error {
	errServerProperty := func(property string) error {
		return fmt.Errorf("missing server property %s", property)
	}
	switch {
	case s.Name == "":
		return errServerProperty("name")
	case s.IPAddress == "":
		return errServerProperty("ipAddress")
	case s.Port == 0:
		return errServerProperty("port")
	case s.Protocol == "":
		return errServerProperty("protocol")
	}
	return nil
}
