package main

import (
	"net/http"
	"time"

	"github.com/GeertJohan/go-metrics/influxdb"
	"github.com/rcrowley/go-metrics"
)

func MetricToInfluxDB(d time.Duration) {
	go influxdb.Influxdb(metrics.DefaultRegistry, d, &influxdb.Config{
		Host:     "localhost:8086",
		Database: "metric",
		Username: "root",
		Password: "root",
	})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world for InfluxDB"))
}

func main() {
	MetricToInfluxDB(time.Second * 1)
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)
}
