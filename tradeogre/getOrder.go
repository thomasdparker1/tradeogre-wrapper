package tradeogre

import (
	"encoding/json"
)

type GetOrder struct {
	Success   bool   `json:"success"`
	Date      string `json:"date"`
	Type      string `json:"type"`
	Market    string `json:"market"`
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	Fulfilled string `json:"fulfilled"`
}

func (to *Keys) GetOrder(UUID string) GetOrder {
	req := to.PrivateRequest("https://tradeogre.com/api/v1/account/order/", nil, UUID)
	var res GetOrder
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
