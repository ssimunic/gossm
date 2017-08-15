package gossm

import (
	"fmt"
)

// Validator is used to validate fields of config structure
type Validator interface {
	// Validate returns <true, nil> if valid, <false, error> if invalid
	Validate() (bool, error)
}

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
	if ms.CheckInterval <= 0 || ms.MaxConnections <= 0 || ms.Timeout <= 0 {
		return false, fmt.Errorf("monitor settings missing")
	}
	return true, nil
}

func (ns *NotificationSettings) Validate() (bool, error) {
	for _, email := range ns.Email {
		if ok, err := email.Validate(); !ok {
			return false, err
		}
	}
	for _, sms := range ns.Sms {
		if ok, err := sms.Validate(); !ok {
			return false, err
		}
	}
	return true, nil
}

func (es *EmailSettings) Validate() (bool, error) {
	if es.Password == "" || es.Username == "" || es.SMTP == "" || es.Port == 0 {
		return false, fmt.Errorf("missing email settings")
	}
	return true, nil
}

func (ss *SmsSettings) Validate() (bool, error) {
	// TODO
	return true, nil
}

func (s *Servers) Validate() (bool, error) {
	if len(*s) == 0 {
		return true, fmt.Errorf("no servers found in config")
	}
	for _, server := range *s {
		if server.Name == "" || server.IPAddress == "" || server.Port == 0 || server.Protocol == "" {
			return true, fmt.Errorf("missing data for server: %#v", server)
		}
	}
	return true, nil
}
