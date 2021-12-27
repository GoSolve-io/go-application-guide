package metrics

import (
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	bufferSize = 50
)

// Provider is a dummy metrics example
type Provider interface {
	Count(tags ...string) error
	Duration(duration time.Duration, tags ...string) error
	Flush()
}

// DummyMetrics implements some dummy metrics provider
type DummyMetrics struct {
	Counts    map[string]uint64
	Durations map[string][]time.Duration
	logger    logrus.FieldLogger
	sync.Mutex
}

// NewDummy returns a new dummy metrics
func NewDummy(logger logrus.FieldLogger) *DummyMetrics {
	return &DummyMetrics{
		logger: logger,
	}
}

// Count increases the count for given tags
func (dm *DummyMetrics) Count(tags ...string) error {
	if len(tags) == 0 {
		return ErrInvalidTags
	}

	dm.Lock()
	defer dm.Unlock()
	if dm.Counts == nil {
		dm.Counts = make(map[string]uint64)
	}

	key := strings.Join(tags, ".")
	if _, exists := dm.Counts[key]; !exists {
		dm.Counts[key] = 0
	}

	dm.Counts[key]++

	return nil
}

// Duration stores the duration for given tags, will log the average
func (dm *DummyMetrics) Duration(duration time.Duration, tags ...string) error {
	if len(tags) == 0 {
		return ErrInvalidTags
	}

	if duration == 0 {
		return ErrInvalidAttribute
	}

	dm.Lock()
	defer dm.Unlock()
	if dm.Durations == nil {
		dm.Durations = make(map[string][]time.Duration)
	}

	key := strings.Join(tags, ".")
	if _, exists := dm.Durations[key]; !exists {
		dm.Durations[key] = make([]time.Duration, 0, bufferSize)
	}

	dm.Durations[key] = append(dm.Durations[key], duration)

	return nil
}

// Flush flushed the stored data to stdout
func (dm *DummyMetrics) Flush() {
	dm.Lock()
	defer dm.Unlock()

	for key, value := range dm.Counts {
		dm.logger.Printf("Count: %s - %d\n", key, value)
		delete(dm.Counts, key)
	}

	for key, durations := range dm.Durations {
		var avg int64
		for _, d := range durations {
			avg = avg + d.Nanoseconds()
		}
		avg = avg / int64(len(durations))
		dm.logger.Printf("Duration: %s - %d\n", key, avg)
		delete(dm.Durations, key)
	}
}
