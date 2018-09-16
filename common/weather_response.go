package common

//WeatherResponse weather response construct
type WeatherResponse struct {
	WindSpeed   float32 `json:"wind_speed"`
	Temperature float32 `json:"temperature_degrees"`
}

//NewWeatherResponse returns a new weather response
func NewWeatherResponse(windSpeed float32, temperature float32) *WeatherResponse {
	return &WeatherResponse{
		WindSpeed:   windSpeed,
		Temperature: temperature,
	}
}
