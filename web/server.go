package main

import (
	"encoding/json"
	"log"
	"log-analyzer/collector"
	"log-analyzer/internal/stats"
	"net/http"
)

func main() {
	http.HandleFunc("/metrics", handleMetrics)
	http.Handle("/", http.FileServer(http.Dir("./ui")))

	log.Println("Serving latency dashboard on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	raw := collector.GetMetrics()
	log.Printf("[metrics] raw data: %+v\n", raw)

	type Stat struct {
		Label string `json:"label"`
		Count int    `json:"count"`
		P50   string `json:"p50"`
		P95   string `json:"p95"`
		P99   string `json:"p99"`
	}

	var output []Stat
	for label, data := range raw {
		log.Printf("[debug] %s: count=%d samples=%v\n", label, data.Count, data.Samples)

		p50, p95, p99 := stats.CalculateStats(data.Samples)
		output = append(output, Stat{
			Label: label,
			Count: data.Count,
			P50:   p50.String(),
			P95:   p95.String(),
			P99:   p99.String(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
