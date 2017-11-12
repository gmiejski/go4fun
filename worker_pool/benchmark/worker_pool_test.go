package benchmark

import (
	"testing"
	"org.miejski/go4fun/worker_pool/common"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/spf13/afero"
	"org.miejski/go4fun/worker_pool/go_by_example"
	"org.miejski/go4fun/worker_pool/marcio"
)

var (
	QueueSize = 20
	WorkerCount = 100
	MessagesCount = 1000
)

func BenchmarkGoByExample(b *testing.B) {

	b.StopTimer()
	payloadSet, err := common.ReadOrSave("worker_pool_test_data", afero.NewOsFs(), func() []time.Duration {
		return common.GeneratePayloadTimes(MessagesCount, 20, 100)
	})
	assert.NoError(b, err)

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		jobCounter := common.NewJobCounter(MessagesCount)
		jobCounter.Start()
		jd := go_by_example.NewJobDoer(QueueSize, WorkerCount, jobCounter)
		timer := time.NewTimer(10 * time.Second)
		b.StartTimer()

		for payload := 0; payload < MessagesCount; payload++ {
			jd.Do(&common.WaitingPayload{WaitingTime: payloadSet[payload]})
		}

		select {
		case value := <-jobCounter.FinishedChannel():
			assert.EqualValues(b, MessagesCount, value)
			timer.Stop()
		case <-timer.C:
			assert.FailNow(b, "Timed out waiting for finishing jobs")
		}
	}
}

func BenchmarkMarcioWorkerPool(b *testing.B) {
	b.StopTimer()
	payloadSet, err := common.ReadOrSave("worker_pool_test_data", afero.NewOsFs(), func() []time.Duration {
		return common.GeneratePayloadTimes(MessagesCount, 20, 100)
	})
	assert.NoError(b, err)

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		jobCounter := common.NewJobCounter(MessagesCount)
		jobCounter.Start()
		jd := marcio.NewMarcioJobDoer(QueueSize, WorkerCount, jobCounter)
		timer := time.NewTimer(10 * time.Second)
		b.StartTimer()

		for payload := 0; payload < MessagesCount; payload++ {
			jd.Do(&common.WaitingPayload{WaitingTime: payloadSet[payload]})
		}

		select {
		case value := <-jobCounter.FinishedChannel():
			assert.EqualValues(b, MessagesCount, value)
			timer.Stop()
		case <-timer.C:
			assert.FailNow(b, "Timed out waiting for finishing jobs")
		}
	}
}
