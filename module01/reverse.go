package module01

import (
	"strings"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var sb strings.Builder
	// some characters (e.g kanji) take more one byte
	// use rune for versatile solutions
	wordRunes := []rune(word)

	for i := len(wordRunes) - 1; i >= 0; i-- {
		sb.WriteRune(wordRunes[i])
	}

	return sb.String()
}
