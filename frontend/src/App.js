import './App.css';
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from '@tanstack/react-query';
import HourlyWeather from "./HourlyWeatherComponent/HourlyWeather.js"
import CurrentWeather from "./CurrentWeatherComponent/CurrentWeather.js"
import DailyWeather from './DailyWeatherComponent/DailyWeather.js';

const hourlyWeatherClient = new QueryClient()
const dailyWeatherClient = new QueryClient()
const currentWeatherClient = new QueryClient()

function FetchHourlyWeather() {
  const { isPending, error, data } = useQuery({
    queryFn: () =>
      fetch('api/weather/hourly').then(
        (res) => res.json(),
      ),
      refetchInterval: 30000,
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

  return(< DailyWeather data={data} />)
}

function FetchCurrentWeather() {
  const { isPending, error, data } = useQuery({
    queryFn: () =>
      fetch('api/weather/current').then(
        (res) => res.json(),
      ),
      refetchInterval: 30000,
  })

  if (isPending) return 'Loading...'

  if (error) return 'An error has ocurred ' + error.message

  return(
    < CurrentWeather data={data} />
  )
}

function App() {
  return (
    <div class="box">
      <div class="flex-container">
        <QueryClientProvider client={currentWeatherClient}>
          <FetchCurrentWeather />
        </QueryClientProvider>
      </div>
      <div class="flex-row">
        <QueryClientProvider client={hourlyWeatherClient}>
          <FetchHourlyWeather />
        </QueryClientProvider>
        <QueryClientProvider client={dailyWeatherClient}>
          <FetchDailyWeather />
        </QueryClientProvider>
      </div>
      {/* <div class="flex-row">
        
      </div> */}
    </div>
  );
}

export default App;
