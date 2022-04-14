package pangram

import "strings"

func IsPangram(input string) bool {
	alphabets := map[rune]bool{}
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			alphabets[letter] = true
		}
	}
	return len(alphabets) == 26
}
