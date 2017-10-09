package gossm

import (
	"github.com/ssimunic/gossm/notify"
)

type Settings struct {
	Monitor       *MonitorSettings
	Notifications *NotificationSettings
}

type MonitorSettings struct {
	CheckInterval             int `json:"checkInterval"`
	Timeout                   int `json:"timeout"`
	MaxConnections            int `json:"maxConnections"`
	ExponentialBackoffSeconds int `json:"exponentialBackoffSeconds"`
}

type NotificationSettings struct {
	Email    []*notify.EmailSettings    `json:"email"`
	Sms      []*notify.SmsSettings      `json:"sms"`
	Telegram []*notify.TelegramSettings `json:"telegram"`
	Pushover []*notify.PushoverSettings `json:"pushover"`
	Webhook  []*notify.WebhookSettings  `json:"webhook"`
}

func (n *NotificationSettings) GetNotifiers() (notifiers notify.Notifiers) {
	for _, email := range n.Email {
		emailNotifier := &notify.EmailNotifier{Settings: email}
		notifiers = append(notifiers, emailNotifier)
	}
	for _, sms := range n.Sms {
		smsNotifier := &notify.SmsNotifier{Settings: sms}
		notifiers = append(notifiers, smsNotifier)
	}

	for _, telegram := range n.Telegram {
		telegramNotifier := &notify.TelegramNotifier{Settings: telegram}
		notifiers = append(notifiers, telegramNotifier)
	}
	for _, pushover := range n.Pushover {
		pushoverNotifier := &notify.PushoverNotifier{Settings: pushover}
		notifiers = append(notifiers, pushoverNotifier)
	}
	for _, webhook := range n.Webhook {
		webhookNotifier := &notify.WebhookNotifier{Settings: webhook}
		notifiers = append(notifiers, webhookNotifier)
	}
	return
}
