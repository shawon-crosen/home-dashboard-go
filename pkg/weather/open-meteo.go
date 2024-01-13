package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/shawon-crosen/dashboard-go/pkg/config"
)

const url = "https://api.open-meteo.com/v1/forecast"

func NewForecastParams(conf config.WeatherConfig) ForecastParams {
	fp := ForecastParams{}

	fp.latitude = conf.Latitude
	fp.longitude = conf.Longitude
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
	fp.tempUnit = conf.TempUnit
	fp.windSpeedUnit = conf.WindSpeedUnit
	fp.precipUnit = conf.PrecipUnit
	fp.timeFormat = "iso8601"
	fp.timeZone = conf.TimeZone
	start := time.Now().Format(time.RFC3339)
	fp.startHour = start[:len(start)-9]
	end := time.Now().Local().Add(7 * 24 * time.Hour).Format(time.RFC3339)
	fp.endHour = end[:len(end)-9]
	fp.startDate = time.Now().Format("2006-01-02")
	fp.endDate = time.Now().Local().Add(7 * 24 * time.Hour).Format("2006-01-02")

	return fp
}

func windDirection(w float64) string {
	var dir string

	switch {
	case (w == 0):
		dir = "N"
	case (w > 0 && w < 90):
		dir = "NNE"
	case (w == 90):
		dir = "E"
	case (w > 90 && w < 180):
		dir = "SSE"
	case (w == 180):
		dir = "S"
	case (w > 180 && w < 270):
		dir = "SSW"
	case (w == 270):
		dir = "W"
	case (w > 270):
		dir = "NNW"
	}

	return dir
}

func (w Weather) GetData(fType string) *ForecastResponse {
	target := ForecastResponse{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
	}

	var params string
	switch {
	case fType == "hourly":
		params = strings.Join(w.Params.hourly, ",")
	case fType == "daily":
		params = strings.Join(w.Params.daily, ",")
	case fType == "current":
		params = strings.Join(w.Params.current, ",")
	}

	q := req.URL.Query()
	q.Add("latitude", fmt.Sprintf("%v", w.Params.latitude))
	q.Add("longitude", fmt.Sprintf("%v", w.Params.longitude))
	q.Add("elevation", fmt.Sprintf("%v", w.Params.elevation))
	q.Add(fType, params)
	q.Add("temperature_unit", w.Params.tempUnit)
	q.Add("wind_speed_unit", w.Params.windSpeedUnit)
	q.Add("precipitation_unit", w.Params.precipUnit)
	q.Add("timeformat", w.Params.timeFormat)
	q.Add("timezone", w.Params.timeZone)
	switch {
	case fType == "hourly":
		q.Add("start_hour", w.Params.startHour)
		q.Add("end_hour", w.Params.endHour)
	case fType == "daily":
		q.Add("start_date", w.Params.startDate)
		q.Add("end_date", w.Params.endDate)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := w.Client.Get(req.URL.String())

	if err != nil {
		log.Println(err)
		return nil
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&target)

	if err != nil {
		log.Println(err)
	}

	return &target
}

func (w Weather) FormatData(fr ForecastResponse, fType string) Forecast {
	f := Forecast{}

	switch {
	case fType == "hourly":
		for i := 0; i < 11; i++ {
			h := HourlyForecast{}
			h.Temp = fmt.Sprintf("%v%v", fr.HourlyData.Temperature[i], fr.HourlyUnits.Temperature)
			h.ApparentTemp = fmt.Sprintf("%v%v", fr.HourlyData.ApparentTemperature[i], fr.HourlyUnits.ApparentTemperature)
			h.CloudCover = fmt.Sprintf("%v%v", fr.HourlyData.CloudCover[i], fr.HourlyUnits.CloudCover)
			h.Humidity = fmt.Sprintf("%v%v", fr.HourlyData.RelativeHumidity[i], fr.HourlyUnits.RelativeHumidity)
			h.Precip = fmt.Sprintf("%v %ves", fr.HourlyData.Precipitation[i], fr.HourlyUnits.Precipitation)
			h.WeatherCode = fmt.Sprintf("%v", fr.HourlyData.WeatherCode[i])
			h.WindDirection = fmt.Sprintf("%v", windDirection(fr.HourlyData.WindDirection[i]))
			h.WindGusts = fmt.Sprintf("%v %v", fr.HourlyData.WindGusts[i], fr.HourlyUnits.WindGusts)
			h.WindSpeed = fmt.Sprintf("%v %v", fr.HourlyData.WindSpeed[i], fr.HourlyUnits.WindSpeed)

			f.Hourly = append(f.Hourly, h)
		}
	case fType == "daily":
		for j := 0; j < 7; j++ {
			d := DailyForecast{}

			d.TempMax = fmt.Sprintf("%v%v", fr.DailyData.TemperatureMax[j], fr.DailyUnits.TemperatureMax)
			d.TempMin = fmt.Sprintf("%v%v", fr.DailyData.TemperatureMin[j], fr.DailyUnits.TemperatureMin)
			d.PrecipProbability = fmt.Sprintf("%v%v", fr.DailyData.PrecipitationProbabilityMean[j], fr.DailyUnits.PrecipitationProbabilityMean)
			d.PrecipTotal = fmt.Sprintf("%v %ves", fr.DailyData.PrecipitationTotal[j], fr.DailyUnits.PrecipitationTotal)
			d.Sunrise = fmt.Sprintf("%v", fr.DailyData.Sunrise[j])
			d.Sunset = fmt.Sprintf("%v", fr.DailyData.Sunset[j])
			d.WeatherCode = fmt.Sprintf("%v", fr.DailyData.WeatherCode[j])
			d.WindGustsMax = fmt.Sprintf("%v %v", fr.DailyData.WindGustsMax[j], fr.DailyUnits.WindGustsMax)
			d.WindSpeedMax = fmt.Sprintf("%v %v", fr.DailyData.WindSpeedMax[j], fr.DailyUnits.WindSpeedMax)

			f.Daily = append(f.Daily, d)
		}
	case fType == "current":
		c := CurrentForecast{}
		c.Temp = fmt.Sprintf("%v%v", fr.CurrentData.Temperature, fr.CurrentUnits.Temperature)
		c.ApparentTemp = fmt.Sprintf("%v%v", fr.CurrentData.ApparentTemperature, fr.CurrentUnits.ApparentTemperature)
		c.CloudCover = fmt.Sprintf("%v%v", fr.CurrentData.CloudCover, fr.CurrentUnits.CloudCover)
		c.Humidity = fmt.Sprintf("%v%v", fr.CurrentData.RelativeHumidity, fr.CurrentUnits.RelativeHumidity)
		c.Precip = fmt.Sprintf("%v %ves", fr.CurrentData.Precipitation, fr.CurrentUnits.Precipitation)
		c.WeatherCode = fmt.Sprintf("%v", fr.CurrentData.WeatherCode)
		c.WindDirection = fmt.Sprintf("%v", windDirection(fr.CurrentData.WindDirection))
		c.WindGusts = fmt.Sprintf("%v %v", fr.CurrentData.WindGusts, fr.CurrentUnits.WindGusts)
		c.WindSpeed = fmt.Sprintf("%v %v", fr.CurrentData.WindSpeed, fr.CurrentUnits.WindSpeed)

		f.Current = c
	}

	return f
}
