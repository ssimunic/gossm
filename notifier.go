package gossm

import (
	"github.com/ssimunic/gossm/logger"
	"time"
)

type Notifier interface {
	Notify(text string)
}

type Notifiers []Notifier

type EmailNotifier struct {
	settings *EmailSettings
}

type SmsNotifier struct {
	settings *SmsSettings
}

func (notifiers Notifiers) NotifyAll(text string) {
	for _, notifier := range notifiers {
		go notifier.Notify(text)
	}
}
func (e *EmailNotifier) Notify(text string) {
	logger.Logln("Sending email notification: ", text)
	time.Sleep(time.Second * 3)
}

func (s *SmsNotifier) Notify(text string) {
	logger.Logln("Sending sms notification: ", text)
	time.Sleep(time.Second * 3)
	// TODO
}
