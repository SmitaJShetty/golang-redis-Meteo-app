package common

import (
	"time"
)

const YahooSydneyWOEID string = "1105779"
const YahooAPIURL string = "https://query.yahooapis.com/v1/public/yql?q=select%20item.condition%2C%20wind%20from%20weather.forecast%20where%20woeid%20%3D%20<CITYCODE>&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"

type YahooWeatherResponse struct {
	Query Query `json:"query"`
}

type Query struct {
	Count    int       `json:"count"`
	Created  time.Time `json:"created"`
	Language string    `json:"lang"`
	Results  Results   `json:"results"`
}

type Results struct {
	Channel Channel `json:"channel"`
}

type Channel struct {
	Wind YahooWind `json:"wind"`
	Item Item      `json:"item"`
}

type YahooWind struct {
	Chill     string `json:"chill"`
	Direction string `json:"direction"`
	Speed     string `json:"speed"`
}

type Item struct {
	Condition Condition `json:"condition"`
}

type Condition struct {
	Code string `json:"code"`
	Date string `json:"date"` //TODO: convert to date
	Temp string `json:"temp"`
	Text string `json:"text"`
}
