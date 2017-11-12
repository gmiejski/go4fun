package marcio

import (
	"log"
	"fmt"
	"org.miejski/go4fun/worker_pool/common"
)



// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan common.Payload
	JobChannel chan common.Payload
	quit       chan bool
	jobCounter common.JobCounter
}

func NewWorker(workerPool chan chan common.Payload, jobCounter common.JobCounter) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan common.Payload),
		quit:       make(chan bool),
		jobCounter: jobCounter}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w *Worker) Start() { // TODO change for pointer receiver and compare performance
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				if err := job.LongLastingOperation(); err != nil {
					log.Print(fmt.Sprintf("Error uploading to S3: %s", err.Error()))
				}
				w.jobCounter.Done(job)
			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
