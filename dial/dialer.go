package dial

import (
	"net"
)

// Dialer is used to test connections
type Dialer struct {
	semaphore chan struct{}
}

// Status saves information about connection
type Status struct {
	Ok  bool
	Err error
}

// NewDialer returns pointer to new Dialer
func NewDialer(concurrentConnections int) *Dialer {
	return &Dialer{
		semaphore: make(chan struct{}, concurrentConnections),
	}
}

// NewWorker is used to send address over NetAddressTimeout to make request and receive status over DialerStatus
// Blocks until slot in semaphore channel for concurrency is free
func (d *Dialer) NewWorker() (chan<- NetAddressTimeout, <-chan Status) {
	netAddressTimeoutCh := make(chan NetAddressTimeout)
	dialerStatusCh := make(chan Status)

	d.semaphore <- struct{}{}
	go func() {
		netAddressTimeout := <-netAddressTimeoutCh
		conn, err := net.DialTimeout(netAddressTimeout.Network, netAddressTimeout.Address, netAddressTimeout.Timeout)

		dialerStatus := Status{}

		if err != nil {
			dialerStatus.Ok = false
			dialerStatus.Err = err
		} else {
			dialerStatus.Ok = true
			conn.Close()
		}
		dialerStatusCh <- dialerStatus
		<-d.semaphore
	}()

	return netAddressTimeoutCh, dialerStatusCh
}
