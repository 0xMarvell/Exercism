package raindrops

import (
	"strconv"
)

// Convert Function converts an integer to either a "Pling", "Plang" or "Plong" sound
// depicting the sound of rainfrops
func Convert(number int) string {
	var drops string

	if number%3 == 0 {
		drops += "Pling"
	}
	if number%5 == 0 {
		drops += "Plang"
	}
	if number%7 == 0 {
		drops += "Plong"
	}
	if drops == "" {
		drops = strconv.Itoa(number)
	}
	return drops

}
