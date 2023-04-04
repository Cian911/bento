package influxdb

import (
	"context"
	"fmt"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxDbClient struct {
	Url         string
	Token       string
	Port        int
	Database    string
	Measurement string
	Field       string

	c influxdb2.Client
}

func NewClient(url, token, database, measurement, field string, port int) *InfluxDbClient {
	client := &InfluxDbClient{
		Url:         url,
		Token:       token,
		Port:        port,
		Database:    database,
		Measurement: measurement,
		Field:       field,
	}
	client.c = influxdb2.NewClient(
		fmt.Sprintf("%s:%d", client.Url, client.Port),
		fmt.Sprintf("%s", client.Token),
	)

	return client
}

func (i *InfluxDbClient) QueryCo2Field() interface{} {
	queryApi := i.c.QueryAPI("")

	result, err := queryApi.Query(
		context.Background(),
		fmt.Sprintf(`from(bucket:"%s")|> range(start: -15s) |> filter(fn: (r) => r._measurement == "%s") |> filter(fn: (r) => r._field == "%s")`, i.Database, i.Measurement, i.Field),
	)

	if err != nil {
		log.Fatalf("Error connecting to InfluxDB: %v", err)
	}

	if result.Err() != nil {
		log.Fatalf("InfluxDB query error: %v", result.Err())
	}

	for result.Next() {
		return result.Record().Value()
	}

	return nil
}
