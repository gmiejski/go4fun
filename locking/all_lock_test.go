package locking

import (
	"testing"
	"time"
	"fmt"
)

var services, inputs = generateInputs(100, time.Millisecond*10, time.Millisecond*30)

func TestLockingTime(t *testing.T) {
	manager := NewAllLockManager(services)
	tic := time.Now()
	for _, lock := range inputs {
		manager.lock(lock)
	}
	fmt.Println(time.Since(tic).String())
}
