package server

import (
	"github.com/prometheus/client_golang/prometheus/collectors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route", "status"},
	)

	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "route", "status"},
	)
)

func init() {
	prometheus.MustRegister(requestDuration, requestCount)
	prometheus.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{
		Namespace: "go-microservice-template-service",
	}))
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &statusRecorder{ResponseWriter: w, status: 200}
		start := time.Now()

		next.ServeHTTP(recorder, r)

		routePattern := chi.RouteContext(r.Context()).RoutePattern()
		duration := time.Since(start).Seconds()

		requestDuration.WithLabelValues(r.Method, routePattern, http.StatusText(recorder.status)).Observe(duration)
		requestCount.WithLabelValues(r.Method, routePattern, http.StatusText(recorder.status)).Inc()
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}
