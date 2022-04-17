package romannumerals

import "errors"

// ToRomanNumeral converts a number to it's roman numeral equivalent
func ToRomanNumeral(input int) (string, error) {
	aToR := []struct {
		regular int
		roman   string
	}{
		{1, "I"}, {4, "IV"}, {5, "V"}, {9, "IX"},
		{10, "X"}, {40, "XL"}, {50, "L"}, {90, "XC"},
		{100, "C"}, {400, "CD"}, {500, "D"}, {900, "CM"},
		{1000, "M"},
	}

	if input <= 0 || input > 3000 {
		return "", errors.New("out of range")
	}

	res := ""
	for i := len(aToR) - 1; i >= 0; i-- {
		for input >= aToR[i].regular {
			input -= aToR[i].regular
			res += aToR[i].roman
		}
	}

	return res, nil
}
