package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type SensorData struct {
	ID       uint `gorm:"primaryKey"`
	Time     time.Time
	TestNo   uint
	Encoder1 int
	Encoder2 int
	Encoder3 int
	Encoder4 int
	Ax       int
	Ay       int
	Az       int
	Rx       int
	Ry       int
	Rz       int
	Altitude int
	Temp     int
}

func main() {
	// Örnek olarak seri porttan gelen veri
	dataFromSerialPort := "123,456,789,101,102,103,104,105,106,107,108,109,110,111,112"

	// Veriyi ayrıştır
	dataFields := strings.Split(dataFromSerialPort, ",")

	// SensorData struct'ını doldur
	var sensorData SensorData
	if len(dataFields) >= 15 {
		sensorData.Encoder1 = convertToInt(dataFields[0])
		sensorData.Encoder2 = convertToInt(dataFields[1])
		sensorData.Encoder3 = convertToInt(dataFields[2])
		sensorData.Encoder4 = convertToInt(dataFields[3])
		sensorData.Ax = convertToInt(dataFields[4])
		sensorData.Ay = convertToInt(dataFields[5])
		sensorData.Az = convertToInt(dataFields[6])
		sensorData.Rx = convertToInt(dataFields[7])
		sensorData.Ry = convertToInt(dataFields[8])
		sensorData.Rz = convertToInt(dataFields[9])
		sensorData.Altitude = convertToInt(dataFields[10])
		sensorData.Temp = convertToInt(dataFields[11])
	} else {
		log.Fatal("Veri eksik")
	}

	// Struct'ı kontrol et
	fmt.Printf("%+v\n", sensorData)
}

// Veriyi int'e dönüştürmek için yardımcı fonksiyon
func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
