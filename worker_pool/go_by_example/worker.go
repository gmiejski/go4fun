package go_by_example

import (
	"org.miejski/go4fun/worker_pool/common"
	"fmt"
	"log"
)

type worker struct {
	jobs    chan common.Payload
	counter common.JobCounter
}

func (w *worker) start() {
	go func() {
		for {
			job := <-w.jobs
			if err := job.LongLastingOperation(); err != nil {
				log.Print(fmt.Sprintf("Error uploading to S3: %s", err.Error()))
			}
			w.counter.Done(job)
		}
	}()
}

func NewWorker(jobs chan common.Payload, counter common.JobCounter) *worker {
	return &worker{jobs: jobs, counter: counter}
}
