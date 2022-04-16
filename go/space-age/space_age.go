// Package space calculates how old someone would be on any planet.
package space

type Planet string

const earthInSeconds float64 = 31557600

// Age calculates how old someone would be on any planet, given an age in seconds.
func Age(seconds float64, planet Planet) float64 {
	var planets = map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return seconds / (earthInSeconds * planets[planet])
}
