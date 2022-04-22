package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	//panic("Please implement the Transform function")
	scrabble := make(map[string]int)

	for i, v := range in {
		for _, letter := range v {
			scrabble[strings.ToLower(letter)] = i
		}
	}

	return scrabble
}
