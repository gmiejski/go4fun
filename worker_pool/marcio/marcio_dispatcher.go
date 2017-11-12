package marcio

import "org.miejski/go4fun/worker_pool/common"

type MarcioDispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan common.Payload
	maxWorkers int
	jobQueue   chan common.Payload
}

func NewDispatcher(jobQueue chan common.Payload, maxWorkers int) *MarcioDispatcher {
	pool := make(chan chan common.Payload, maxWorkers)
	return &MarcioDispatcher{WorkerPool: pool, maxWorkers: maxWorkers, jobQueue: jobQueue}
}

func (d *MarcioDispatcher) Run(jobCounter common.JobCounter) {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, jobCounter)
		worker.Start()
	}

	go d.dispatch()
}

func (d *MarcioDispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			// a job request has been received
			go func(job common.Payload) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
