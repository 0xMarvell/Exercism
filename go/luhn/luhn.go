package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	//panic("Please implement the Valid function")
	id = strings.ReplaceAll(id, " ", "")

	if len(id) < 2 {
		return false
	}

	isSecond := len(id)%2 == 0
	sum := 0
	for _, ch := range id {

		if !unicode.IsDigit(ch) {
			return false
		}

		n := int(ch - '0')
		// double every second digit
		if isSecond {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		isSecond = !isSecond
	}

	return sum%10 == 0
}
