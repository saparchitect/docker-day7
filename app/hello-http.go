package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "http",
			Name:      "requests",
			Help:      "This is my counter",
		})

	gauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "http",
			Name:      "gauge",
			Help:      "This is my gauge",
		})

	histogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "http",
			Name:      "latency",
			Help:      "This is my histogram",
		})
)

func randomMilliWait() {
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Millisecond * time.Duration(rand.Int31n(3000)))
}

func hello(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		histogram.Observe(time.Since(start).Seconds())
	}()
	randomMilliWait()
	counter.Add(1)
	gauge.Add(rand.Float64()*15 - 5)

	helloEnv := os.Getenv("HELLO")

	if helloEnv == "" {
		helloEnv = "hello"
	}

	arg := "friend"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	log.Printf("%s %s\n", helloEnv, arg)

	fmt.Fprintf(w, "%s %s\n", helloEnv, arg)
}

func main() {
	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", hello)
	log.Println("Starting server on 0.0.0.0:8090")
	http.ListenAndServe(":8090", nil)
}
