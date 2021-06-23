package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// load and read config
	config := readConfigFromEnviroment()

	// create http router
	router := mux.NewRouter()

	// start web service
	if err := http.ListenAndServe(config.Host + ":" + config.Port, router); err != nil {
		panic(err)
	}

	fmt.Println(config)
}