package types

import (
	"encoding/json"
)

type House struct {
	Id        int    `json:"id"`
	Address   string `json:"address"`
	Homeowner string `json:"homeowner"`
	Price     int    `json:"price"`
	PhotoURL  string `json:"photoURL"`
}

type Houses struct {
	Houses []House `json:"houses"`
	Ok     bool    `json:"ok"`
}

func (houses Houses) Parse(jsonData []byte) Houses {
	err := json.Unmarshal(jsonData, &houses)
	if err != nil {
		// TODO: Implement log
	}

	return houses
}
