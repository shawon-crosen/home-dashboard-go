package openmeteo

import (
	"net/http"
	"time"
)

const url = "https://api.open-meteo.com/v1/forecast"

type Weather struct {
	client http.Client
	params ForecastParams
}

type ForecastParams struct {
	latitude      float64
	longitude     float64
	elevation     float64
	hourly        []string
	daily         []string
	current       []string
	tempUnit      string
	windSpeedUnit string
	precipUnit    string
	timeFormat    string
	timeZone      string
	pastDays      int64
	forecastDays  int64
	startDate     string
	endDate       string
}

func NewForecastParams() ForecastParams {
	fp := ForecastParams{}

	fp.latitude = 41.88
	fp.longitude = -87.62
	fp.hourly = []string{
		"temperature_2m",
		"relative_humidity_2m",
		"apparent_temperature",
		"cloud_cover",
		"wind_speed_10m",
		"wind_direction_10m",
		"wind_gusts_10m",
		"precipitation",
		"weather_code",
	}
	fp.daily = []string{
		"temperature_2m_max",
		"temperature_2m_min",
		"precipitation_sum",
		"precipitation_probability_mean",
		"weather_code",
		"sunrise",
		"sunset",
		"wind_speed_10m_max",
		"wind_gusts_10m_max",
	}
	fp.current = []string{
		"temperature_2m",
		"relative_humidity_2m",
		"apparent_temperature",
		"cloud_cover",
		"wind_speed_10m",
		"wind_direction_10m",
		"wind_gusts_10m",
		"precipitation",
		"weather_code",
	}
	fp.tempUnit = "fahrenheit"
	fp.windSpeedUnit = "mph"
	fp.precipUnit = "inch"
	fp.timeFormat = "iso8601"
	fp.timeZone = "CST"
	fp.pastDays = 0
	fp.forecastDays = 7
	fp.startDate = time.Now().Format("2023-01-01")
	fp.endDate = time.Now().Local().Add(7 * 24 * time.Hour).GoString()

	return fp
}

func (w Weather) GetData(p ForecastParams) map[string]interface{} {
	data = map[string]interface{}{}

	return data
}
