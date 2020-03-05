package main

import (
	"github.com/soypita/flights-searcher-api/entities"
	"github.com/soypita/flights-searcher-api/gateways"
	"sort"
	"sync"
	"testing"
)

const (
	testRawResponse = "resources/test/Response_test.xml"
	rawResponseSize = 3
	rawChargesSize = 3
	cheapestFlights = 546.80
	expensiveFlights = 1379.80
	slowestFlightPrice = 546.80
	fastestFlightPrice = 1379.80
	optimalCarrierName = "Emirates"
)

func initTest() (*entities.FlightsResponse, error) {
	wg := sync.WaitGroup{}
	extractData := entities.FlightsResponse{}
	var err error
	wg.Add(1)
	go gateways.ExtractDataFromRawResponse(&wg, testRawResponse, &extractData, err)
	wg.Wait()
	return &extractData, err
}

func TestExtractDataFromRawResponse(t *testing.T)  {
	extractData, err := initTest()
	if err != nil {
		t.Error("Failed to extract data from raw response")
	}

	if len(extractData.Resp) != rawResponseSize {
		t.Error("Incorrect size of extracted data from raw response", "expected", rawResponseSize, "get", len(extractData.Resp))
	}

	if len(extractData.Resp[0].FlightPricing.Charges) != rawChargesSize {
		t.Error("Incorrect size of extracted charges from raw response", "expected", rawChargesSize, "get", len(extractData.Resp[0].FlightPricing.Charges))
	}
}

func TestSortersCheapExpensiveForFlightResponses(t *testing.T) {
	extractData, err := initTest()
	if err != nil {
		t.Error("Failed to extract data from raw response")
	}

	sort.Sort(entities.CheapFlightSorter(extractData.Resp))

	if extractData.Resp[0].FlightPricing.Charges[2].Amount != cheapestFlights {
		t.Error("Incorrect cheap sorted flight", "expected", cheapestFlights, "get", extractData.Resp[0].FlightPricing.Charges[2].Amount)
	}

	if extractData.Resp[2].FlightPricing.Charges[2].Amount != expensiveFlights {
		t.Error("Incorrect cheap sorted flight", "expected", expensiveFlights, "get", extractData.Resp[2].FlightPricing.Charges[2].Amount)
	}
}

func TestSortersFastSlowForFlightResponses(t *testing.T) {
	extractData, err := initTest()
	if err != nil {
		t.Error("Failed to extract data from raw response")
	}

	sort.Sort(entities.SlowFlightSorter(extractData.Resp))

	if extractData.Resp[0].FlightPricing.Charges[2].Amount != slowestFlightPrice {
		t.Error("Incorrect slow sorted flight", "expected", slowestFlightPrice, "get", extractData.Resp[0].FlightPricing.Charges[2].Amount)
	}

	if extractData.Resp[2].FlightPricing.Charges[2].Amount != fastestFlightPrice {
		t.Error("Incorrect fast sorted flight", "expected", fastestFlightPrice, "get", extractData.Resp[2].FlightPricing.Charges[2].Amount)
	}
}

func TestSortersOptimalFlightResponses(t *testing.T) {
	extractData, err := initTest()
	if err != nil {
		t.Error("Failed to extract data from raw response")
	}

	sort.Sort(entities.OptimalFlightSorter(extractData.Resp))

	if extractData.Resp[0].Flights.Flights[0].Carrier.CarrierName != optimalCarrierName {
		t.Error("Incorrect optimal carrier name", "expected", optimalCarrierName, "get", extractData.Resp[0].Flights.Flights[0].Carrier.CarrierName)
	}
}