package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log-analyzer/collector"
	"log-analyzer/internal/stats"
	"log-analyzer/middleware"
	"math/rand"
	"net/http"
	"time"
)

func addCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		addCORS(w)
		start := time.Now()
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fmt.Fprintln(w, "Hello, Latency Lens!")
		duration := time.Since(start)
		collector.Record("GET /hello", duration)
	})

	http.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		addCORS(w)
		start := time.Now()
		time.Sleep(time.Duration(300+rand.Intn(300)) * time.Millisecond)
		fmt.Fprintln(w, "This is a slow endpoint.")
		duration := time.Since(start)
		collector.Record("GET /slow", duration)
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		addCORS(w)
		raw := collector.GetMetrics()

		type Stat struct {
			Label string  `json:"label"`
			Count int     `json:"count"`
			P50   float64 `json:"p50"`
			P95   float64 `json:"p95"`
			P99   float64 `json:"p99"`
		}

		var output []Stat
		for label, data := range raw {
			log.Printf("[metrics] %s: count=%d samples=%v\n", label, data.Count, data.Samples)
			p50, p95, p99 := stats.CalculateStats(data.Samples)
			output = append(output, Stat{
				Label: label,
				Count: data.Count,
				P50:   float64(p50.Milliseconds()),
				P95:   float64(p95.Milliseconds()),
				P99:   float64(p99.Milliseconds()),
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(output)
	})

	wrappedMux := middleware.HTTPMiddleware(http.DefaultServeMux)

	fmt.Println("[server] Listenig on :3000")
	err := http.ListenAndServe(":3000", wrappedMux)
	if err != nil {
		log.Fatal(err)
	}
}
