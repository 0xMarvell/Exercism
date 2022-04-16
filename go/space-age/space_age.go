package space

import "strings"

type Planet string

const earthInSeconds float64 = 31557600

func Age(seconds float64, planet Planet) float64 {
	//panic("Please implement the Age function")
	planetLower := strings.ToLower(string(planet))
	var age float64

	switch {
	case planetLower == "Mercury":
		mercuryInSeconds := earthInSeconds * 0.2408467
		age = seconds / mercuryInSeconds
		return age
	case planetLower == "Venus":
		venusInSeconds := earthInSeconds * 0.61519726
		age = seconds / venusInSeconds
		return age
	case planetLower == "Mars":
		marsInSeconds := earthInSeconds * 1.8808158
		age = seconds / marsInSeconds
		return age
	case planetLower == "Jupiter":
		jupiterInSeconds := earthInSeconds * 11.862615
		age = seconds / jupiterInSeconds
		return age
	case planetLower == "Saturn":
		saturnInSeconds := earthInSeconds * 29.447498
		age = seconds / saturnInSeconds
		return age
	case planetLower == "Uranus":
		uranusInSeconds := earthInSeconds * 84.016846
		age = seconds / uranusInSeconds
		return age
	case planetLower == "Neptune":
		mercuryInSeconds := earthInSeconds * 164.79132
		age = seconds / mercuryInSeconds
		return age
	default:
		age = seconds / earthInSeconds
		return age
	}
}
