package cli

import (
	"fmt"
	"log-analyzer/collector"
	"os"
	"sort"
	"strings"
	"time"
)

// percentile returns the nth percentile value from sorted durations.
func percentile(data []time.Duration, p float64) time.Duration {
	if len(data) == 0 {
		return 0
	}
	idx := int(float64(len(data)-1) * p)
	return data[idx]
}

// Visualize prints P50/P95/P99 latency per label.
func Visualize() {
	stats := collector.Snapshot()
	if len(stats) == 0 {
		fmt.Println("No latency data recorded.")
		os.Exit(0)
	}

	fmt.Printf("%-30s  %6s  %6s  %6s  %5s\n", "Label", "P50", "P95", "P99", "N")
	fmt.Println(strings.Repeat("-", 60))

	for label, samples := range stats {
		sort.Slice(samples, func(i, j int) bool {
			return samples[i] < samples[j]
		})
		p50 := percentile(samples, 0.50)
		p95 := percentile(samples, 0.95)
		p99 := percentile(samples, 0.99)
		fmt.Printf("%-30s  %6s  %6s  %6s  %5d\n", label, p50, p95, p99, len(samples))
	}
}
