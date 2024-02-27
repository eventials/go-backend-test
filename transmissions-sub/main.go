package transmissions_sub

import (
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
	"time"
	"transmissions-sub/database"
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

func alterNextTransmissionStatus() {
	var transmission *Transmission

	database.DB.Where("date <= ? AND status = ? AND date + make_interval(0, 0, 0, 0, 0, duration * 1000000000)", time.Now(), Upcoming).First(&transmission)

}

func main() {
	database.Connect()

	s := gocron.NewScheduler(time.UTC)

	// 4
	s.Every(1).Minutes().Do(alterNextTransmissionStatus)

	// 5
	s.StartBlocking()
}
