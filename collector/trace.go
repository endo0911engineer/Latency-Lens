package collector

import (
	"log"
	"sync"
	"time"
)

type entry struct {
	Count   int
	Samples []time.Duration
}

var (
	mu      sync.Mutex
	metrics = make(map[string]*entry)
)

// Record stores the latency for a given label (e.g. route).
func Record(label string, d time.Duration) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := metrics[label]; !ok {
		metrics[label] = &entry{Samples: make([]time.Duration, 0, 100)}
	}
	metrics[label].Count++
	metrics[label].Samples = append(metrics[label].Samples, d)

	log.Printf("[recorded] %s: %v (count: %d)", label, d, metrics[label].Count)
}

// Snapshot returns a copy of the current metrics data.
func Snapshot() map[string][]time.Duration {
	mu.Lock()
	defer mu.Unlock()

	copyMap := make(map[string][]time.Duration, len(metrics))
	for label, e := range metrics {
		copyMap[label] = append([]time.Duration(nil), e.Samples...)
	}
	return copyMap
}

// GetMetrics returns a snapshot of all metrics.
func GetMetrics() map[string]entry {
	mu.Lock()
	defer mu.Unlock()

	// copy to avoid race conditions
	snapshot := make(map[string]entry)
	for k, v := range metrics {
		sampleCopy := make([]time.Duration, len(v.Samples))
		copy(sampleCopy, v.Samples)
		snapshot[k] = entry{
			Count:   v.Count,
			Samples: sampleCopy,
		}
	}
	return snapshot
}
