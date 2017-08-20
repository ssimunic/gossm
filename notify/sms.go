package notify

type SmsSettings struct {
	Sms string `json:"sms"`
}

type SmsNotifier struct {
	Settings *SmsSettings
}

func (s *SmsNotifier) Notify(text string) {
	// TODO
}

func (ss *SmsSettings) Validate() (bool, error) {
	// TODO
	return true, nil
}
