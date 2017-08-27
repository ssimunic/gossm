package notify

// Initializer is used to initialize something
type Initializer interface {
	Initialize()
}

// Notifier is used to send messages
type Notifier interface {
	// Notify sends text over notifier, returns error message if failed
	Notify(text string) error
}

type Notifiers []Notifier

func (notifiers Notifiers) NotifyAll(text string) {
	for _, notifier := range notifiers {
		go notifier.Notify(text)
	}
}
