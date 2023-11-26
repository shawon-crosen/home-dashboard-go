package cta

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Stops
// LaSalle 41340
// Harrison 41490
// Roosevelt 41400

const url = "http://lapi.transitchicago.com/api/1.0/ttarrivals.aspx"

type Station struct {
	Client http.Client
	Id     int
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

func (s Station) GetTrains() StationResponse {
	target := StationResponse{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
	}

	q := req.URL.Query()
	q.Add("mapid", fmt.Sprint(s.Id))
	q.Add("key", os.Getenv("CTA_API_KEY"))
	q.Add("outputType", "JSON")

	req.URL.RawQuery = q.Encode()

	resp, err := s.Client.Get(req.URL.String())

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	// r, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(r))

	err = json.NewDecoder(resp.Body).Decode(&target)

	if err != nil {
		fmt.Println("response packing error")
		fmt.Println(err)
	}

	return target
}
