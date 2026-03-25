//Package weather provides tools to give the forcast of the weather
//current condition and location.
package weather

var (
    // CurrentCondition represents the weather currently.
	CurrentCondition string
    // CurrentLocation represents the location of the weather currently.
	CurrentLocation  string
)
//Forecast returns a striing value of the CurrentLocation and  CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
