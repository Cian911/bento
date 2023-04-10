package main

import (
	"flag"
	"log"
	"time"

	"github.com/cian911/blauberg-vento/pkg/config"
	"github.com/cian911/blauberg-vento/pkg/fan"
	"github.com/cian911/blauberg-vento/pkg/influxdb"
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
  parsedConfig := config.ParseConfig(configFile)
  stream := influxdb.NewClient(parsedConfig.S)

  var fans []*fan.Fan 
  for _, v := range parsedConfig.F.Fans { 
    f := fan.NewFan( 
      v.IPAddress, 
      v.ID, 
      v.Password, 
      v.Port, 
    ) 
    fans = append(fans, f) 
  } 
    
  // Test fan is connected and working 
  //f2 := fans[1] 
  //f2.ChangeFanSpeed(HIGH_SPEED) 
  
  // Poll stream every 5 seconds to get co2 data

}


