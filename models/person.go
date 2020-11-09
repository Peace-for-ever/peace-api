package models

import (
	"github.com/brianvoe/gofakeit"
)

//Person struct
type Person struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Location  Location `json:"location"`
}

//RandomPerson generator with max location param
func RandomPerson(maxX, maxY int) (p Person) {
	p.Location.RandomLocation(10, 10)

	p.FirstName = gofakeit.Person().FirstName
	p.LastName = gofakeit.Person().LastName

	return p
}
