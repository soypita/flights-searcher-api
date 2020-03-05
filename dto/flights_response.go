package dto

import "github.com/soypita/flights-searcher-api/entities"

type FlightResults struct {
	Flights  []entities.Flight `json:"flights"`
	Prices   []entities.Charge `json:"prices"`
	Currency string            `json:"currency"`
}

type FlightsResponse struct {
	Limit   int             `json:"limit"`
	Offset  int             `json:"offset"`
	Total   int             `json:"total"`
	Results []FlightResults `json:"results"`
}

type FlightsStatistics struct {
	OptimalFlight FlightResults `json:"optimalFlight"`
	CheapestFlight FlightResults `json:"cheapestFlight"`
	FastFlight FlightResults `json:"fastFlight"`
	SlowFlight FlightResults `json:"slowFlight"`
	ExpensiveFlight FlightResults `json:"expensiveFlight"`
}
