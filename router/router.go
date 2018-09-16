package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func Start(listenAddress string) {
	r := mux.NewRouter()

	addRoutes(r)

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
	go func() {
		err := http.ListenAndServe(listenAddress, r)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Listening on port:", listenAddress)
	}()
}
