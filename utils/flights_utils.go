package utils

import (
	"github.com/soypita/flights-searcher-api/dto"
	"github.com/soypita/flights-searcher-api/entities"
	"sort"
	"sync"
)

func GetFlightStatistics(in []entities.Response) dto.FlightsStatistics {
	resStatistic := dto.FlightsStatistics{}
	wg := sync.WaitGroup{}

	arrToSortCheap := make([]entities.Response, len(in))
	copy(arrToSortCheap, in)
	arrToSortSlow := make([]entities.Response, len(in))
	copy(arrToSortSlow, in)
	arrToSortOptimal := make([]entities.Response, len(in))
	copy(arrToSortOptimal, in)

	wg.Add(3)

	go findCheapestFlight(&wg, arrToSortCheap)
	go findSlowFLight(&wg, arrToSortSlow)
	go findOptimalFlight(&wg, arrToSortOptimal)

	wg.Wait()
	collectDataToStatisticResponse(arrToSortCheap, arrToSortSlow, arrToSortOptimal, &resStatistic)
	return resStatistic
}

func findCheapestFlight(wg *sync.WaitGroup, in []entities.Response) {
	defer wg.Done()
	sort.Sort(entities.CheapFlightSorter(in))
}

func findSlowFLight(wg *sync.WaitGroup, in []entities.Response) {
	defer wg.Done()
	sort.Sort(entities.SlowFlightSorter(in))
}

func findOptimalFlight(wg *sync.WaitGroup, in []entities.Response) {
	defer wg.Done()
	sort.Sort(entities.OptimalFlightSorter(in))
}

func collectDataToStatisticResponse(arrToSortCheap []entities.Response, arrToSortSlow []entities.Response, arrToSortOptimal []entities.Response, resStatistic *dto.FlightsStatistics) {
	resStatistic.CheapestFlight.Flights = make([]entities.Flight, len(arrToSortCheap[0].Flights.Flights))
	resStatistic.CheapestFlight.Prices = make([]entities.Charge, len(arrToSortCheap[0].FlightPricing.Charges))
	copy(resStatistic.CheapestFlight.Flights, arrToSortCheap[0].Flights.Flights)
	copy(resStatistic.CheapestFlight.Prices, arrToSortCheap[0].FlightPricing.Charges)
	resStatistic.CheapestFlight.Currency = arrToSortCheap[0].FlightPricing.Currency

	resStatistic.ExpensiveFlight.Flights = make([]entities.Flight, len(arrToSortCheap[len(arrToSortCheap)-1].Flights.Flights))
	resStatistic.ExpensiveFlight.Prices = make([]entities.Charge, len(arrToSortCheap[len(arrToSortCheap)-1].FlightPricing.Charges))
	copy(resStatistic.ExpensiveFlight.Flights, arrToSortCheap[len(arrToSortCheap)-1].Flights.Flights)
	copy(resStatistic.ExpensiveFlight.Prices, arrToSortCheap[len(arrToSortCheap)-1].FlightPricing.Charges)
	resStatistic.ExpensiveFlight.Currency = arrToSortCheap[len(arrToSortCheap)-1].FlightPricing.Currency

	resStatistic.SlowFlight.Flights = make([]entities.Flight, len(arrToSortSlow[0].Flights.Flights))
	resStatistic.SlowFlight.Prices = make([]entities.Charge, len(arrToSortSlow[0].FlightPricing.Charges))
	copy(resStatistic.SlowFlight.Flights, arrToSortSlow[0].Flights.Flights)
	copy(resStatistic.SlowFlight.Prices, arrToSortSlow[0].FlightPricing.Charges)
	resStatistic.SlowFlight.Currency = arrToSortSlow[0].FlightPricing.Currency

	resStatistic.FastFlight.Flights = make([]entities.Flight, len(arrToSortSlow[len(arrToSortSlow)-1].Flights.Flights))
	resStatistic.FastFlight.Prices = make([]entities.Charge, len(arrToSortSlow[len(arrToSortSlow)-1].FlightPricing.Charges))
	copy(resStatistic.FastFlight.Flights, arrToSortSlow[len(arrToSortSlow)-1].Flights.Flights)
	copy(resStatistic.FastFlight.Prices, arrToSortSlow[len(arrToSortSlow)-1].FlightPricing.Charges)
	resStatistic.FastFlight.Currency = arrToSortSlow[len(arrToSortSlow)-1].FlightPricing.Currency

	resStatistic.OptimalFlight.Flights = make([]entities.Flight, len(arrToSortOptimal[0].Flights.Flights))
	resStatistic.OptimalFlight.Prices = make([]entities.Charge, len(arrToSortOptimal[0].FlightPricing.Charges))
	copy(resStatistic.OptimalFlight.Flights, arrToSortOptimal[0].Flights.Flights)
	copy(resStatistic.OptimalFlight.Prices, arrToSortOptimal[0].FlightPricing.Charges)
	resStatistic.OptimalFlight.Currency = arrToSortOptimal[0].FlightPricing.Currency
}
