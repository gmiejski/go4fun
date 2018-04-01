package locking

import (
	"testing"
	"time"
	"fmt"
)

var services, inputs = generateInputs(1000, time.Millisecond*50, time.Millisecond*100)

func TestLockingTime(t *testing.T) {
	manager := NewAllLockManager(services)
	tic := time.Now()
	fmt.Println(fmt.Sprintf("******** %s ********", manager.name()))
	for _, lock := range inputs {
		manager.lock(lock)
	}
	fmt.Println(fmt.Sprintf("Manager: %s -> time %s", manager.name(), time.Since(tic).String()))

	manager = NewGrainedLock(services)
	tic = time.Now()
	fmt.Println(fmt.Sprintf("******** %s ********", manager.name()))
	for _, lock := range inputs {
		manager.lock(lock)
	}
	fmt.Println(fmt.Sprintf("Manager: %s -> time %s", manager.name(), time.Since(tic).String()))
}
