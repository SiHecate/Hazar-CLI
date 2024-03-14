package model

import "time"

type SensorData struct {
	ID       uint `gorm:"primaryKey"`
	Time     time.Time
	TestNo   uint
	Encoder1 int
	Encoder2 int
	Encoder3 int
	Encoder4 int
	Ax       float64
	Ay       float64
	Az       float64
	Rx       float64
	Ry       float64
	Rz       float64
	Altitude float64
	Temp     float64
}
