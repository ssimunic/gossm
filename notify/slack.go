package notify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type SlackSettings struct {
	BearerToken string `json:"bearerToken"`
	ChannelID   string `json:"channelId"`
}

type SlackNotifier struct {
	Settings *SlackSettings
}

func (s *SlackNotifier) Notify(text string) error {
	payload := map[string]interface{}{"channel": s.Settings.ChannelID, "text": "Hello.. :tada:"}
	bytes, _ := json.Marshal(payload)

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, "https://slack.com/api/chat.postMessage", strings.NewReader(string(bytes)))
	r.Header.Add("Authorization", "Bearer "+s.Settings.BearerToken)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(string(bytes))))

	_, err := client.Do(r)
	return err
}

func (s *SlackNotifier) Initialize() {
}

func (ss *SlackSettings) Validate() error {
	errSlackProperty := func(property string) error {
		return fmt.Errorf("missing slack property %s", property)
	}
	switch {
	case ss.BearerToken == "":
		return errSlackProperty("bearerToken")
	case ss.ChannelID == "":
		return errSlackProperty("channelId")
	}
	return nil
}

func (s *SlackNotifier) String() string {
	return fmt.Sprintf("Slack Bot %s on ChannelID %s", s.Settings.BearerToken, s.Settings.ChannelID)
}
