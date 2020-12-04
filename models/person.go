package models

import (
	"github.com/brianvoe/gofakeit"
)

//Person struct
type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//RandomPerson generator with max location param
func RandomPerson() (p Person) {
	p.FirstName = gofakeit.Person().FirstName
	p.LastName = gofakeit.Person().LastName

	return p
}
