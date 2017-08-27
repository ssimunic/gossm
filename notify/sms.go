package notify

type SmsSettings struct {
	Sms string `json:"sms"`
}

type SmsNotifier struct {
	Settings *SmsSettings
}

func (s *SmsNotifier) Notify(text string) error {
	// TODO
	return nil
}

func (ss *SmsSettings) Validate() error {
	// TODO
	return nil
}

func (s *SmsNotifier) String() string {
	// TODO
	return ""
}
