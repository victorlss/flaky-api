package home_vision

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/victorlss/flaky-api/pkg/shared/types"
	"io/ioutil"
	"log"
)

var (
	baseUrl = "https://app-homevision-staging.herokuapp.com/api_project"
)

const (
	retries      = 5
	pagesToFetch = 10
)

// Fetch data from Home Vision API.
func Fetch(houseChan chan types.House) {
	pageNumber := 0
	for pageNumber < pagesToFetch {
		pageNumber++

		httpClient := retryablehttp.NewClient()
		httpClient.RetryMax = retries

		url := fmt.Sprintf("%s/houses?page=%d", baseUrl, pageNumber)
		resp, err := httpClient.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		houses := parse(body)

		for _, house := range houses.Houses {
			houseChan <- house
		}
	}

	close(houseChan)
}

// Parse data from json to Houses type
func parse(jsonData []byte) types.Houses {
	var houses types.Houses
	_ = json.Unmarshal(jsonData, &houses)

	return houses
}
