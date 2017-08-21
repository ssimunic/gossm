package notify

type SmsSettings struct {
	Sms string `json:"sms"`
}

type SmsNotifier struct {
	Settings *SmsSettings
}

func (s *SmsNotifier) Notify(text string) (bool, error) {
	// TODO
	return false, nil
}

func (ss *SmsSettings) Validate() (bool, error) {
	// TODO
	return true, nil
}

func (s *SmsNotifier) String() string {
	// TODO
	return ""
}
