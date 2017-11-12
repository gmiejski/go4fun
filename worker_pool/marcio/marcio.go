package marcio

import (
	"org.miejski/go4fun/worker_pool/common"
)

func NewMarcioJobDoer(dispatcherjobQueueSize int, workersCount int, counter common.JobCounter) common.JobDoer {
	jobQueue := make(chan common.Payload, dispatcherjobQueueSize)

	dispatcher := NewDispatcher(jobQueue, workersCount)
	dispatcher.Run(counter)
	return &marcioJobDoer{jobQueue: jobQueue}
}

type marcioJobDoer struct {
	jobQueue chan common.Payload
}

func (jd *marcioJobDoer) Do(payload common.Payload) {
	jd.jobQueue <- payload
}
