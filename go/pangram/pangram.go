package pangram

import "strings"

func IsPangram(input string) bool {
	//panic("Please implement the IsPangram function")
	alphabets := map[rune]bool{}
	for _, letter := range strings.ToLower(input) {
		if letter >= 'a' && letter <= 'z' {
			alphabets[letter] = true
		}
	}
	return len(alphabets) == 26
}
