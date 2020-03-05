package entities

type Flight struct {
	Carrier struct {
		ID          string `xml:"id,attr" json:"carrierId"`
		CarrierName string `xml:",chardata" json:"carrierName"`
	} `xml:"Carrier" json:"carrier"`
	FlightNumber  int    `xml:"FlightNumber" json:"flightNumber"`
	Source        string `xml:"Source" json:"source"`
	Destination   string `xml:"Destination" json:"destination"`
	DepartureTime string `xml:"DepartureTimeStamp" json:"departureTime"`
	ArrivalTime   string `xml:"ArrivalTimeStamp" json:"arrivalTime"`
	Class         string `xml:"Class" json:"flightClass"`
	NumberOfStops int    `xml:"NumberOfStops" json:"stops"`
	TicketType    string `xml:"TicketType" json:"ticketType"`
}

type Charge struct {
	Amount     float32 `xml:",chardata" json:"amount"`
	Type       string  `xml:"type,attr" json:"type"`
	ChargeType string  `xml:"ChargeType,attr" json:"chargeType"`
}

type Pricing struct {
	Currency string   `xml:"currency,attr"`
	Charges  []Charge `xml:"ServiceCharges"`
}

type Response struct {
	Flights struct {
		Flights []Flight `xml:"Flight"`
	} `xml:"OnwardPricedItinerary>Flights"`
	FlightPricing Pricing `xml:"Pricing"`
}

type FlightsResponse struct {
	Resp []Response `xml:"PricedItineraries>Flights"`
}
