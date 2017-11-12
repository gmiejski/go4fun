package common

import "time"

type Payload interface {
	LongLastingOperation() error
}

type WaitingPayload struct {
	WaitingTime time.Duration
}

func (d *WaitingPayload) LongLastingOperation() error {
	time.Sleep(d.WaitingTime)
	println("d")
	return nil
}
