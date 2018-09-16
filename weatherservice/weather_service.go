package weatherservice

import (
	"AP/common"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//WeatherServices list of 3rd party weather services known to the application's weather service
var WeatherServices []Provider
var currentAccessible int

const SERVICE_TIMEOUT int = 2

//WeatherService service interface
type WeatherService interface {
	GetWeather() (*common.WeatherResponse, error)
	LoadWeatherServices() error
}

//SimpleWeatherService simple weather service
type SimpleWeatherService struct {
}

//GetWeather gets weather service url from a list of services
func (sws *SimpleWeatherService) GetWeather() (*common.WeatherResponse, error) {
	provider, providerErr := sws.getWeatherServiceProvider()
	if providerErr != nil {
		providerErr := fmt.Errorf("GetWeatherServiceURL:Err:(%v)", providerErr)
		return nil, providerErr
	}

	if provider == nil {
		providerEmptyErr := fmt.Errorf("GetWeatherServiceURL: Weather service provider was empty")
		return nil, providerEmptyErr
	}

	weatherResponse, weatherResponseErr := sws.invokeWeatherAPI(provider)
	if weatherResponseErr != nil {
		return nil, weatherResponseErr
	}

	if weatherResponse == nil {
		return nil, fmt.Errorf("getWeather:Empty weatherresponse")
	}

	return weatherResponse, nil
}

func (sws *SimpleWeatherService) getWeatherServiceProvider() (*Provider, error) {
	if len(WeatherServices) == 0 {
		return nil, fmt.Errorf("No weather services")
	}

	serviceAvailable := false

	for !serviceAvailable {
		fmt.Println("looking for service", WeatherServices[currentAccessible].Type)
		isAlive, isAliveErr := sws.isAlive(&WeatherServices[currentAccessible])
		if isAliveErr != nil {
			fmt.Println(isAliveErr)
		}

		if !isAlive {
			currentAccessible++

			if currentAccessible >= len(WeatherServices) {
				currentAccessible = 0
			}
		} else {
			serviceAvailable = true
			break
		}
	}

	return &WeatherServices[currentAccessible], nil
}

func (sws *SimpleWeatherService) isAlive(provider *Provider) (bool, error) {
	if provider == nil {
		return false, fmt.Errorf("provider empty")
	}

	if provider.URL == "" {
		return false, fmt.Errorf("provider url empty")
	}

	timeOut := time.Duration(time.Duration(SERVICE_TIMEOUT) * time.Second)
	httpClient := http.Client{
		Timeout: timeOut,
	}
	response, connErr := httpClient.Get(provider.URL)
	if connErr != nil {
		return false, connErr
	}

	if response == nil {
		return false, fmt.Errorf("invalid response")
	}

	if response.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Response not OK")
	}

	return true, nil
}

func (sws *SimpleWeatherService) invokeWeatherAPI(provider *Provider) (*common.WeatherResponse, error) {
	var weatherResponse *common.WeatherResponse
	var weatherResponseErr error

	apiResponse, responseErr := http.Get(provider.URL)
	if responseErr != nil {
		return nil, responseErr
	}

	if apiResponse == nil {
		return nil, fmt.Errorf("Error while processing response, empty response")
	}

	switch provider.Type {
	case PROVIDER_YAHOO:
		var response common.YahooWeatherResponse
		parseErr := json.NewDecoder(apiResponse.Body).Decode(&response)
		if parseErr != nil {
			return nil, parseErr
		}

		weatherResponse, weatherResponseErr = sws.parseValues(response.Query.Results.Channel.Wind.Speed, response.Query.Results.Channel.Item.Condition.Temp)
		if weatherResponseErr != nil {
			return nil, weatherResponseErr
		}

	case PROVIDER_OPENWEATHER:
		var response common.OpenWeatherResponse
		fmt.Println(apiResponse.Body)
		parseErr := json.NewDecoder(apiResponse.Body).Decode(&response)
		if parseErr != nil {
			return nil, parseErr
		}

		weatherResponse = common.NewWeatherResponse(response.Wind.Speed, response.Wind.Degree)
	}

	return weatherResponse, nil
}

func (sws *SimpleWeatherService) parseValues(strSpeed string, strTemp string) (*common.WeatherResponse, error) {
	speed, speedParseErr := strconv.ParseFloat(strSpeed, 32)
	if speedParseErr != nil {
		return nil, speedParseErr
	}

	temp, tempParseErr := strconv.ParseFloat(strTemp, 32)
	if tempParseErr != nil {
		return nil, tempParseErr
	}

	return common.NewWeatherResponse(float32(speed), float32(temp)), nil
}

//LoadWeatherServices loads available weather services
func (sws *SimpleWeatherService) LoadWeatherServices() error {
	//gathers all weather services and loads them into an array, can be env variable or fed from another api
	url, resolveErr := sws.resolveLocation(common.YahooAPIURL, common.YahooSydneyWOEID)
	if resolveErr != nil {
		fmt.Println("LoadWeatherServices: yahoo url unresolved for location")
	}

	if url != "" {
		newWService := NewProvider(PROVIDER_YAHOO, url)
		WeatherServices = append(WeatherServices, *newWService)
	}

	url, resolveErr = sws.resolveLocation(common.OpenWeatherURL, common.OpenWeatherSydney)
	if resolveErr != nil {
		fmt.Println("LoadWeatherServices: Open weather url unresolved for location")
	}

	if url != "" {
		newWService := NewProvider(PROVIDER_OPENWEATHER, url)
		WeatherServices = append(WeatherServices, *newWService)
	}

	if len(WeatherServices) == 0 {
		return fmt.Errorf("LoadWeatherServices: Weather services not available")
	}

	return nil
}

//resolveLocation, resolves location in url; can be moved to a service if more values are required to be parsed
func (sws *SimpleWeatherService) resolveLocation(url string, cityCode string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("resolveLocation: empty location url")
	}

	return strings.Replace(url, "<CITYCODE>", cityCode, 1), nil
}
