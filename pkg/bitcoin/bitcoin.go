package bitcoin

import (
	"encoding/json"
	"log"
	"net/http"
)

// PriceIndex of Bitcoin.
type PriceIndex struct {
	Time struct {
		Updated string `json:"updated"`
	} `json:"time"`
	PriceIndex map[string]struct {
		Code      string  `json:"code"`
		Symbol    string  `json:"symbol"`
		RateFloat float64 `json:"rate_float"`
	} `json:"bpi"`
}

// GetPriceIndex from CoinDesk.
func GetPriceIndex() *PriceIndex {
	res, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		log.Printf("Error fetch()ing: %v", err)
		return nil
	}
	defer res.Body.Close()

	var bpi *PriceIndex
	err = json.NewDecoder(res.Body).Decode(&bpi)
	if err != nil {
		log.Printf("Error JSON decoding: %v", err)
		return nil
	}
	return bpi
}
