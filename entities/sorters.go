package entities

import (
	"log"
	"time"
)

const chargeTypeToCompare = "TotalAmount"
const typeToCompare = "SingleAdult"

type CheapFlightSorter []Response

func (f CheapFlightSorter) Len() int      { return len(f) }
func (f CheapFlightSorter) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f CheapFlightSorter) Less(i, j int) bool {
	firstPrices := f[i].FlightPricing.Charges
	secondPrices := f[j].FlightPricing.Charges

	firstAmount, secondAmount := extractAmountToCompare(firstPrices, secondPrices)

	return firstAmount < secondAmount
}

type SlowFlightSorter []Response

func (f SlowFlightSorter) Len() int      { return len(f) }
func (f SlowFlightSorter) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f SlowFlightSorter) Less(i, j int) bool {
	firstFlights := f[i].Flights.Flights
	secondFlights := f[j].Flights.Flights

	firstDuration, err := extractDurationTimeForFlight(firstFlights)

	if err != nil {
		log.Println(err)
		return false
	}

	secondDuration, err := extractDurationTimeForFlight(secondFlights)

	return firstDuration > secondDuration
}

type OptimalFlightSorter []Response

func (f OptimalFlightSorter) Len() int      { return len(f) }
func (f OptimalFlightSorter) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f OptimalFlightSorter) Less(i, j int) bool {
	firstFlights := f[i].Flights.Flights
	secondFlights := f[j].Flights.Flights
	firstPrices := f[i].FlightPricing.Charges
	secondPrices := f[j].FlightPricing.Charges

	firstDuration, err := extractDurationTimeForFlight(firstFlights)

	if err != nil {
		log.Println(err)
		return false
	}

	secondDuration, err := extractDurationTimeForFlight(secondFlights)
	firstAmount, secondAmount := extractAmountToCompare(firstPrices, secondPrices)

	return  firstAmount + float32(firstDuration) < secondAmount + float32(secondDuration)
}

func extractDurationTimeForFlight(flight []Flight) (time.Duration, error) {
	layout := "2006-01-02T1504"
	startTime, err := time.Parse(layout, flight[0].ArrivalTime)
	if err != nil {
		return 0, err
	}
	finishTime, err := time.Parse(layout, flight[len(flight) - 1].ArrivalTime)
	if err != nil {
		return 0, err
	}
	diff := finishTime.Sub(startTime)
	return diff, nil
}

func extractAmountToCompare(firstCharges []Charge, secondCharge []Charge) (float32, float32) {
	var firstAmount float32 = 0.0
	var secondAmount float32 = 0.0

	for i := 0; i < len(firstCharges); i++ {
		if firstCharges[i].ChargeType == chargeTypeToCompare && firstCharges[i].Type == typeToCompare {
			firstAmount = firstCharges[i].Amount
		}
	}

	for i := 0; i < len(secondCharge); i++ {
		if secondCharge[i].ChargeType == chargeTypeToCompare && secondCharge[i].Type == typeToCompare {
			secondAmount = secondCharge[i].Amount
		}
	}

	return firstAmount, secondAmount
}
