package types

// Houses is a JSON representation of Home Vision API response
type Houses struct {
	Houses []House `json:"houses"`
	Ok     bool    `json:"ok"`
}
