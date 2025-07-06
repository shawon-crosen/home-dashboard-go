import fstyle from "./forecast.module.css"

export default function HourlyWeather({ data }) {
    return (
        <ul class="flex-container">
            {data.Hourly.map((hour, index) => (
                <li class="flex-item" key={index} className={fstyle.hourly}>
                    <div>
                        <span>{hour.Time}</span>
                        <span className={fstyle.temp}>{hour.Temp.Value}{hour.Temp.Unit}</span>
                        <span className={fstyle.humidity}>{hour.Humidity.Value}</span>
                        <span className={fstyle.cloudcover}>{hour.CloudCover.Value}</span>
                        <span className={fstyle.windspeed}>{hour.WindSpeed.Value}</span>
                        <span className={fstyle.winddirection}>{hour.WindDirection.Value}</span>
                        <span className={fstyle.windgusts}>{hour.WindGusts.Value}</span>
                        <span className={fstyle.weathercode}>{hour.WeatherCode.Value}</span>
                    </div>
                </li>
            ))}
        </ul>
    )
}

// function cloudCover({ coverPercentage }) {
    
// }