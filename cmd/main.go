package main

import (
	"flag"
	"sync"

	"github.com/cian911/blauberg-vento/pkg/config"
	"github.com/cian911/blauberg-vento/pkg/fan"
	"github.com/cian911/blauberg-vento/pkg/influxdb"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "config", "", "Path to config file to parse.")
	flag.Parse()
	parsedConfig := config.ParseConfig(configFile)
	stream := influxdb.NewClient(parsedConfig.I)

	var fans []*fan.Fan
	for _, v := range parsedConfig.F {
		f := fan.NewFan(
			v.IPAddress,
			v.ID,
			v.Password,
			v.Name,
			v.Port,
			v.MaxedTimeout,
		)
		fans = append(fans, f)
	}

	// Test fan is connected and working
	f2 := fans[1]

	// Poll stream every 5 seconds to get co2 data
	var wg sync.WaitGroup
	wg.Add(1)
	stream.Poll(f2)
	wg.Wait()
}
