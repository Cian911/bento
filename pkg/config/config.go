package config

import (
	"log"

	"github.com/spf13/viper"
	"github.com/cian911/blauberg-vento/pkg/fan"
)

func ParseConfig(configFile string) fan.Fans {
  viper.SetConfigFile(configFile)
  var fans fan.Fans

  if err := viper.ReadInConfig(); err == nil {
    log.Println("Using config file: ", viper.ConfigFileUsed())

    err := viper.Unmarshal(&fans)

    if err != nil {
      log.Fatalf("Unable to decode config file. Please check data is in the correct format: %v", err)
    }

    if fans.Fans == nil || len(fans.Fans) == 0 {
      log.Fatalf("Unable to decode config file. Please check data is in the correct format: %v", fans.Fans)
    }
  }

  return fans
}
