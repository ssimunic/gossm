package dialer

import (
	"net"
)

// Dialer is used to test connections
type Dialer struct {
	concurrentConnections int
	semaphore             chan struct{}
}

// Status saves information about connection
type Status struct {
	Ok  bool
	Err error
}

// New returns pointer to new Dialer
func New(concurrentConnections int) *Dialer {
	return &Dialer{
		concurrentConnections: concurrentConnections,
		semaphore:             make(chan struct{}, concurrentConnections),
	}
}

// NewWorker is used to send address over NetAddressTimeout to make request
// and receive status over DialerStatus
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
