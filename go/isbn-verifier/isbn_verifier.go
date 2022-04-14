// Pacakge isbn checks if the inputted string is a valid ISBN number
package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN checks if the inputted string is a valid ISBN-1o
func IsValidISBN(isbn string) bool {
	newIsbn := strings.ReplaceAll(isbn, "-", "")
	result := 0

	if len(newIsbn) != 10 {
		return false
	}

	for i, v := range newIsbn {
		if i < 9 && unicode.IsLetter(v) {
			return false
		}
		digit := int(v - '0')
		if v == 'X' {
			digit = 10
		}

		result += digit * (10 - i)
	}

	return result%11 == 0
}
