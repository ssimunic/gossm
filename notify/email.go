package notify

import (
	"fmt"
	"net/smtp"
	"strings"
)

type EmailSettings struct {
	SMTP     string
	Port     int
	Username string
	Password string
	From     string
	To       []string
}

type EmailNotifier struct {
	Settings *EmailSettings
	auth     smtp.Auth
}

func (es *EmailSettings) Validate() error {
	errEmailProperty := func(property string) error {
		return fmt.Errorf("missing email property %s", property)
	}
	switch {
	case es.Username == "":
		return errEmailProperty("username")
	case es.Password == "":
		return errEmailProperty("password")
	case es.SMTP == "":
		return errEmailProperty("smtp")
	case es.Port == 0:
		return errEmailProperty("port")
	case es.From == "":
		return errEmailProperty("from")
	case len(es.To) == 0:
		return errEmailProperty("to")
	}
	return nil
}

func (e *EmailNotifier) Initialize() {
	e.auth = smtp.PlainAuth(
		"",
		e.Settings.Username,
		e.Settings.Password,
		e.Settings.SMTP,
	)
}

func (e *EmailNotifier) String() string {
	return fmt.Sprintf("email %s at %s:%d", e.Settings.Username, e.Settings.SMTP, e.Settings.Port)
}

func (e *EmailNotifier) Notify(text string) error {
	formattedReceipets := strings.Join(e.Settings.To, ", ")
	msg := "From: " + e.Settings.From + "\n" +
		"To: " + formattedReceipets + "\n" +
		"Subject: GOSSM Notification\n\n" +
		text + " not reached."

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", e.Settings.SMTP, e.Settings.Port),
		e.auth,
		e.Settings.From,
		e.Settings.To,
		[]byte(msg),
	)
	if err != nil {
		return fmt.Errorf("error sending email: %s", err)
	}
	return nil
}
