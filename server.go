package main

import (
	"golang-fifa-world-cup-web-service/handlers"
	"net/http"
)

func main() {
	//data.PrintUsage()

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/winners", handlers.WinnersHandler)
	http.ListenAndServe(":35000", nil)
}
