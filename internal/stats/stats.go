package stats

import (
	"sort"
	"time"
)

func CalculateStats(samples []time.Duration) (p50, p95, p99 time.Duration) {
	if len(samples) == 0 {
		return 0, 0, 0
	}
	sorted := make([]time.Duration, len(samples))
	copy(sorted, samples)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	p := func(percentile float64) time.Duration {
		index := int(float64(len(sorted)) * percentile)
		if index >= len(sorted) {
			index = len(sorted) - 1
		}
		return sorted[index]
	}

	return p(0.50), p(0.95), p(0.99)
}
