package common

const OpenWeatherURL string = "http://api.openweathermap.org/data/2.5/weather?q=<CITYCODE>&appid=2326504fb9b100bee21400190e4dbe6d"
const OpenWeatherSydney string = "Sydney, AU"

type OpenWeatherResponse struct {
	ID          uint         `json:"id"`
	Name        string       `json:"name"`
	Cod         int          `json:"cod"`
	Coordinates Coordinates  `json:"coord"`
	Weather     []WeatherObj `json:"weather"`
	Base        string       `jsons:"base"`
	Main        Main         `json:"main"`
	Visibility  float64      `json:"visibility"`
	Wind        Wind         `json:"wind"`
	Dt          uint         `json:"dt"`
	Sys         Sys          `json:"sys"`
}

type Coordinates struct {
	Longitude float32 `json:"long"`
	Latitude  float32 `json:"lat"`
}

type Weather struct {
	WeatherObjs []WeatherObj `json:"weather"`
}

type WeatherObj struct {
	ID          uint   `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Pressure float32 `json:"pressure"`
	Humidity float32 `json:"humidity"`
	TempMin  float32 `json:"temp_min"`
	TempMax  float32 `json:"temp_max"`
}

type Wind struct {
	Speed  float32 `json:"speed"`
	Degree float32 `json:"deg"`
}

type Clouds struct {
	All uint `json:"all"`
}

type Sys struct {
	Type    int     `json:"type"`
	ID      uint    `json:"id"`
	Message float32 `json:"message"`
	Country string  `json:"country"`
	Sunrise uint    `json:"sunrise"`
	Sunset  uint    `json:"sunset"`
}
