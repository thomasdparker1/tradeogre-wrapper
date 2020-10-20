package tradeogre

import (
	"encoding/json"
	"fmt"
)

type SellStruct struct {
	Success      bool   `json:"success"`
	UUID         string `json:"uuid"`
	Bnewbalavail string `json:"bnewbalavail"`
	Snewbalavail string `json:"snewbalavail"`
}

func (to *Keys) SellOrder(market string, price float64, amount float64) SellStruct {
	params := map[string]string{
		"market": market,
		"quantity": fmt.Sprintf("%.8f", amount),
		"price": fmt.Sprintf("%.8f", price),
	}
	req := to.PrivateRequest("https://tradeogre.com/api/v1/order/sell", params, "")
	var res SellStruct
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
