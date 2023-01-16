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
  encodeData("03", off_request, off_value)
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
  conn := connect()
  payload := build_headers() + data
  payload = HEADER + payload + checksum(payload)
  bytes, _ := hex.DecodeString(payload)
  //addr := &net.UDPAddr{Port: 4000, IP: net.ParseIP("192.168.0.88")}
  res, err := conn.Write(bytes)
  if err != nil {
    log.Fatal(err)
  }
  return res
}

func receive(conn *net.UDPConn) {
  buf := make([]byte, 4096)
  _, err := bufio.NewReader(conn).Read(buf)

  if err != nil {
    log.Fatalf("Error reading client: %v", err)
  } else {
    fmt.Printf("%s\n", buf)
  }

}

func build_headers() string {
  id_size := get_size("004F00384B435705")
  pwd_size := get_size("1111")
  id := hex.EncodeToString([]byte("004F00384B435705"))
  password := hex.EncodeToString([]byte("1111"))
  devieType := "02"
  return fmt.Sprintf("%s%x%x%x%x", devieType, id_size, id, pwd_size, password)
}

func get_size(str string) int {
  res, err := strconv.Atoi(fmt.Sprintf("%02d", hex.EncodedLen(len(str))))
  if err != nil {
    log.Fatal(err)
  }

  return res
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
      n_out += "fe" + fmt.Sprintf("%02d", hex.EncodedLen(val_bytes)) + out[2:4]
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
  chksum := fmt.Sprintf("%04d", hex.EncodedLen(sum(hexToTuple(msg))))
  //byte_arr, _ := hex.DecodeString(chksum)
  return chksum
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
