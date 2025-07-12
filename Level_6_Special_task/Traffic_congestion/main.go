package main

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	cof := 1.0

	switch weather.Condition {
	case HeavyRain:
		cof += 0.2
	case Snow:
		cof += 0.15
	case Rain:
		cof += 0.125
	}

	if weather.WindSpeed > 15 {
		cof += 0.1
	}

	return cof
}
