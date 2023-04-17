package fan

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

const PROTOCOL = "udp"

type Fan struct {
	// Fan IP Address
	IPAddress string
	// Fan unique ID
	ID string
	// Port the fan UDP server runs on - default is 4000
	Port int
	// Fan password
	Password string
	// Name of the fan
	Name string
	// IsWorking sets the state of the fan
	IsWorking bool
	// Maxed Timeout is the time in seconds the fans will run at max when triggered
	MaxedTimeout int

	conn *net.UDPConn
}

// Create a new Fan
func NewFan(ip_addr, id, pwd, name string, port, maxedTimeout int) *Fan {
	fan := &Fan{
		IPAddress:    ip_addr,
		ID:           id,
		Port:         port,
		Password:     pwd,
		Name:         name,
		MaxedTimeout: maxedTimeout,
	}
	fan.Connect()

	return fan
}

// Change fan speed. Possible param values (1..5)
func (f *Fan) ChangeFanSpeed(speed string) string {
	data := f.encodedata(OP_WRITE_RETURN_REQUEST, OP_SPEED_MODE_REQUEST, speed)
	response := f.send(data)

	return response
}

// Change fan operation mode
// Possible params are "in, out, invert"
func (f *Fan) ChangeFanOperation(operation string) string {
	op := ""
	switch operation {
	case "in":
		op = OP_AIR_IN
	case "out":
		op = OP_AIR_OUT
	case "invert":
		op = OP_AIR_INVERT
	default:
		log.Println("Fan operation not recognised. Defaulting to invert operation.")
		op = OP_AIR_INVERT
	}

	data := f.encodedata(OP_WRITE_RETURN_REQUEST, OP_AIRFLOW_REQUEST, op)
	response := f.send(data)

	return response
}

// Connect to fan
func (f *Fan) Connect() {
	server, err := net.ResolveUDPAddr(PROTOCOL, fmt.Sprintf("%s:%d", f.IPAddress, f.Port))
	if err != nil {
		log.Fatalf("Could not Connect to fan (%s) udp server: %v", f.Name, err)
	}

	conn, err := net.DialUDP(PROTOCOL, nil, server)
	if err != nil {
		log.Fatalf("Could not Connect to fan (%s): %v", f.Name, err)
	}

	f.conn = conn
}

func (f *Fan) PollMaxedTimeout() {
	timeInFuture := time.Now().Add(time.Duration(f.MaxedTimeout) * time.Second).Unix()
	go func() {
		// Set the polling period to MaxedTimeout
		ticker := time.NewTicker(time.Duration(f.MaxedTimeout) * time.Second)
		for {
			select {
			case <-ticker.C:
				log.Println("MaxedTimeout Polling...")

				if time.Now().Unix() >= timeInFuture {
					log.Printf("MaxedTimeout complete for fan %s. Reducing speed.", f.Name)
					f.ChangeFanSpeed(LOW_SPEED)
					f.IsWorking = false
				}
			}
		}
	}()
}

// Send and receive data from fan
func (f *Fan) send(data string) string {
	payload := f.buildRequestHeaders() + data
	payload = HEADER + payload + checksum(payload)
	byteData, err := hex.DecodeString(payload)

	if err != nil {
		log.Fatalf("Could not decode payload data: %v", err)
	}

	_, err = f.conn.Write(byteData)
	if err != nil {
		log.Fatalf("Could not write to fan (%s): %v", f.Name, err)
	}

	bufferedData := make([]byte, 64)
	_, err = f.conn.Read(bufferedData)
	fmt.Printf("Buffered Data: %v\n", bufferedData)
	fmt.Printf("Buffered Data string: %v\n", string(bufferedData))

	if err != nil {
		log.Fatalf("Could not read response data from fan (%s): %v", f.Name, err)
	}

	responseStr := string(bufferedData[25:])
	return responseStr
}

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

func (f *Fan) encodedata(operation, param, value string) string {
	out := ""
	parameter := ""
	val_bytes := 0

	for i := 0; i < len(param); i += 4 {
		n_out := ""
		out = param[i:(i + 4)]
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
	byte_array, err := hex.DecodeString(chksum)
	if err != nil {
		log.Fatalf("Could not decode checksum for fan: %v", err)
	}
	ck := fmt.Sprintf("%02x", byte_array[1]) + fmt.Sprintf("%02x", byte_array[0])
	return ck
}

func hexToTuple(msg string) []int64 {
	result := []int64{}
	val := int64(0)
	err := errors.New("")

	for i := 0; i < len(msg); i += 2 {
		if (i + 2) > len(msg) {
			val = 1
		} else {
			val, err = strconv.ParseInt(msg[i:(i+2)], 16, 16)
			if err != nil {
				log.Fatalf("Could no convert hex to tuple: %v", err)
			}
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
