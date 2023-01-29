package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strconv"
)

const (
  // Turn on Ventilation
  on_request = "0001"
  on_value = "01"

  // Turn off Ventilation
  off_request = "0001"
  off_value = "00"

  write_return = "03"

  HEADER = "FDFD"
)

func main() {
  //encodeData("03", on_request, on_value)
  //conn := connect()
  //receive(conn)
  fmt.Println(fmt.Sprintf("%x", len("004F00384B435705")))
  
  fmt.Println("-----------------------------------")
  encodeData(write_return, on_request, on_value)
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

func send(data string) int {
  //conn := connect()
  payload := build_headers() + data
  fmt.Println("Payload: "+payload)
  payload = HEADER + payload + checksum(payload)
  fmt.Printf("payload: %v\n\n", payload)
  bytes, err := hex.DecodeString(payload)
  if err != nil {
    log.Fatalf("Could not decode payload: %v", err)
  }
  fmt.Println(bytes)
  /*res, err := conn.Write(bytes)*/
  /*fmt.Printf("response: %v", res)*/
  /*if err != nil {*/
    /*log.Fatal(err)*/
  /*}*/
  /*return res*/
  return 1
}

func receive(conn *net.UDPConn) {
  b := make([]byte, 4096)
  _, err := bufio.NewReader(conn).Read(b)
  fmt.Println(b)

  if err != nil {
    log.Fatalf("Error reading client: %v", err)
  } else {
    fmt.Printf("Buff: %v\n", b)
  }

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

    parameter += n_out + value
    if out == "0077" {
      value = ""
    }
  }

  data := operation + parameter
  conn := connect()
  send(data)
  receive(conn)
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
