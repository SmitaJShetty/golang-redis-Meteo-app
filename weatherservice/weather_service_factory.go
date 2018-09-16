package weatherservice

//NewWeatherService returns a new weather service
func NewWeatherService() WeatherService {
	return &SimpleWeatherService{}
}
