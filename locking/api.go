package locking

import (
	"strconv"
	"time"
	"math/rand"
)

type LockingManager interface {
	lock(LockServiceNow)
	name() string
}

type LockServiceNow struct {
	services []string
	time     time.Duration
}

func generateInputs(servicesCount int, minTime time.Duration, maxTime time.Duration) ([]Service, []LockServiceNow) {
	services := generateServices(servicesCount)
	requests := generateRequests(minTime, maxTime, services)
	return services, requests
}

func generateRequests(min, max time.Duration, services []Service) []LockServiceNow {
	requests := make([]LockServiceNow, 0)

	for i := 0; i < 100; i++ {
		r := rand.Int63n(int64(max - min))
		lockedCount := rand.Intn(10) + 1
		thisRequestServices := make([]string, 0)
		permutation := rand.Perm(len(services))
		for i := 0; i < lockedCount; i++ {
			thisRequestServices = append(thisRequestServices, services[permutation[i]].id)
		}
		requests = append(requests, LockServiceNow{time: time.Duration(r), services: thisRequestServices})
	}
	return requests
}

func generateServices(servicesCount int) []Service {
	result := make([]Service, servicesCount)
	for i := 0; i < servicesCount; i++ {
		result[i] = NewService("service-" + strconv.Itoa(i))
	}
	return result
}
