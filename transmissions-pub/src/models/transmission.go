package models

import (
	"gorm.io/gorm"
	"time"
)

type TransmissionStatus uint8

const (
	Upcoming TransmissionStatus = iota
	Live
	Ended
)

type Transmission struct {
	gorm.Model

	Title    string        `json:"title" binding:"required"`
	Date     time.Time     `json:"date" binding:"required"`
	Duration time.Duration `json:"duration" binding:"required"`
	Status   TransmissionStatus
}
