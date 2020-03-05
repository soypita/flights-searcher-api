package gateways

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"sync"

	"github.com/soypita/flights-searcher-api/entities"
)

// In real application we assume that we get data from gateway dynamically
const (
	firstGatewayResponse  = "resources/Response_1.xml"
	secondGatewayResponse = "resources/Response_2.xml"
)

func extractDataFromRawResponse(wg *sync.WaitGroup, gatewayResponse string, response *entities.FlightsResponse, err error) {
	xmlFile, err := os.Open(gatewayResponse)
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, response)
	wg.Done()
}

func prepareData() ([]entities.Response, error) {
	wg := sync.WaitGroup{}
	firstResponse := entities.FlightsResponse{}
	secondResponse := entities.FlightsResponse{}
	wg.Add(2)

	var err1 error
	var err2 error
	go extractDataFromRawResponse(&wg, firstGatewayResponse, &firstResponse, err1)
	go extractDataFromRawResponse(&wg, secondGatewayResponse, &secondResponse, err2)

	wg.Wait()
	if err1 != nil {
		return nil, err1
	}

	if err2 != nil {
		return nil, err2
	}

	result := append(firstResponse.Resp, secondResponse.Resp...)

	return result, nil
}

var prepData []entities.Response = nil

func GetAllFlightTickets() (resp []entities.Response, err error) {
	// mock call to real sales gateway
	if prepData == nil {
		resp, err = prepareData()
	}
	return resp, err
}
