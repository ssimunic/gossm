package gossm

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/ssimunic/gossm/logger"
)

type Initializer interface {
	Initialize()
}

type Notifier interface {
	Notify(text string)
}

type Notifiers []Notifier

type EmailNotifier struct {
	settings *EmailSettings
	auth     smtp.Auth
}

type SmsNotifier struct {
	settings *SmsSettings
}

func (notifiers Notifiers) NotifyAll(text string) {
	for _, notifier := range notifiers {
		go notifier.Notify(text)
	}
}

func (notifiers Notifiers) NotifyAllWithDelay(text string, delay time.Duration) {
	<-time.After(delay)
	for _, notifier := range notifiers {
		go notifier.Notify(text)
	}
}

func (e *EmailNotifier) Initialize() {
	logger.Logln("Authenticating", e.settings.Username)
	e.auth = smtp.PlainAuth(
		"",
		e.settings.Username,
		e.settings.Password,
		e.settings.SMTP,
	)
}

func (e *EmailNotifier) Notify(text string) {
	formattedReceipets := strings.Join(e.settings.To, ", ")
	msg := "From: " + e.settings.From + "\n" +
		"To: " + formattedReceipets + "\n" +
		"Subject: GOSSM Notification\n\n" +
		text + " not reached."

	logger.Logln("Sending email notification:", text)
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", e.settings.SMTP, e.settings.Port),
		e.auth,
		e.settings.From,
		e.settings.To,
		[]byte(msg),
	)
	if err != nil {
		logger.Logln(err)
	}
}

func (s *SmsNotifier) Notify(text string) {
	logger.Logln("Sending sms notification:", text)
	// TODO
}
