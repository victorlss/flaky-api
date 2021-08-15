package home_vision

import (
	"github.com/victorlss/flaky-api/pkg/api/home-vision/types"
	"io/ioutil"
	"net/http"
)

// Fetch houses data from Home Vision API.
func Fetch(url string) types.Houses {
	res, err := http.Get(url)
	if err != nil {
		// TODO: Implement log
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// TODO: Implement log
	}

	return types.Houses{}.Parse(body)
}
