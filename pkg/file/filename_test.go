package file

import (
	"github.com/stretchr/testify/assert"
	"github.com/victorlss/flaky-api/pkg/api/home-vision/types"
	"testing"
)

func TestGenerateFilename(t *testing.T) {
	t.Run("should generate file", func(t *testing.T) {
		expected := "id-000-4_Pumpkin_Hill_St_Antioch_TN_37013.jpg"
		houseMock := types.House{
			Id:        0,
			Address:   "4 Pumpkin Hill St. Antioch, TN 37013",
			Homeowner: "Nicole Bone",
			Price:     105124,
			PhotoURL:  "https://image.shutterstock.com/image-photo/big-custom-made-luxury-house-260nw-374099713.jpg",
		}

		actual := GenerateFilename(houseMock)

		assert.Equal(t, expected, actual)
	})
}
