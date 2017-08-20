package dial

import (
	"time"
)

// NetAddress is transport unit for Dialer
type NetAddress struct {
	Network string
	Address string
}

// NetAddressTimeout is tuple of NetAddress and attached Timeout
type NetAddressTimeout struct {
	NetAddress
	Timeout time.Duration
}
