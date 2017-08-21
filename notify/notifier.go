package notify

// Initializer is used to initialize something
type Initializer interface {
	Initialize()
}

// Notifier is used to send messages
type Notifier interface {
	// Notify sends text over notifier which returns status and error message
	Notify(text string) (bool, error)
}

type Notifiers []Notifier

func (notifiers Notifiers) NotifyAll(text string) {
	for _, notifier := range notifiers {
		go notifier.Notify(text)
	}
}
