import './App.css';
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from '@tanstack/react-query';
import HourlyWeather from "./HourlyWeatherComponent/HourlyWeather.js"
import CurrentWeather from "./CurrentWeatherComponent/CurrentWeather.js"
import DailyWeather from './DailyWeatherComponent/DailyWeather.js';
import Quote from './QuoteComponent/Quote.js';

const hourlyWeatherClient = new QueryClient()
const dailyWeatherClient = new QueryClient()
const currentWeatherClient = new QueryClient()
const quoteClient = new QueryClient()

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
      refetchInterval: 30000,
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

function FetchQuotes() {
  const { isPending, error, data } = useQuery({
    queryFn: () =>
      fetch('api/quotes/random').then(
        (res) => res.json(),
      ),
      refetchInterval: 86400000,
  })

  if (isPending) return 'Loading...'

  if (error) return 'An error has ocurred ' + error.message

  return(
    < Quote quote={data} />
  )
}

function App() {
  return (
    <div>
      <div class="parent">
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
        </div>
        <div class="qbox">
          <QueryClientProvider client={quoteClient}>
            <FetchQuotes />
          </QueryClientProvider>
        </div>
      </div>
    </div>
  );
}

export default App;
