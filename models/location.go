package models

import (
	"math/rand"
	"time"
)

type Location struct {
	X int
	Y int
}

func (l *Location) RandomLocation(maxX, maxY int) {
	rand.Seed(time.Now().UnixNano())
	l.X = rand.Intn(maxX-0+1) + 0
	l.Y = rand.Intn(maxY-0+1) + 0
}
