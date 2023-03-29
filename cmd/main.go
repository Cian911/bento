package main

import (
	"flag"

	"github.com/cian911/blauberg-vento/pkg/config"
	"github.com/cian911/blauberg-vento/pkg/fan"
)

const (
  LOW_SPEED = "01"
  MID_SPEED = "02"
  HIGH_SPEED = "03"
)

var configFile string

func main() {
  flag.StringVar(&configFile, "config", "", "Path to config file to parse.")  
  flag.Parse()

  yamlFans := config.ParseConfig(configFile)
  var fans []*fan.Fan
  for _, v := range yamlFans.Fans {
    f := fan.NewFan(
      v.IPAddress,
      v.ID,
      v.Password,
      v.Port,
    )
    fans = append(fans, f)
  }
  
  // Test fan is connected and working
  f2 := fans[1]
  f2.ChangeFanSpeed(HIGH_SPEED)
}

