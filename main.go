package main

import (
	"AP/router"
	"fmt"
)

//import "github.com/go-redis/redis"

func main() {
	//listenAddress := os.Getenv("LISTEN_URL")
	listenAddress := "localhost:8090"
	router.Start(listenAddress)
	fmt.Printf("Server listening on: %s ...", listenAddress)
	fmt.Println()

	select {}
}
