package quotes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

const url = "https://thequoteshub.com/api/"

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
	Author   string   `json:"author"`
	AuthorID string   `json:"author_id"`
	ID       int      `json:"id"`
	Tags     []string `json:"tags"`
	Text     string   `json:"text"`
}

func (q Quotes) getQuoteNum() int {
	target := Response{}
	req, err := http.NewRequest("GET", url+"authors/"+"Toni+Morrison", nil)
	if err != nil {
		return 0
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

	return target.Pagination.Total
}

func (q Quotes) getQuotesByTag(size int) (Response, error) {
	target := Response{}
	req, err := http.NewRequest("GET", url+"tags"+fmt.Sprintf("/art?page=1&page_size=%v", size), nil)
	if err != nil {
		return target, err
	}

	query := req.URL.Query()
	resp, err := q.Client.Get(req.URL.String())
	req.URL.RawQuery = query.Encode()

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

func (q Quotes) getQuotesByAuthor(size int) (Response, error) {
	target := Response{}
	req, err := http.NewRequest("GET", url+"authors"+fmt.Sprintf("/Toni+Morrison?page=1&page_size=%v", size), nil)
	if err != nil {
		return target, err
	}

	query := req.URL.Query()
	resp, err := q.Client.Get(req.URL.String())
	req.URL.RawQuery = query.Encode()

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

func (q Quotes) QueryForQuote() *Quote {
	size := q.getQuoteNum()
	if size < 1 {
		fmt.Println("No quotes")
	}
	quotes, err := q.getQuotesByAuthor(size)
	if err != nil {
		fmt.Println("Error getting quotes")
	}

	r := rand.IntN(size)

	if quotes.Quotes[r].Author == "Neil Gaiman" {
		r = rand.IntN(size)
	}

	return &quotes.Quotes[r]
}
