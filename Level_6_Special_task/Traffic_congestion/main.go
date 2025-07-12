package main

type WeatherCondition int

const (
	Clear WeatherCondition = iota + 1
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	cof := float64(Clear)

	switch weather.Condition {
	case HeavyRain:
		cof *= 0.2
	case Snow:
		cof *= 0.15
	case Rain:
		cof *= 0.125
	}

	if weather.WindSpeed > 15 {
		cof *= 0.1
	}

	return cof
}
