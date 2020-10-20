package tradeogre

import (
	"encoding/json"
)

type Orders []struct {
	UUID     string `json:"uuid"`
	Date     int    `json:"date"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Market   string `json:"market"`
}

func (to *Keys) Orders(market string) Orders {
	params := map[string]string{
		"market": market,
	}
	req := to.PrivateRequest("https://tradeogre.com/api/v1/account/orders", params, "")
	var res Orders
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
