import './App.css';
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from '@tanstack/react-query';
import HourlyWeather from "./HourlyWeatherComponent/HourlyWeather.js"

const hourlyWeatherClient = new QueryClient()
const dailyWeatherClient = new QueryClient()
const currentWeatherClient = new QueryClient()
const trainClient = new QueryClient()

function FetchHourlyWeather() {
  const { isPending, error, data } = useQuery({
    queryFn: () =>
      fetch('api/weather/hourly').then(
        (res) => res.json(),
      ),
  })

  if (isPending) return 'Loading...'

  if (error) return 'An error has ocurred ' + error.message

  return(< HourlyWeather data={data} />)
  // return(
  //   <ul class="flex-container">
  //       {data.Hourly.map((hour, index) => (
  //         <li class="flex-item" key={index} className={fstyle.hourly}>
  //           <div>
  //             Temp: {hour.Temp}
  //             Humidity: {hour.Humidity}
  //           </div>
  //         </li>
  //     ))}
  //   </ul>
    // <div>
    //   {data.Hourly.map((hour, index) => (
    //     <ul class="flex-container">
    //       <li class="flex-item" key={index} className={fstyle.hourly}>
    //         <div>
    //           Temp: {hour.Temp}
    //         </div>
    //         <div>
    //           Humidity: {hour.Humidity}
    //         </div>
    //       </li>
    //     </ul>
    //   ))}
    // </div>
    //         <p>{hourly.Time}</p>
    //         <p>{hourly.Temperature}</p>
    //         <p>{hourly.RelativeHumidity}</p>
    //         <p>{hourly.ApparentTemperature}</p>
    //         <p>{hourly.CloudCover}</p>
    //         <p>{hourly.WindSpeed}</p>
    //         <p>{hourly.WindDirection}</p>
    //         <p>{hourly.WindGusts}</p>
    //         <p>{hourly.Precipitation}</p>
    //         <p>{hourly.WeatherCode}</p>
}

function FetchDailyWeather() {
  const { isPending, error, data } = useQuery({
    queryFn: () =>
      fetch('api/weather/daily').then(
        (res) => res.json(),
      ),
  })

  if (isPending) return 'Loading...'

  if (error) return 'An error has ocurred ' + error.message

  return(
    <ul>
      {JSON.stringify(data.Daily)}
    </ul>
  )
}

function FetchCurrentWeather() {
  const { isPending, error, data } = useQuery({
    queryFn: () =>
      fetch('api/weather/current').then(
        (res) => res.json(),
      ),
  })

  if (isPending) return 'Loading...'

  if (error) return 'An error has ocurred ' + error.message

  return(
    <ul>
      {JSON.stringify(data.Current)}
    </ul>
  )
}

function FetchTrains() {
  const {isPending, error, data} = useQuery({
    queryFn: () =>
      fetch('api/cta').then(
        (res) => res.json(),
      ),
  })

  if (isPending) return 'Loading...'

  if (error) return 'An error has occurred ' + error.message



  return(
    <ul>
      {JSON.stringify(data.StationResponse)}
    </ul>
  )
}

function App() {
  return (
    <div className="Dashboard">
      <QueryClientProvider client={hourlyWeatherClient}>
        <span>Hourly Weather</span>
        <FetchHourlyWeather/>
      </QueryClientProvider>
      {/* <QueryClientProvider client={dailyWeatherClient}>
        <FetchDailyWeather />
      </QueryClientProvider>
      <QueryClientProvider client={currentWeatherClient}>
        <FetchCurrentWeather />
      </QueryClientProvider>
      <QueryClientProvider client={trainClient}>
        <FetchTrains />
      </QueryClientProvider> */}
    </div>
  );
}

export default App;
