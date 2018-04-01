package locking

import (
	"sync"
	"time"
)

func NewGrainedLock(services []Service) LockingManager {
	locks := generateLocks(services)
	serviceMap := toMap(services)
	return &grainedLock{locks: locks, services: serviceMap}
}

func toMap(services []Service) map[string]*Service {
	result := make(map[string]*Service)
	for _, service := range services {
		result[service.id] = &service
	}
	return result
}

func generateLocks(services []Service) map[string]*sync.Mutex {
	result := make(map[string]*sync.Mutex)
	for _, service := range services {
		result[service.id] = &sync.Mutex{}
	}
	return result
}

type grainedLock struct {
	locks    map[string]*sync.Mutex
	services map[string]*Service
}

func (lock *grainedLock) name() string {
	return "FineGrainedLocking"
}

func (lock *grainedLock) lock(rq LockServiceNow) {
	for _, x := range rq.services {
		lock.locks[x].Lock()
	}
	go lock.unlock(rq.services, rq.time)
}

func (lock *grainedLock) unlock(services []string, duration time.Duration) {
	time.Sleep(duration)
	for _, serviceName := range services {
		lock.services[serviceName].visit(duration)
	}
	for _, serviceName := range services {
		lock.locks[serviceName].Unlock()
	}
}
