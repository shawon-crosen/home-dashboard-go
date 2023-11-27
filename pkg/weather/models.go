package weather

import (
	"net/http"
)

type Weather struct {
	Client http.Client
	Params ForecastParams
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
	startDate     string
	endDate       string
}

type ForecastResponse struct {
	Latitude     float64            `json:"latitude"`
	Longitude    float64            `json:"longitude"`
	Elevation    float64            `json:"elevation,omitempty"`
	GenTimeMs    float64            `json:"generationtime_ms"`
	UtcOffset    int64              `json:"utc_offset_seconds"`
	Timezone     string             `json:"timezone"`
	TimezoneAbbr string             `json:"timezone_abbreviation"`
	HourlyData   HourlyResponse     `json:"hourly"`
	HourlyUnits  HourUnitsResponse  `json:"hourly_units"`
	DailyData    DailyResponse      `json:"daily"`
	DailyUnits   DailyUnitsResponse `json:"daily_units"`
}

type HourlyResponse struct {
	Time                []string  `json:"time"`
	Temperature         []float64 `json:"temperature_2m"`
	RelativeHumidity    []float64 `json:"relative_humidity_2m"`
	ApparentTemperature []float64 `json:"apparent_temperature"`
	CloudCover          []float64 `json:"cloud_cover"`
	WindSpeed           []float64 `json:"wind_speed_10m"`
	WindDirection       []float64 `json:"wind_direction_10m"`
	WindGusts           []float64 `json:"wind_gusts_10m"`
	Precipitation       []float64 `json:"precipitation"`
	WeatherCode         []float64 `json:"weather_code"`
}

type HourUnitsResponse struct {
	Temperature         string `json:"temperature_2m"`
	RelativeHumidity    string `json:"relative_humidity_2m"`
	ApparentTemperature string `json:"apparent_temperature"`
	CloudCover          string `json:"cloud_cover"`
	WindSpeed           string `json:"wind_speed_10m"`
	WindDirection       string `json:"wind_direction_10m"`
	WindGusts           string `json:"wind_gusts_10m"`
	Precipitation       string `json:"precipitation"`
	WeatherCode         string `json:"weather_code"`
}

type DailyResponse struct {
	TemperatureMax               []float64 `json:"temperature_2m_max"`
	TemperatureMin               []float64 `json:"temperature_2m_min"`
	PrecipitationTotal           []float64 `json:"precipitation_sum"`
	PrecipitationProbabilityMean []float64 `json:"precipitation_probability_mean"`
	WeatherCode                  []float64 `json:"weather_code"`
	Sunrise                      []string  `json:"sunrise"`
	Sunset                       []string  `json:"sunset"`
	WindSpeedMax                 []float64 `json:"wind_speed_10m_max"`
	WindGustsMax                 []float64 `json:"wind_gusts_10m_max"`
}

type DailyUnitsResponse struct {
	TemperatureMax               string `json:"temperature_2m_max"`
	TemperatureMin               string `json:"temperature_2m_min"`
	PrecipitationTotal           string `json:"precipitation_sum"`
	PrecipitationProbabilityMean string `json:"precipitation_probability_mean"`
	WeatherCode                  string `json:"weather_code"`
	Sunrise                      string `json:"sunrise"`
	Sunset                       string `json:"sunset"`
	WindSpeedMax                 string `json:"wind_speed_10m_max"`
	WindGustsMax                 string `json:"wind_gusts_10m_max"`
}

type Forecast struct {
	Hourly []HourlyForecast
	Daily  []DailyForecast
}

type HourlyForecast struct {
	Temp          string
	Humidity      string
	ApparentTemp  string
	CloudCover    string
	WindSpeed     string
	WindDirection string
	WindGusts     string
	Precip        string
	WeatherCode   string
}

type DailyForecast struct {
	TempMax           string
	TempMin           string
	PrecipTotal       string
	PrecipProbability string
	WeatherCode       string
	Sunrise           string
	Sunset            string
	WindSpeedMax      string
	WindGustsMax      string
}
