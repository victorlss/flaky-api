package types

// House is a JSON representation of house data
type House struct {
	Id        int    `json:"id"`
	Address   string `json:"address"`
	Homeowner string `json:"homeowner"`
	Price     int    `json:"price"`
	PhotoURL  string `json:"photoURL"`
}
