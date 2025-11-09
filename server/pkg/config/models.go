package config

type Config struct {
	WeatherConfig WeatherConfig `yaml:"weather"`
}

type WeatherConfig struct {
	Latitude      float64 `yaml:"latitude"`
	Longitude     float64 `yaml:"longitude"`
	TempUnit      string  `yaml:"temperature_unit"`
	WindSpeedUnit string  `yaml:"speed_unit"`
	PrecipUnit    string  `yaml:"precipitation_unit"`
	TimeZone      string  `yaml:"timezone"`
}
