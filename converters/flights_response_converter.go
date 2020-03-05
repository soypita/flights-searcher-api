package converters

import (
	"sync"

	"github.com/soypita/flights-searcher-api/dto"
	"github.com/soypita/flights-searcher-api/entities"
)

// ConvertToFlightsResponse - converter to FlightsResponse dto from api
func ConvertToFlightsResponse(flights []entities.Response) *dto.FlightsResponse {
	resp := &dto.FlightsResponse{}

	resp.Results = make([]dto.FlightResults, len(flights))
	wg := sync.WaitGroup{}
	for i := 0; i < len(flights); i++ {
		wg.Add(1)
		go fillFlightResult(&wg, &resp.Results[i], &flights[i])
	}

	wg.Wait()
	return resp
}

func fillFlightResult(wg *sync.WaitGroup, target *dto.FlightResults, source *entities.Response) {
	defer wg.Done()
	target.Flights = source.Flights.Flights[:]
	target.Prices = source.FlightPricing.Charges[:]
	target.Currency = source.FlightPricing.Currency[:]
}
