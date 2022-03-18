package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	//panic("Please implement the IsIsogram function")
	lowercase := strings.ToLower(word)
	read_word := make(map[rune]bool)
	for _, v := range lowercase {
		if !unicode.IsLetter(v) {
			continue
		}
		if read_word[v] {
			return false
		}
		read_word[v] = true
	}
	return true
}
