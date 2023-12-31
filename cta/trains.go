package cta

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Stops
// LaSalle 41340
// Harrison 41490
// Roosevelt 41400

const url = "http://lapi.transitchicago.com/api/1.0/ttarrivals.aspx"

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

	err = json.NewDecoder(resp.Body).Decode(&target)

	if err != nil {
		fmt.Println("response packing error")
		fmt.Println(err)
	}

	return target
}

func (s Station) FormatData(sr StationResponse) StationData {
	td := StationData{}
	nRoutes := make(map[string][]TrainData)
	sRoutes := make(map[string][]TrainData)

	// for _, t := range sr.Root.Etas {
	// 	routes[t.RouteColor] = []TrainData{}
	// }

	for _, t := range sr.Root.Etas {
		r := TrainData{}

		// bool conversions
		app, _ := strconv.Atoi(t.Approaching)
		delay, _ := strconv.Atoi(t.Delayed)
		fault, _ := strconv.Atoi(t.Fault)
		scheduled, _ := strconv.Atoi(t.Scheduled)

		r.Color = t.RouteColor
		r.StopName = t.StationName
		r.StopDescription = t.StopDescription
		r.DestinationNumber = t.DestinationNumber
		r.Destinationname = t.DestinationName
		r.Directioncode = t.Directioncode

		r.Approaching = app != 0
		r.Delayed = delay != 0
		r.Fault = fault != 0
		r.Scheduled = scheduled != 0

		r.PredictionGenerated = t.PredictionGenerated
		r.ArrivalTime = t.ArrivalTime

		r.ArrivalMinutes = parseArrival(t.ArrivalTime)

		switch t.Directioncode {
		case "1":
			nRoutes[t.RouteColor] = append(nRoutes[t.RouteColor], r)
		case "5":
			sRoutes[t.RouteColor] = append(sRoutes[t.RouteColor], r)
		}
	}

	td.North = nRoutes
	td.South = sRoutes

	return td
}

func parseArrival(at string) float64 {
	now := time.Now()
	arr, err := time.ParseInLocation("2006-01-02T15:04:05", at, time.Local)
	if err != nil {
		log.Println(err)
	}

	return math.Round(arr.Sub(now).Minutes())
}
