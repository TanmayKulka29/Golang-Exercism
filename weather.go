// Package weather provides the tools to get current weather condition of the city.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the city whose weather is to be displayed.
var CurrentLocation string

// Forecast returns the city and its weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
