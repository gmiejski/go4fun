package common

type JobCounter interface {
	Done(payload Payload)
	Count() int
	Start()
	FinishedChannel() chan int
}

func NewJobCounter(waitFor int) JobCounter {
	return &jobCounter{doneChannel: make(chan Payload, 100000), waitFor: waitFor, finishChannel: make(chan int, 1)}
}

type jobCounter struct {
	curr          int
	doneChannel   chan Payload
	started       bool
	waitFor       int
	finishChannel chan int
}

func (jc *jobCounter) FinishedChannel() chan int {
	return jc.finishChannel
}

func (jc *jobCounter) Start() {
	jc.started = true
	go func() {
		for {
			<-jc.doneChannel
			jc.curr += 1
			if jc.waitFor == jc.curr {
				jc.finishChannel <- jc.waitFor
			}
		}
	}()
}

func (jc *jobCounter) WaitFor(int) chan int {
	doneChannel := make(chan int)
	return doneChannel
}

func (jc *jobCounter) Count() int {
	return jc.curr
}

func (jc *jobCounter) Done(payload Payload) {
	if !jc.started {
		panic("Job counter not started!")
	}
	jc.doneChannel <- payload
}
