package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strconv"
)

const (
  // Turn on Ventilation
  unit_operation_request = "0001"
  on_value = "01"

  // Turn off Ventilation
  off_value = "00"

  // Extract air temperature
  extract_air_temp_request = "0020"

  // Wifi stuff
  wifi_client_name_request = "0095"
  wifi_operation_mode_request = "0094"
  wifi_status_request = "00A1"
  ip_address_assigned_to_module_request = "00A3"

  // Speed mode
  speed_mode_request = "0002"

  // All possible functions
  write_return = "03"
  read = "01"
  write = "02"
  inc = "04"
  dec = "05"

  HEADER = "FDFD"
)

func main() {
  //encodeData("03", on_request, on_value)
  //conn := connect()
  //receive(conn)
  fmt.Println(fmt.Sprintf("%x", len("004F00384B435705")))
  
  fmt.Println("-----------------------------------")
  //encodeData(write_return, on_request, on_value)
  encodeData(read, "0093", "")
}

func connect() *net.UDPConn {
  // Hostname is: ESP-4605B8
  udpServer, err := net.ResolveUDPAddr("udp", "192.168.0.72:4000")
  if err != nil {
    log.Fatal(err)
  }

  conn, err := net.DialUDP("udp", nil, udpServer)
  if err != nil {
    log.Fatal(err)
  }

  return conn
}

func send(data string, conn *net.UDPConn) int {
  payload := build_headers() + data
  fmt.Println("Payload: "+payload)
  payload = HEADER + payload + checksum(payload)
  fmt.Printf("payload: %v\n\n", payload)
  bytes, err := hex.DecodeString(payload)
  if err != nil {
    log.Fatalf("Could not decode payload: %v", err)
  }
  fmt.Println(bytes)
  res, err := conn.Write(bytes)
  if err != nil {
    log.Fatal(err)
  }
  
  buf := make([]byte, 64)
  _, _ = conn.Read(buf)

  fmt.Println(buf)
  s := string(buf[25:])
  s1 := string(buf)
  fmt.Printf("\n%v\n", s1)
  fmt.Printf("\n%v\n", s)

  return res
}

func build_headers() string {
  id_size := get_size("004F00384B435705")
  pwd_size := get_size("1111")
  id := fmt.Sprintf("%x", "004F00384B435705")
  password := fmt.Sprintf("%x", "1111")
  fmt.Println(fmt.Sprintf("PWD: %x", "1111"))
  devieType := "02"
  return fmt.Sprintf("%s%s%s%s%s", devieType, id_size, id, pwd_size, password)
}

func get_size(str string) string {
  encoding := fmt.Sprintf("%x", len(str))
  res, err := strconv.Atoi(encoding)
  if err != nil {
    log.Fatal(err)
  }

  t := fmt.Sprintf("%02d", res)

  return t
}

func encodeData(operation, param, value string) {
  out := ""
  parameter := ""
  val_bytes := 0

  for i := 0; i < len(param); i += 4{
    n_out := ""
    out = param[i : (i + 4)]
    if out == "0077" && value == "" {
      value = "0101"
    }
    if value != "" {
      val_bytes = int(len(value) / 2)
    } else {
      val_bytes = 0
    }
    if out[:2] != "00" {
      n_out = "ff" + out[:2]
    }
    if val_bytes > 1 {
      n_out += "fe" + fmt.Sprintf("%02x", val_bytes) + out[2:4]
    } else {
      n_out += out[2:4]
    }

    fmt.Printf("Out: %s, Param: %s Out2: %s val_bytes: %d\n", param, out, out[:2], val_bytes)

    parameter += n_out + value
    if out == "0077" {
      value = ""
    }
  }

  data := operation + parameter
  conn := connect()
  send(data, conn)
}

func checksum(msg string) string {
  chksum := fmt.Sprintf("%04x", sum(hexToTuple(msg)))
  byte_array, _ := hex.DecodeString(chksum)
  ck := fmt.Sprintf("%02x", byte_array[1]) + fmt.Sprintf("%02x", byte_array[0])
  fmt.Println(ck)
  return ck
}

func hexToTuple(msg string) []int64 {
  result := []int64{}
  val := int64(0)
  for i := 0; i < len(msg); i += 2 {
    if (i + 2) > len(msg) {
      val = 1
    } else {
      val, _ = strconv.ParseInt(msg[i : (i + 2)], 16, 16)
    }

    result = append(result, val)
  }

  return result
}

func sum(arr []int64) int {
    sum := int64(0)
    for _, valueInt := range arr {
        sum += valueInt
    }
    return int(sum)
}
