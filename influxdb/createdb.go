package main

import (
	"fmt"
	"time"

	client "github.com/influxdb/influxdb/client/v2"
)

const (
	MyDB     = "square_holes"
	username = "root"
	password = "root"
)

func main() {
	// Make a client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		fmt.Println(err)
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		fmt.Println(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		fmt.Println(err)
	}

	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)
}
