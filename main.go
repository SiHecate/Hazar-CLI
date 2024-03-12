package main

import (
	"fmt"
	"hazar/model"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tarm/serial"
	serialRead "go.bug.st/serial"
)

var SerialPort *serial.Port

func main() {
	ComPorts()
	/*
		as usual /dev/ttyACM0
		chown /dev/ttyACM0 for ubuntu
		sudo su
		type your password
		cd /
		cd dev
		chown username ttyUSB0
	*/
	ComRead()

}

func ComRead() {
	if err := ConnectComPort("/dev/ttyACM0"); err != nil {
		log.Fatal(err)
	}
	defer SerialPort.Close()

	for {
		buf := make([]byte, 256)
		n, err := SerialPort.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		data := string(buf[:n])
		SaveToFile(data)

		var datas model.SensorData

		fields := strings.Fields(string(buf))
		for i := 0; i < len(fields); i++ {
			switch fields[i] {

			case "Time:":
				datas.Time = time.Now()

			case "Encoder1:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println("deneme")
				}
				datas.Encoder1 = value

			case "Encoder2:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Encoder2 = value

			case "Encoder3:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Encoder3 = value

			case "Encoder4:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Encoder4 = value

			case "Ax:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Ax = value

			case "Ay:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Ay = value

			case "Az:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Az = value

			case "Rx:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Rx = value

			case "Ry:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Ry = value

			case "Rz:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Rz = value

			case "Altitude:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Altitude = value

			case "Temp:":
				value, err := strconv.Atoi(fields[i+1])
				if err != nil {
					fmt.Println()
				}
				datas.Temp = value
			}
			// fmt.Println(datas)
		}
	}

}

func FileCheck(fileName string) error {
	_, err := os.Stat(fileName)
	if err == nil {
		return err
	}
	if err != nil {
		os.Create("data.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
	return err
}

func SaveToDatabase(data string) {

}

func SaveToFile(data string) error {
	fileName := "data.txt"
	err := FileCheck(fileName)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(data + "\n"); err != nil {
		return err
	}
	return nil
}

func ConnectComPort(address string) error {
	c := &serial.Config{
		Name: address,
		Baud: 115200,
	}
	s, err := serial.OpenPort(c)
	if err != nil {
		return err
	}
	SerialPort = s
	fmt.Println("Connection success")
	return nil
}

func ComPorts() {
	ports, err := serialRead.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found avaliable")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}
