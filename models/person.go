package models

import (
	"github.com/Pallinder/go-randomdata"
)

type Person struct {
	FirstName string
	LastName  string
	Location
}

func RandomPerson(maxX, maxY int) (p Person) {
	p.Location.RandomLocation(10, 10)

	p.FirstName = randomdata.FirstName(randomdata.RandomGender)
	p.LastName = randomdata.LastName()

	return p
}
