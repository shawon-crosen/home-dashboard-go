import fstyle from "./forecast.module.css"

export default function HourlyWeather({ data }) {
    return (
        <ul class="flex-container">
            {data.Hourly.map((hour, index) => (
                <li class="flex-item" key={index} className={fstyle.hourly}>
                    <div>
                        <span className={fstyle.temp}>{hour.Temp}</span>
                        <span className={fstyle.humidity}>{hour.Humidity}</span>
                        <span className={fstyle.cloudcover}>{hour.CloudCover}</span>
                        <span className={fstyle.windspeed}>{hour.WindSpeed}</span>
                        <span className={fstyle.winddirection}>{hour.WindDirection}</span>
                        <span className={fstyle.windgusts}>{hour.WindGusts}</span>
                        <span className={fstyle.weathercode}>{hour.WeatherCode}</span>
                    </div>
                </li>
            ))}
        </ul>
    )
}