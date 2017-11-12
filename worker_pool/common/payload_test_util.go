package common

import (
	"time"
	"math/rand"
	"github.com/spf13/afero"
	"path/filepath"
	"os"
	"bufio"
)

const basePath = "/tmp/go4fun/worker_pool/"

func GeneratePayloadTimes(count int, minWaitMiliss int, maxWaitMiliss int) []time.Duration {
	var results []time.Duration
	for i := 0; i < count; i++ {
		var i = time.Millisecond * time.Duration(rand.Intn(maxWaitMiliss-minWaitMiliss)+minWaitMiliss)
		results = append(results, i)
	}
	return results
}

func Save(durations []time.Duration, filename string, fs afero.Fs) (string, error) {
	err := fs.MkdirAll(basePath, 0777)
	if err != nil {
		return "", err
	}
	resultFileName := filepath.Join(basePath, filename)
	file, err := fs.OpenFile(resultFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return "", err
	}
	for _, duration := range durations {
		_, err := file.WriteString(duration.String() + "\n")
		if err != nil {
			return "", err
		}
	}
	err = file.Close()
	if err != nil {
		return "", err
	}
	return resultFileName, nil
}

func ReadPayloadsFromFile(fullFilePath string, fs afero.Fs) ([]time.Duration, error) {
	fileHandle, err := fs.Open(fullFilePath)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(fileHandle)

	var durationsResult []time.Duration

	for fileScanner.Scan() {
		dur, err := time.ParseDuration(fileScanner.Text())
		if err != nil {
			return nil, err
		}
		durationsResult = append(durationsResult, dur)
	}

	return durationsResult, nil
}

func ReadOrSave(filename string, fs afero.Fs, generatingFunc func() []time.Duration) ([]time.Duration, error) {
	durations, err := ReadPayloadsFromFile(filepath.Join(basePath, filename), fs)
	if err == nil {
		return durations, nil
	}

	durations = generatingFunc()

	_, err = Save(durations, filename, fs)
	if err != nil {
		return nil, err
	}

	return durations, nil
}
