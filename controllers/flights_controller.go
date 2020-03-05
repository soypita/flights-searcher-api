package controllers

import (
	"encoding/json"
	"github.com/soypita/flights-searcher-api/converters"
	"github.com/soypita/flights-searcher-api/gateways"
	"github.com/soypita/flights-searcher-api/utils"
	"log"
	"net/http"
	"strconv"
)

type FlightsController struct {
}

func (c *FlightsController) GetAllFlightsResponse(w http.ResponseWriter, r *http.Request) {
	limitValue := r.FormValue("limit")
	startValue := r.FormValue("offset")
	start := 0
	finish := 0
	limit := 0
	var err error

	if startValue != "" {
		start, err = strconv.Atoi(startValue)
		if err != nil {
			log.Printf("Failed to parse query param, %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	if limitValue != "" {
		limit, err = strconv.Atoi(limitValue)
		if err != nil {
			log.Printf("Failed to parse query param, %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		finish = limit + start
	}

	gatewayResponse, err := gateways.GetAllFlightTickets()
	gatewayResponseLen := len(gatewayResponse)

	if err != nil {
		log.Printf("Error to get data from gateway, %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if limit == 0 || limit > gatewayResponseLen {
		limit = gatewayResponseLen
	}

	if finish == 0 || finish > gatewayResponseLen {
		finish = gatewayResponseLen
	}

	resp := converters.ConvertToFlightsResponse(gatewayResponse[start:finish])
	resp.Limit = limit
	resp.Offset = start
	resp.Total = gatewayResponseLen
	json.NewEncoder(w).Encode(resp)
}

func (c *FlightsController) GetStatsForFlightsResponses(w http.ResponseWriter, r *http.Request) {
	gatewayResponse, err := gateways.GetAllFlightTickets()

	if err != nil {
		log.Printf("Error to get data from gateway, %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(utils.GetFlightStatistics(gatewayResponse))
}
