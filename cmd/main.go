package main

import (
	"context"
	"flag"
	"fmt"

	influxClient "github.com/influxdata/influxdb-client-go/v2"
)

const (
	LOW_SPEED  = "01"
	MID_SPEED  = "02"
	HIGH_SPEED = "03"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "config", "", "Path to config file to parse.")
	flag.Parse()

	/*  yamlFans := config.ParseConfig(configFile) */
	/* var fans []*fan.Fan */
	/* for _, v := range yamlFans.Fans { */
	/*   f := fan.NewFan( */
	/*     v.IPAddress, */
	/*     v.ID, */
	/*     v.Password, */
	/*     v.Port, */
	/*   ) */
	/*   fans = append(fans, f) */
	/* } */
	/*  */
	/* // Test fan is connected and working */
	/* f2 := fans[1] */
	/* f2.ChangeFanSpeed(HIGH_SPEED) */

	c := influxClient.NewClient("http://192.168.0.172:8086", fmt.Sprintf("%s:%s", "user", "password"))
	queryApi := c.QueryAPI("")
	result, err := queryApi.Query(context.Background(), `from(bucket:"flink_home")|> range(start: -15s) |> filter(fn: (r) => r._measurement == "Co2Data") |> filter(fn: (r) => r._field == "co2")`)
	if err == nil {
		for result.Next() {
			fmt.Printf("row: %v\n", result.Record().Value())
		}
		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}

	c.Close()
}
