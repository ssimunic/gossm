package notify

import (
	"fmt"
	"net/http"
	"net/url"
)

type WebhookSettings struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

type WebhookNotifier struct {
	Settings *WebhookSettings
}

func (w *WebhookNotifier) Notify(srv string) error {
	switch w.Settings.Method {
	case "GET":
		u, err := url.Parse(w.Settings.Url)
		if err != nil {
			return err
		}
		u.Query().Add("server", srv)
		_, err = http.Get(u.String())
		return err
	case "POST":
		values := url.Values{}
		values.Add("server", srv)
		_, err := http.PostForm(w.Settings.Url, values)
		return err
	default:
		return fmt.Errorf("Invalid Method %s", w.Settings.Method)
	}
}

func (w *WebhookNotifier) Initialize() {
}

func (ws *WebhookSettings) Validate() error {
	errWebhookProperty := func(property string) error {
		return fmt.Errorf("missing webhook property %s", property)
	}
	switch {
	case ws.Url == "":
		return errWebhookProperty("url")
	case ws.Method == "":
		return errWebhookProperty("method")
	}
	return nil
}

func (w *WebhookNotifier) String() string {
	return fmt.Sprintf("Webhook with method %s to URL %s", w.Settings.Method, w.Settings.Url)
}
