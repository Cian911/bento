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

	c := influxClient.NewClient("http://192.168.0.172:8086", fmt.Sprintf("%s:%s", "root", "password"))
	queryApi := c.QueryAPI("")
	_, err := queryApi.Query(context.Background(), `from(bucket:"Co2Data")|> range(start: -15s)`)
	/* if err == nil { */
	/*   for result.Next() { */
	/*     if result.TableChanged() { */
	/*       fmt.Printf("table: %s\n", result.TableMetadata().String()) */
	/*     } */
	/*  */
	/*     fmt.Printf("row: %s\n", result.Record().String()) */
	/*   } */
	/*   if result.Err() != nil { */
	/*     fmt.Printf("Query error: %s\n", result.Err().Error()) */
	/*   } */
	/* } else { */
	/*   panic(err) */
	/* } */

	fmt.Printf("Error: %v", err)

	c.Close()
}
