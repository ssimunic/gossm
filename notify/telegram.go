package notify

import (
	"fmt"
	"net/http"
	"net/url"
)

type TelegramSettings struct {
	BotToken string `json:"botToken"`
	ChatID   string `json:"chatId"`
}

type TelegramNotifier struct {
	Settings *TelegramSettings
}

func (s *TelegramNotifier) Notify(text string) error {
	values := url.Values{}
	values.Add("chat_id", s.Settings.ChatID)
	values.Add("parse_mode", "markdown")
	values.Add("text", "*[Error]* _GOSSM_\nServer "+text+" not reached.")
	_, err := http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.Settings.BotToken), values)
	return err
}

func (e *TelegramNotifier) Initialize() {
}

func (ts *TelegramSettings) Validate() error {
	errTelegramProperty := func(property string) error {
		return fmt.Errorf("missing telegram property %s", property)
	}
	switch {
	case ts.BotToken == "":
		return errTelegramProperty("bot_token")
	case ts.ChatID == "":
		return errTelegramProperty("chat_id")
	}
	return nil
}

func (t *TelegramNotifier) String() string {
	return fmt.Sprintf("Telegram Bot %s with ChatID %s", t.Settings.BotToken, t.Settings.ChatID)
}
