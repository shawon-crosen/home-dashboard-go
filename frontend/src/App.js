import './App.css';
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from '@tanstack/react-query';
import HourlyWeather from "./HourlyWeatherComponent/HourlyWeather.js"
import CurrentWeather from "./CurrentWeatherComponent/CurrentWeather.js"

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
      refetchInterval: 300000,
  })

  if (isPending) return 'Loading...'

  if (error) return 'An error has ocurred ' + error.message

  return(< HourlyWeather data={data} />)
  
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
    < CurrentWeather data={data} />
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
      <QueryClientProvider client={currentWeatherClient}>
        <span>Current Weather</span>
        <FetchCurrentWeather />
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
