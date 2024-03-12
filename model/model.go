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
	Ax       int
	Ay       int
	Az       int
	Rx       int
	Ry       int
	Rz       int
	Altitude int
	Temp     int
}
