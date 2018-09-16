package weatherservice

import (
	"AP/common"
	"AP/datastoreservice"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const EXPIRY int = 3

//GetWeather gets weather
func GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		city = "Sydney"
	}

	weatherSvc := NewWeatherService()

	weatherSvc.LoadWeatherServices()

	cacheSvc := datastoreservice.NewCacheService()
	cachedResponse, cacheErr := cacheSvc.Get(city)
	if cacheErr != nil {
		fmt.Printf(cacheErr.Error())
	}

	if cachedResponse != nil {
		fmt.Println("cachedResponse")
		response, _ := json.Marshal(*cachedResponse)
		common.SendResult(w, r, response)

	} else {
		weatherResponse, weatherResponseErr := weatherSvc.GetWeather()
		if weatherResponseErr != nil {
			common.SendErrorResponse(w, r, common.NewAppError(fmt.Sprintf("GetWeather: Error while fetching weather details from service; Err:(%v)", weatherResponseErr), http.StatusInternalServerError))
			return
		}

		if weatherResponse == nil {
			common.SendErrorResponse(w, r, common.NewAppError("GetWeather: Error while fetching weather details from service", http.StatusInternalServerError))
			return
		}

		setErr := cacheSvc.Set(city, weatherResponse, time.Duration(EXPIRY)*time.Second)
		if setErr != nil {
			fmt.Println(setErr)
		}

		response, _ := json.Marshal(weatherResponse)
		common.SendResult(w, r, response)
	}
}
