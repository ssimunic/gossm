package notify

type SlackSettings struct {
	Slack string `json:"slack"`
}

type SlackNotifier struct {
	Settings *SlackSettings
}

func (s *SlackNotifier) Notify(text string) error {
	// TODO
	return nil
}

func (ss *SlackSettings) Validate() error {
	// TODO
	return nil
}

func (s *SlackNotifier) String() string {
	// TODO
	return ""
}
