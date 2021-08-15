package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

		expected := Houses{
			Houses: []House{
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

		actual := Houses{}.Parse([]byte(housesJsonMock))

		assert.Equal(t, expected, actual)
	})
}
