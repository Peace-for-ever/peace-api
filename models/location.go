package models

import (
	"github.com/brianvoe/gofakeit"
)

//Location struct
type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//RandomLocation generate random location (x,y) from maxX,maxY
func (l *Location)RandomLocation(maxX, maxY int) {
	l.X = gofakeit.Number(0, maxX)
	l.Y = gofakeit.Number(0, maxY)
}
