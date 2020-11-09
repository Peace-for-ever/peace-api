package models

import (
	"github.com/brianvoe/gofakeit"
)

//GenerateSentence from anything
func GenerateSentence() (s string) {
	return gofakeit.Sentence(10)
}
