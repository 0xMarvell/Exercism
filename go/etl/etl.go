package etl

import "strings"

// Transform transforms data from map into an array
// then returns the data in new format
func Transform(in map[int][]string) map[string]int {
	scrabble := make(map[string]int)

	for i, v := range in {
		for _, letter := range v {
			scrabble[strings.ToLower(letter)] = i
		}
	}

	return scrabble
}
