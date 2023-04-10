package config

import (
	"log"

	"github.com/cian911/blauberg-vento/pkg/fan"
	"github.com/cian911/blauberg-vento/pkg/influxdb"
	"github.com/spf13/viper"
)

type Config struct {
  F fan.Fans
  S influxdb.InfluxDbClient
}

func ParseConfig(configFile string) Config {
  viper.SetConfigFile(configFile)
  var c Config

  if err := viper.ReadInConfig(); err == nil {
    log.Println("Using config file: ", viper.ConfigFileUsed())

    err := viper.Unmarshal(&c)

    if err != nil {
      log.Fatalf("Unable to decode config file. Please check data is in the correct format: %v", err)
    }

    if c.F.Fans == nil || len(c.F.Fans) == 0 {
      log.Fatalf("Unable to decode config file. Please check data is in the correct format: %v", c.F.Fans)
    }

    if c.S.Url == "" {
      log.Fatal("Unable to decode config file. Please check data is in the correct format.")
    }
  }

  return c
}
