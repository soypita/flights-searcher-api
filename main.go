package main

import (
	"github.com/gorilla/mux"
	"github.com/soypita/flights-searcher-api/controllers"
	"log"
	"net/http"
)

func main() {
	controller := controllers.FlightsController{}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/flights/all", controller.GetAllFlightsResponse).Methods("GET")
	router.HandleFunc("/flights/stats", controller.GetStatsForFlightsResponses).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
