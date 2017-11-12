package go_by_example

import "org.miejski/go4fun/worker_pool/common"

type jobDoer struct {
	jobs chan common.Payload
}

func (jd *jobDoer) Do(payload common.Payload) {
	jd.jobs <- payload
}

func NewJobDoer(jobsQueueSize int, workersCount int, counter common.JobCounter) common.JobDoer {
	jobs := make(chan common.Payload, jobsQueueSize)

	for i := 0; i < workersCount; i++ {
		w := NewWorker(jobs, counter)
		w.start()
	}

	return &jobDoer{jobs: jobs}
}
