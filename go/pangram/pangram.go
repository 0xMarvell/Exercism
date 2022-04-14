// Package panagram checks if given string is a panagram
package pangram

import "strings"

// IsPanagram checks if the given string is a panagram
func IsPangram(input string) bool {
	alphabets := map[rune]bool{}
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			alphabets[letter] = true
		}
	}
	return len(alphabets) == 26
}
