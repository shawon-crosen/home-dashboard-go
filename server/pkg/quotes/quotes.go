package quotes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const url = "https://thequoteshub.com/api/quotes"

type Quotes struct {
	Client http.Client
}

type Response struct {
	Pagination Pagination
	Quotes     []Quote
}

type Pagination struct {
	Next     string
	Page     int
	PageSize int
	Pages    int
	Total    int
}

type Quote struct {
	Author   string
	AuthorID string
	ID       int
	Tags     []string
	Text     string
}

func (q Quotes) getQuotes() (Response, error) {
	target := Response{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return target, err
	}

	resp, err := q.Client.Get(req.URL.String())

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&target)

	if err != nil {
		log.Println(err)
	}

	return target, nil
}

func (q Quotes) getQuotesByAuthor(aID string) (Response, error) {
	target := Response{}
	req, err := http.NewRequest("GET", url+"/authors"+"/"+aID, nil)
	if err != nil {
		return target, err
	}

	resp, err := q.Client.Get(req.URL.String())

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&target)

	if err != nil {
		log.Println(err)
	}

	return target, nil
}
