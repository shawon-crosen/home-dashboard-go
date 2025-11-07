import fstyle from "./forecast.module.css"

export default function DailyWeather({ data }) {
    return (
        // <div class="flex-container">
        data.Daily.map((day, index) => (
            <li class="flex-item-d" key={index}>
                <div className={fstyle.time}>
                    <span className={fstyle.time}>
                        {day.Date}
                    </span>
                </div>
                <div class= {fstyle.child}>
                    <span className={fstyle.temp}>
                        {day.TempMean.Value}{day.TempMean.Unit}
                    </span>
                </div>
                <div className={fstyle.child}>
                    <span className={fstyle.weatherstatus}>
                        {getWeatherData(day.WeatherCode)[0]}
                    </span>
                </div>
                <div className={fstyle.child}>
                    <span>
                        <img
                            src={`http://openweathermap.org/img/wn/${getWeatherData(day.WeatherCode)[1]}d@2x.png`}
                        />
                    </span>
                </div>
                {/* <div className={fstyle.weatherdetails}>
                    <span className={fstyle.weatherdetailschild}>Feels Like</span>
                    <span className={fstyle.weatherdetailschild}>{day.ApparentTemp.Value}{day.ApparentTemp.Unit}</span>
                </div> */}
                <div className={fstyle.weatherdetails}>
                    <span className={fstyle.weatherdetailschild}>Wind Speed</span>
                    <span className={fstyle.weatherdetailschild}>{day.WindSpeedMax.Value}{day.WindSpeedMax.Unit}</span>
                </div>
                <div className={fstyle.weatherdetails}>
                    <span className={fstyle.weatherdetailschild}>Wind Gusts</span>
                    <span className={fstyle.weatherdetailschild}> {day.WindGustsMax.Value}{day.WindGustsMax.Unit}</span>
                </div>
                <div className={fstyle.weatherdetails}>
                    <span className={fstyle.weatherdetailschild}>Precipitation</span>
                    <span className={fstyle.weatherdetailschild}>{day.PrecipProbability.Value}{day.PrecipProbability.Unit}</span>
                </div>
                <div className={fstyle.weathercode}>{day.WeatherCode.Value}</div>
            </li>
        ))
        // </div>
    )
}

function isDay(dayCode) {
    switch(dayCode) {
        case 0:
            return "n"
        case 1:
            return "d"  
    }
}

function getWeatherData(weatherCode) {
    switch(weatherCode) {
        case 0:
            return ["Clear Sky", "01"]
        case 1:
        case 2:
            return ["Partly Cloudy", "02"]
        case 3:
            return ["Overcast", "03"]
        case 45:
        case 48:
            return ["Fog", "50"]
        case 51:
            return ["Light Drizzle", "09"]
        case 53:
            return ["Moderate Drizzle", "09"]
        case 55:
            return ["Heavy Drizzle", "09"]
        case 56:
            return ["Light Freezing Drizzle", "09"]
        case 57:
            return ["Heavy Freezing Drizzle", "09"]
        case 61:
            return ["Light Rain", "09"]
        case 63:
            return ["Moderate Rain", "09"]
        case 65:
            return ["Heavy Rain", "09"]
        case 66:
            return ["Light Freezing Rain", "10"]
        case 67:
            return ["Heavy Freezing Rain", "10"]
        case 71:
        case 77:
            return ["Light Snow Fall", "13"]
        case 73:
            return ["Moderate Snow Fall", "13"]
        case 75:
            return ["Heavy Snow Fall", "13"]
        case 80:
            return ["Light Rain Showers", "10"]
        case 81:
            return ["Moderate Rain Showers", "10"]
        case 82:
            return ["Heavy Rain Showers", "10"]
        case 85:
            return ["Light Snow Showers", "13"]
        case 86:
            return ["Heavy Snow Showers", "13"]
        case 95:
        case 96:
        case 99:
            return ["Thunderstorms", "11"]
    }
}