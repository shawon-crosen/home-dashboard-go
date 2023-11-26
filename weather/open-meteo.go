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
	fp.startDate = time.Now().Format("2006-01-02")
	fp.endDate = time.Now().Local().Add(7 * 24 * time.Hour).Format("2006-01-02")

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
		fmt.Println("response packing error")
		fmt.Println(err)
	}

	return target
}

func (w Weather) FormatData(fr ForecastResponse) Forecast {
	f := Forecast{}

	for i := 0; i < 11; i++ {
		h := HourlyForecast{}
		h.Temp = fmt.Sprintf("%v\u0xB0", fr.HourlyData.Temperature[i])
		h.ApparentTemp = fmt.Sprintf("%v%v", fr.HourlyData.ApparentTemperature[i], fr.HourlyUnits.ApparentTemperature[0])
		h.CloudCover = fmt.Sprintf("%v%v", fr.HourlyData.CloudCover[i], fr.HourlyUnits.CloudCover[0])
		h.Humidity = fmt.Sprintf("%v%v", fr.HourlyData.RelativeHumidity[i], fr.HourlyUnits.RelativeHumidity[0])
		h.Precip = fmt.Sprintf("%v%v", fr.HourlyData.Precipitation[i], fr.HourlyUnits.Precipitation[0])
		h.WeatherCode = fmt.Sprintf("%v%v", fr.HourlyData.WeatherCode[i], fr.HourlyUnits.WeatherCode[0])
		h.WindDirection = fmt.Sprintf("%v%v", fr.HourlyData.WindDirection[i], fr.HourlyUnits.WindDirection[0])
		h.WindGusts = fmt.Sprintf("%v%v", fr.HourlyData.WindGusts[i], fr.HourlyUnits.WindGusts[0])
		h.WindSpeed = fmt.Sprintf("%v%v", fr.HourlyData.WindSpeed[i], fr.HourlyUnits.WindSpeed[0])

		f.Hourly = append(f.Hourly, h)
		fmt.Println(h)
	}

	for j := 0; j < 7; j++ {
		d := DailyForecast{}

		d.TempMax = fmt.Sprintf("%v%v", fr.DailyData.TemperatureMax[j], fr.DailyUnits.TemperatureMax[0])
		d.TempMin = fmt.Sprintf("%v%v", fr.DailyData.TemperatureMin[j], fr.DailyUnits.TemperatureMin[0])
		d.PrecipProbability = fmt.Sprintf("%v%v", fr.DailyData.PrecipitationProbabilityMean[j], fr.DailyUnits.PrecipitationProbabilityMean[0])
		d.PrecipTotal = fmt.Sprintf("%v%v", fr.DailyData.PrecipitationTotal[j], fr.DailyUnits.PrecipitationTotal[0])
		d.Sunrise = fmt.Sprintf("%v%v", fr.DailyData.Sunrise[j], fr.DailyUnits.Sunrise[0])
		d.Sunset = fmt.Sprintf("%v%v", fr.DailyData.Sunset[j], fr.DailyUnits.Sunset[0])
		d.WeatherCode = fmt.Sprintf("%v%v", fr.DailyData.WeatherCode[j], fr.DailyUnits.WeatherCode[0])
		d.WindGustsMax = fmt.Sprintf("%v%v", fr.DailyData.WindGustsMax[j], fr.DailyUnits.WindGustsMax[0])
		d.WindSpeedMax = fmt.Sprintf("%v%v", fr.DailyData.WindSpeedMax[j], fr.DailyUnits.WindSpeedMax[0])

		f.Daily = append(f.Daily, d)
		fmt.Println(d)
	}

	return f
}
