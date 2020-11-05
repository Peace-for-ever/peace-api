package models

import (
	"github.com/Pallinder/go-randomdata"
)

type Person struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Location  Location `json:"location"`
}

func RandomPerson(maxX, maxY int) (p Person) {
	p.Location.RandomLocation(10, 10)

	p.FirstName = randomdata.FirstName(randomdata.RandomGender)
	p.LastName = randomdata.LastName()

	return p
}
