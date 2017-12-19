package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	helloCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hello_calls",
			Help: "Number of hello calls.",
		},
		[]string{"url"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(helloCounter)
}

func handler(w http.ResponseWriter, r *http.Request) {
	helloCounter.With(prometheus.Labels{"url": "/hello"}).Inc()
	fmt.Fprintf(w, "12 clouds demo at CloudAustin 2017!")
}

func main() {
	http.HandleFunc("/hello", handler)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
