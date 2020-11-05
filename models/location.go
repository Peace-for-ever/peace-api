package models

import (
	"github.com/Pallinder/go-randomdata"
)

//Location struct
type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

//RandomLocation generate random location (x,y) from maxX,maxY
func (l *Location) RandomLocation(maxX, maxY int) {
	l.X = randomdata.Number(0, maxX)
	l.Y = randomdata.Number(0, maxY)
}
