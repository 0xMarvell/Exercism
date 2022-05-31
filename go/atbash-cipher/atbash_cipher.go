package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	cipher := "zyxwvutsrqponmlkjihgfedcba"
	var temp string
	var result string
	s = strings.ToLower(s)
	for _, v := range s {
		idx := strings.Index(alpha, string(v))
		if idx != -1 {
			val := cipher[idx]
			temp += string(val)
		} else if !unicode.IsPunct(v) && !unicode.IsSpace(v) {
			temp += string(v)
		}
	}
	for i, v := range temp {
		result += string(v)
		if (i+1)%5 == 0 && (i+1) != len(temp) {
			result += " "
		}
	}
	return result
}
