package models

import "github.com/tjarratt/babble"

//GenerateSentence from anything
func GenerateSentence() (s string) {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	return babbler.Babble()
}
