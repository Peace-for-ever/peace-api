package models

import (
	"time"

	"github.com/brianvoe/gofakeit"
)

//Drone struct
type Drone struct {
	Citizen     string    `json:"citizen"`
	Message     string    `json:"message"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Date        time.Time `json:"date"`
	Battery     int       `json:"battery"`
	Temperature int       `json:"temperature"`
}

func GenerateEvent() (d Drone) {
	d.Message = gofakeit.Sentence(10)
	d.Latitude = gofakeit.Latitude()
	d.Longitude = gofakeit.Longitude()
	d.Citizen = gofakeit.Person().FirstName + " " + gofakeit.Person().LastName
	d.Date = gofakeit.DateRange(time.Now().AddDate(0, 0, -1), time.Now())
	d.Battery = gofakeit.Number(0, 100)
	d.Temperature = gofakeit.Number(0, 40)

	return d
}
