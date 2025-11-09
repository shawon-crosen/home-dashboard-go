package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/shawon-crosen/dashboard-go/server/pkg/config"
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
		"precipitation_probability",
		"weather_code",
		"is_day",
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
		"temperature_2m_mean",
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
		"is_day",
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
	fmt.Println("QUERY: ", req.URL.RawQuery)

	resp, err := w.Client.Get(req.URL.String())

	if err != nil {
		fmt.Println(err)
	}

	known_errors := []int{404, 400}

	if slices.Contains(known_errors, resp.StatusCode) {
		fmt.Println(q)
		fmt.Println(resp)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&target)

	if err != nil {
		log.Println(err)
	}

	return &target
}

func (w Weather) FormatHourlyData(fr ForecastResponse, fType string) HourlyPayload {
	f := HourlyPayload{}

	for i := 2; i < 10; i += 2 {
		h := HourlyForecast{}
		h.Temp.Value = fr.HourlyData.Temperature[i]
		h.Temp.Unit = fr.HourlyUnits.Temperature

		h.ApparentTemp.Value = fr.HourlyData.ApparentTemperature[i]
		h.ApparentTemp.Unit = fr.HourlyUnits.ApparentTemperature

		h.CloudCover.Value = fr.HourlyData.CloudCover[i]
		h.CloudCover.Unit = fr.HourlyUnits.CloudCover

		h.Humidity.Value = fr.HourlyData.RelativeHumidity[i]
		h.Humidity.Unit = fr.HourlyUnits.RelativeHumidity

		h.Precip.Value = fr.HourlyData.Precipitation[i]
		h.Precip.Unit = fr.HourlyUnits.Precipitation

		h.WeatherCode = fr.HourlyData.WeatherCode[i]

		h.WindDirection = windDirection(fr.HourlyData.WindDirection[i])

		h.WindGusts.Value = fr.HourlyData.WindGusts[i]
		h.WindGusts.Unit = fr.HourlyUnits.WindGusts

		h.WindSpeed.Value = fr.HourlyData.WindSpeed[i]
		h.WindSpeed.Unit = fr.HourlyUnits.WindSpeed

		hTime, err := time.Parse("2006-01-02T15:04", fr.HourlyData.Time[i])
		if err != nil {
			fmt.Println(err)
		}
		h.Time = hTime.Format(time.Kitchen)
		h.IsDay = fr.HourlyData.IsDay[i]

		f.Hourly = append(f.Hourly, h)
	}

	return f
}

func (w Weather) FormatDailyData(fr ForecastResponse, fType string) DailyPayload {
	f := DailyPayload{}
	for j := 1; j < 5; j++ {
		d := DailyForecast{}

		d.TempMean.Value = fr.DailyData.TemperatureMean[j]
		d.TempMean.Unit = fr.DailyUnits.TemperatureMean

		d.TempMax.Value = fr.DailyData.TemperatureMax[j]
		d.TempMax.Unit = fr.DailyUnits.TemperatureMax

		d.TempMin.Value = fr.DailyData.TemperatureMin[j]
		d.TempMin.Unit = fr.DailyUnits.TemperatureMin

		d.PrecipProbability.Value = fr.DailyData.PrecipitationProbabilityMean[j]
		d.PrecipProbability.Unit = fr.DailyUnits.PrecipitationProbabilityMean

		d.PrecipTotal.Value = fr.DailyData.PrecipitationTotal[j]
		d.PrecipTotal.Unit = fr.DailyUnits.PrecipitationTotal

		d.Sunrise, _ = time.Parse(fr.DailyData.Sunrise[j], time.RFC3339)

		d.Sunset, _ = time.Parse(fr.DailyData.Sunset[j], time.RFC3339)

		d.WeatherCode = fr.DailyData.WeatherCode[j]

		d.WindGustsMax.Value = fr.DailyData.WindGustsMax[j]
		d.WindGustsMax.Unit = fr.DailyUnits.WindGustsMax

		d.WindSpeedMax.Value = fr.DailyData.WindSpeedMax[j]
		d.WindSpeedMax.Unit = fr.DailyUnits.WindSpeedMax

		date, _ := time.Parse("2006-01-02T15:04", fr.DailyData.Sunrise[j])
		d.Date = date.Weekday().String()

		f.Daily = append(f.Daily, d)
	}

	return f
}

func (w Weather) FormatCurrentData(fr ForecastResponse, fType string) CurrentPayload {
	f := CurrentPayload{}

	f.Current.Temp.Value = fr.CurrentData.Temperature
	f.Current.Temp.Unit = fr.CurrentUnits.Temperature

	f.Current.ApparentTemp.Value = fr.CurrentData.ApparentTemperature
	f.Current.ApparentTemp.Unit = fr.CurrentUnits.ApparentTemperature

	f.Current.CloudCover.Value = fr.CurrentData.CloudCover
	f.Current.CloudCover.Unit = fr.CurrentUnits.CloudCover

	f.Current.Humidity.Value = fr.CurrentData.RelativeHumidity
	f.Current.Humidity.Unit = fr.CurrentUnits.RelativeHumidity

	f.Current.Precip.Value = fr.CurrentData.Precipitation
	f.Current.Precip.Unit = fr.CurrentUnits.Precipitation

	f.Current.WeatherCode = fr.CurrentData.WeatherCode

	f.Current.WindDirection = windDirection(fr.CurrentData.WindDirection)

	f.Current.WindGusts.Value = fr.CurrentData.WindGusts
	f.Current.WindGusts.Unit = fr.CurrentUnits.WindGusts

	f.Current.WindSpeed.Value = fr.CurrentData.WindSpeed
	f.Current.WindSpeed.Unit = fr.CurrentUnits.WindSpeed

	f.Current.IsDay = fr.CurrentData.IsDay

	return f
}
