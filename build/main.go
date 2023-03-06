package main

import (
    "fmt"
    "net/http"
    "time"

    "math/rand"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

const PROMETHEUS_NAMESPACE = "pet_shop"

const MAX_NUMBER_OF_PETS = 15

func recordMetrics() {
    go func() {
        for {
            if ( rand.Intn(2) == 1 && petsInQueue != MAX_NUMBER_OF_PETS ) { // Enter a client
                petsInQueueMetric.Add(1)
            }else{ // Get out the client
                if ( petsInQueue > 0 ) {
                    petsInQueueMetric.Sub(1)
                    totalPetsServedMetric.Add(1)
                    var nota float64
                    nota = rand.NormFloat64() * 1.5 + 4
                    if ( nota > 5 ) {
                        nota = 5
                    }else if ( nota < 1 ) {
                        nota = 1
                    }
                    npsHistogram.Observe(nota)
                    fmt.Println(nota)
                }
            }
            time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
        }
    }()
}

var (
    petsInQueueMetric = prometheus.NewGauge(prometheus.GaugeOpts{
        Namespace: PROMETHEUS_NAMESPACE,
        Name: "pets_in_queue",
        Help: "number of pets to be cared for",
    })
    totalPetsServedMetric = prometheus.NewCounter( prometheus.CounterOpts{
        Namespace: PROMETHEUS_NAMESPACE,
        Name: "total_pets_served",
        Help: "number of pets served",
    })
    npsHistogram = prometheus.NewHistogram( prometheus.HistogramOpts{
        Namespace: PROMETHEUS_NAMESPACE,
        Name: "score",
        Help: "service's score",
        Buckets: []float64{1,2,3,4,5},
    })
    petsInQueue = 4
)

func init() {
    petsInQueueMetric.Add(float64(petsInQueue))
    prometheus.Register(petsInQueueMetric)
    prometheus.Register(totalPetsServedMetric)
    prometheus.Register(npsHistogram)
}

func main() {
    recordMetrics()

    fmt.Println("metrics in 0.0.0.0:3000/metrics")
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":3000", nil)
}

