package locking

import "time"

func NewService(id string) Service {
	return Service{id: id}
}

type Service struct {
	id          string
	visitsCount int
	visitsTime  time.Duration
}

func (s *Service) visit(duration time.Duration) {
	s.visitsCount++
	s.visitsTime += duration
}
