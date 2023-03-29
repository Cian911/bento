package main

import "github.com/cian911/blauberg-vento/pkg/fan"

const (
  LOW_SPEED = "01"
  MID_SPEED = "02"
  HIGH_SPEED = "03"
)

func main() {
/*  f := fan.NewFan(*/
    /*"192.168.0.72",*/
    /*"004F00384B435705",*/
    /*"1111",*/
    /*4000,*/
  /*)*/

  //f.ChangeFanSpeed(HIGH_SPEED)
  //f.ChangeFanOperation("invert")

  f1 := fan.NewFan(
    "192.168.0.238",
    "001900284B435704",
    "1111",
    4000,
  )

  f1.ChangeFanSpeed(HIGH_SPEED)
}

