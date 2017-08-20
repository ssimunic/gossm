package gossm

import (
	"fmt"
)

func (c *Config) Validate() (bool, error) {
	if ok, err := c.Settings.Validate(); !ok {
		return false, fmt.Errorf("invalid settings: %v", err)
	}
	if ok, err := c.Servers.Validate(); !ok {
		return false, fmt.Errorf("invalid servers: %v", err)
	}
	return true, nil
}

func (s *Settings) Validate() (bool, error) {
	if ok, err := s.Monitor.Validate(); !ok {
		return false, fmt.Errorf("invalid monitor settings: %v", err)
	}
	if ok, err := s.Notifications.Validate(); !ok {
		return false, fmt.Errorf("invalid notification settings: %v", err)
	}
	return true, nil
}

func (ms *MonitorSettings) Validate() (bool, error) {
	// ExponentialBackoffSeconds can be 0, which means when calculated,
	// delay for notifications will always be 1 second
	if ms.CheckInterval <= 0 || ms.MaxConnections <= 0 || ms.Timeout <= 0 || ms.ExponentialBackoffSeconds < 0 {
		return false, fmt.Errorf("monitor settings missing")
	}
	return true, nil
}

func (ns *NotificationSettings) Validate() (bool, error) {
	for _, email := range ns.Email {
		if ok, err := email.Validate(); !ok {
			return false, fmt.Errorf("invalid email settings: %v", err)
		}
	}
	for _, sms := range ns.Sms {
		if ok, err := sms.Validate(); !ok {
			return false, fmt.Errorf("invalid sms settings: %v", err)
		}
	}
	return true, nil
}

func (es *EmailSettings) Validate() (bool, error) {
	errEmailProperty := func(property string) error {
		return fmt.Errorf("missing email property %s", property)
	}
	switch {
	case es.Username == "":
		return false, errEmailProperty("username")
	case es.Password == "":
		return false, errEmailProperty("password")
	case es.SMTP == "":
		return false, errEmailProperty("smtp")
	case es.Port == 0:
		return false, errEmailProperty("port")
	case es.From == "":
		return false, errEmailProperty("from")
	case len(es.To) == 0:
		return false, errEmailProperty("to")
	}
	return true, nil
}

func (ss *SmsSettings) Validate() (bool, error) {
	// TODO
	return true, nil
}

func (servers Servers) Validate() (bool, error) {
	if len(servers) == 0 {
		return false, fmt.Errorf("no servers found in config")
	}

	for _, server := range servers {
		if ok, err := server.Validate(); !ok {
			return false, fmt.Errorf("invalid server settings: %s", err)
		}

	}
	return true, nil
}

func (s *Server) Validate() (bool, error) {
	errServerProperty := func(property string) error {
		return fmt.Errorf("missing server property %s", property)
	}
	switch {
	case s.Name == "":
		return false, errServerProperty("name")
	case s.IPAddress == "":
		return false, errServerProperty("ipAddress")
	case s.Port == 0:
		return false, errServerProperty("port")
	case s.Protocol == "":
		return false, errServerProperty("protocol")
	}
	return true, nil
}
