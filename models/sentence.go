package models

import "github.com/tjarratt/babble"

func GenerateSentence() (s string) {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	return babbler.Babble()
}
