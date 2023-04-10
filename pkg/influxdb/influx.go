package influxdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cian911/blauberg-vento/pkg/fan"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxDbClient struct {
	Url         string
	Token       string
	Port        int
	Database    string
	Measurement string
	Field       string
	Interval    int
	Threshold   int

	c influxdb2.Client
}

func NewClient(client InfluxDbClient) *InfluxDbClient {
	client.c = influxdb2.NewClient(
		fmt.Sprintf("%s:%d", client.Url, client.Port),
		fmt.Sprintf("%s", client.Token),
	)

	return &client
}

func (i *InfluxDbClient) QueryField() interface{} {
	queryApi := i.c.QueryAPI("")

	result, err := queryApi.Query(
		context.Background(),
		fmt.Sprintf(`from(bucket:"%s")|> range(start: -%ds) |> filter(fn: (r) => r._measurement == "%s") |> filter(fn: (r) => r._field == "%s")`, i.Database, i.Interval, i.Measurement, i.Field),
	)

	if err != nil {
		log.Fatalf("Error connecting to InfluxDB: %v", err)
	}

	if result.Err() != nil {
		log.Fatalf("InfluxDB query error: %v", result.Err())
	}

  i.c.Close()

	for result.Next() {
		return result.Record().Value()
	}

	return nil
}

func (i *InfluxDbClient) Poll(f fan.Fan) {
  go func() {
    ticker := time.NewTicker(time.Duration(i.Interval) * time.Second)
    for {
      select {
      case <-ticker.C:
        log.Println("Polling...")
        co2level := i.QueryField().(int)
        
        if co2level >= i.Threshold {
          log.Println("Threshold meet, turning fans to full for 5 minutes.")
          f.ChangeFanSpeed("03")
        }
      }
    }
  }()
}
