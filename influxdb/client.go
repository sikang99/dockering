package main

import (
	"fmt"

	"github.com/rossdylan/influxdbc"
)

func main() {
	database := influxdbc.NewInfluxDB("localhost:8086", "testdb", "root", "root")
	series := influxdbc.NewSeries("Name", "Col1", "Col2")
	series.AddPoint("Col1 data", "Col2 data")
	err := database.WriteSeries([]influxdbc.Series{*series})
	if err != nil {
		fmt.Println(err)
	}
}
