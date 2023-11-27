package config

type Config struct {
	WeatherConfig WeatherConfig `yaml:"weather"`
	CtaConfig     CtaConfig     `yaml:"cta"`
}

type WeatherConfig struct {
	Latitude      float64 `yaml:"latitude"`
	Longitude     float64 `yaml:"longitude"`
	TempUnit      string  `yaml:"temperature_unit"`
	WindSpeedUnit string  `yaml:"speed_unit"`
	PrecipUnit    string  `yaml:"precipitation_unit"`
	TimeZone      string  `yaml:"timezone"`
}

type CtaConfig struct {
	Stations []int  `yaml:"station_ids"`
	Api_key  string `yaml:"api_key"`
}
