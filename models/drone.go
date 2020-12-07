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
	Country     string    `json:"country"`
}

func GenerateEvent() (d Drone) {
	d.Message = gofakeit.Sentence(10)
	city := gofakeit.Number(0, 10595)
	d.Latitude = Cities[city].Latitude
	d.Longitude = Cities[city].Longitude
	d.Country = Cities[city].Country
	d.Citizen = gofakeit.Person().FirstName + " " + gofakeit.Person().LastName
	d.Date = gofakeit.DateRange(time.Now().AddDate(0, 0, -7), time.Now())
	d.Battery = gofakeit.Number(0, 100)
	d.Temperature = gofakeit.Number(0, 40)

	return d
}
