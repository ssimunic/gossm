package gossm

import (
	"fmt"
	"net/smtp"

	"github.com/ssimunic/gossm/logger"
)

type Notifier interface {
	Notify(text string)
}

type Initializer interface {
	Initialize()
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

func (e *EmailNotifier) Notify(text string) {
	msg := "From: " + e.settings.Username + "\n" +
		"To: " + e.settings.Username + "\n" +
		"Subject: GOSSM Notification\n\n" +
		text + " not reached."

	logger.Logln("Sending email notification:", text)
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", e.settings.SMTP, e.settings.Port),
		e.auth,
		e.settings.Username,
		[]string{e.settings.Username},
		[]byte(msg),
	)
	if err != nil {
		logger.Logln(err)
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

func (s *SmsNotifier) Notify(text string) {
	logger.Logln("Sending sms notification:", text)
	// TODO
}
