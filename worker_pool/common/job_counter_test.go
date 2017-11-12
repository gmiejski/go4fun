package common

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func BenchmarkJobCounter(b *testing.B) {

	b.ReportAllocs()
	for benchmark := 0; benchmark < b.N; benchmark++ {
		b.StopTimer()
		expectedMessagesProcessed := 10000
		jc := NewJobCounter(expectedMessagesProcessed)

		timer := time.NewTimer(1 * time.Second)
		defer timer.Stop()

		b.StartTimer()
		jc.Start()

		for payload := 0; payload < expectedMessagesProcessed; payload++ {
			 jc.Done(&WaitingPayload{})
		}

		select {
		case value := <-jc.FinishedChannel():
			assert.EqualValues(b, expectedMessagesProcessed, value)
		case <-timer.C:
			assert.FailNow(b, "Timed out waiting for finishing jobs")
		}
	}
}
