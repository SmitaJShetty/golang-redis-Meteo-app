package router

import (
	"AP/weatherservice"

	"github.com/gorilla/mux"
)

func addRoutes(r *mux.Router) {
	r.HandleFunc("/v1/weather", weatherservice.GetWeather).Methods("GET")
}
