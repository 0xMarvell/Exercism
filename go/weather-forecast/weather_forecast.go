// Package weather provides tools to return the weather forecast for any location it is provided with.
package weather

// CurrentCondition represents the current weather condition for a specific provided location.
var CurrentCondition string

// CurrentLocation represents the porvided location to read the weather forecast from.
var CurrentLocation string

//Forecast takes two string arguments "city" and "condition" and returns a string value that represents the weather forecast for a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
