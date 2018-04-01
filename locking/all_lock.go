package locking

import (
	"sync"
	"time"
)

type allLockManager struct {
	access   sync.Mutex
	services []Service
}

func (lm *allLockManager) name() string {
	return "AllLock"
}

func (lm *allLockManager) lock(rq LockServiceNow) {
	lm.access.Lock()

	for _, x := range rq.services {
		for _, s := range lm.services {
			if s.id == x {
				s.visit(rq.time)
				break
			}
		}
	}
	time.Sleep(rq.time)
	lm.access.Unlock()
}

func NewAllLockManager(services []Service) LockingManager {
	return &allLockManager{access: sync.Mutex{}, services: services}
}
