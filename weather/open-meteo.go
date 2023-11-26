package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const url = "https://api.open-meteo.com/v1/forecast"

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
	pastDays      int64
	forecastDays  int64
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
	WindDirection       []string  `json:"wind_direction_10m"`
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
	Temperature_Max              []float64 `json:"temperature_2m_max"`
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
	Temperature_Max              string `json:"temperature_2m_max"`
	TemperatureMin               string `json:"temperature_2m_min"`
	PrecipitationTotal           string `json:"precipitation_sum"`
	PrecipitationProbabilityMean string `json:"precipitation_probability_mean"`
	WeatherCode                  string `json:"weather_code"`
	Sunrise                      string `json:"sunrise"`
	Sunset                       string `json:"sunset"`
	WindSpeedMax                 string `json:"wind_speed_10m_max"`
	WindGustsMax                 string `json:"wind_gusts_10m_max"`
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

func (w Weather) GetData() ForecastResponse {
	target := ForecastResponse{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
	}

	q := req.URL.Query()
	q.Add("latitude", fmt.Sprintf("%v", w.Params.latitude))
	q.Add("longitude", fmt.Sprintf("%v", w.Params.longitude))
	q.Add("elevation", fmt.Sprintf("%v", w.Params.elevation))
	q.Add("hourly", strings.Join(w.Params.hourly, ","))
	q.Add("daily", strings.Join(w.Params.daily, ","))
	q.Add("current", strings.Join(w.Params.current, ","))
	q.Add("temperature_unit", w.Params.tempUnit)
	q.Add("wind_speed_unit", w.Params.windSpeedUnit)
	q.Add("precipitation_unit", w.Params.precipUnit)
	q.Add("timeformat", w.Params.timeFormat)
	q.Add("timezone", w.Params.timeZone)
	q.Add("past_days", fmt.Sprintf("%v", w.Params.pastDays))
	q.Add("forecast_days", fmt.Sprintf("%v", w.Params.forecastDays))
	q.Add("start_date", w.Params.startDate)
	q.Add("end_date", w.Params.endDate)

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	resp, err := w.Client.Get(req.URL.String())

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&target)

	if err != nil {
		log.Println(err)
	}

	return target
}
