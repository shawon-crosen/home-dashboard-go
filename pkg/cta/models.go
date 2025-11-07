package cta

import (
	"net/http"
)

type AllTrains struct {
	Stations []int
	ApiKey   string
	Client   http.Client
}

type AllTrainsData struct {
	StationResponse []TrainData
}

type StationResponse struct {
	Root TrainResponse `json:"ctatt"`
}

type TrainResponse struct {
	GenTime     string `json:"tmst"`
	ErrorCode   string `json:"errCd"`
	ErrorNumber int    `json:"errNm"`
	Etas        []Eta  `json:"eta"`
}

type Eta struct {
	StationId           string `json:"staId"`
	StopId              string `json:"stpId"`
	StationName         string `json:"staNm"`
	StopDescription     string `json:"stpDe"`
	RouteNumber         string `json:"rn"`
	RouteColor          string `json:"rt"`
	DestinationNumber   string `json:"destSt"`
	DestinationName     string `json:"destNm"`
	Directioncode       string `json:"trDr"`
	PredictionGenerated string `json:"prdt"`
	ArrivalTime         string `json:"arrT"`
	Approaching         string `json:"isApp"`
	Scheduled           string `json:"isSch"`
	Delayed             string `json:"isDly"`
	Fault               string `json:"isFlt"`
	Latitude            string `json:"lat"`
	Longitude           string `json:"lon"`
	Heading             string `json:"heading"`
}

type TrainData struct {
	StopName            string
	StopDescription     string
	Color               string
	DestinationNumber   string
	Destinationname     string
	PredictionGenerated string
	ArrivalTime         string
	Directioncode       string
	ArrivalMinutes      float64
	Approaching         bool
	Scheduled           bool
	Delayed             bool
	Fault               bool
}
