package gossm

type Settings struct {
	Monitor       *MonitorSettings
	Notifications *NotificationSettings
}

type MonitorSettings struct {
	CheckInterval  int `json:"checkInterval"`
	Timeout        int `json:"timeout"`
	MaxConnections int `json:"maxConnections"`
}

type NotificationSettings struct {
	Email []*EmailSettings `json:"email"`
	Sms   []*SmsSettings   `json:"sms"`
}

type EmailSettings struct {
	SMTP     string   `json:"smtp"`
	Port     int      `json:"port"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	From     string   `json:"from"`
	To       []string `json:"to"`
}

type SmsSettings struct {
	Sms string `json:"sms"`
}

func (n *NotificationSettings) GetNotifiers() (notifiers []Notifier) {
	for _, email := range n.Email {
		emailNotifier := &EmailNotifier{settings: email}
		notifiers = append(notifiers, emailNotifier)
	}
	for _, sms := range n.Sms {
		smsNotifier := &SmsNotifier{settings: sms}
		notifiers = append(notifiers, smsNotifier)
	}
	return
}
