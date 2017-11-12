package common

import (
	"testing"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestPayloadGenerationAndReading(t *testing.T) {
	// given
	fs := afero.NewMemMapFs()
	generatedWaitTimes := GeneratePayloadTimes(1000, 20, 100)
	filename, err := Save(generatedWaitTimes, "testFile", fs)
	assert.NoError(t, err)
	// when
	payloads, err := ReadPayloadsFromFile(filename, fs)

	// then
	assert.NoError(t, err)
	assert.Equal(t, generatedWaitTimes, payloads)
}
