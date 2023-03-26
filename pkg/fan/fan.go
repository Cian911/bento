package fan

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strconv"
)

const PROTOCOL = "udp"

type Fan struct {
	IPAddress string
	ID        string
	Port      int
	Password  string

  conn *net.UDPConn
}

// Create a new Fan
func NewFan(ip_addr, id, pwd string, port int) *Fan {
	return &Fan{
    IPAddress: ip_addr,
    ID: id,
    Port: port,
    Password: pwd,
  }
}

// Connect to fan
func (f *Fan) Connect() {
  server, err := net.ResolveUDPAddr(PROTOCOL, fmt.Sprintf("%s:%d", f.IPAddress, f.Port))
  if err != nil {
    log.Fatalf("Could not connect to fan (%s) udp server: %v", f.ID, err)
  }

  conn, err := net.DialUDP(PROTOCOL, nil, server)
  if err != nil {
    log.Fatalf("Could not connect to fan (%s): %v", f.ID, err)
  }

  f.conn = conn
}

// Send data to fan
func (f *Fan) Send() {}

// Receive data from fan
func (f *Fan) Receive() {}

func (f *Fan) buildRequestHeaders() string {
  id_size := getSize(f.ID)
  pwd_size := getSize(f.Password)
  id := fmt.Sprintf("%x", f.ID)
  password := fmt.Sprintf("%x", f.Password)

  return fmt.Sprintf("%s%s%s%s%s", DEVICE_TYPE, id_size, id, pwd_size, password)
}

func getSize(str string) string {
  encoding := fmt.Sprintf("%x", len(str))
  res, err := strconv.Atoi(encoding)
  if err != nil {
    log.Fatalf("Failed to encode request headers: %v", err)
  }

  return fmt.Sprintf("%02d", res)
}

func encodeData(operation, param, value string) string {
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
  return data
}

func checksum(msg string) string {
  chksum := fmt.Sprintf("%04x", sum(hexToTuple(msg)))
  byte_array, _ := hex.DecodeString(chksum)
  ck := fmt.Sprintf("%02x", byte_array[1]) + fmt.Sprintf("%02x", byte_array[0])
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

