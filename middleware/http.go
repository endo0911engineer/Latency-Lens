package middleware

import (
	"log-analyzer/collector"
	"net/http"
	"time"
)

// HTTPMiddleware wraps an http.Handler and records latency metrics.
func HTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(rw, r)

		latency := time.Since(start)
		collector.Record(r.Method+" "+r.URL.Path, latency)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
