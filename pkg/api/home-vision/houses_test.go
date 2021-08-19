package homeVision

import (
	"github.com/stretchr/testify/assert"
	"github.com/victorlss/flaky-api/pkg/shared/types"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchHouses(t *testing.T) {
	t.Run("should fetch houses data", func(t *testing.T) {
		expected := types.House{
			Id:        0,
			Address:   "4 Pumpkin Hill Street Antioch, TN 37013",
			Homeowner: "Nicole Bone",
			Price:     105124,
			PhotoURL:  "https://image.shutterstock.com/image-photo/big-custom-made-luxury-house-260nw-374099713.jpg",
		}

		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = io.WriteString(w, `{
					"houses": [{
						"id": 0,
						"address": "4 Pumpkin Hill Street Antioch, TN 37013",
						"homeowner": "Nicole Bone",
						"price": 105124,
						"photoURL": "https://image.shutterstock.com/image-photo/big-custom-made-luxury-house-260nw-374099713.jpg"
					}],
					"ok": true
				}`)
			}),
		)

		defer server.Close()
		baseUrl = server.URL
		houseChan := make(chan types.House)

		go Fetch(houseChan)
		actual := <-houseChan

		assert.Equal(t, expected, actual)
	})
}

func TestJsonParse(t *testing.T) {
	t.Run("should parse json houses data", func(t *testing.T) {
		housesJsonMock := `{
			"houses": [{
				"id": 0,
				"address": "4 Pumpkin Hill Street Antioch, TN 37013",
				"homeowner": "Nicole Bone",
				"price": 105124,
				"photoURL": "https://image.shutterstock.com/image-photo/big-custom-made-luxury-house-260nw-374099713.jpg"
			}],
			"ok": true
		}`

		expected := types.Houses{
			Houses: []types.House{
				{
					Id:        0,
					Address:   "4 Pumpkin Hill Street Antioch, TN 37013",
					Homeowner: "Nicole Bone",
					Price:     105124,
					PhotoURL:  "https://image.shutterstock.com/image-photo/big-custom-made-luxury-house-260nw-374099713.jpg",
				},
			},
			Ok: true,
		}

		actual := parse([]byte(housesJsonMock))

		assert.Equal(t, expected, actual)
	})
}
