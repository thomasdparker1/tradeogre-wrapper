package tradeogre

import (
	"encoding/json"
)

type TickerStruct struct {
	Success      bool   `json:"success"`
	Initialprice string `json:"initialprice"`
	Price        string `json:"price"`
	High         string `json:"high"`
	Low          string `json:"low"`
	Volume       string `json:"volume"`
	Bid          string `json:"bid"`
	Ask          string `json:"ask"`
}

func (to *Keys) Ticker(market string) TickerStruct {
	req := to.PrivateRequest("https://tradeogre.com/api/v1/ticker", nil,"/"+market)
	var ticker TickerStruct
	_ = json.Unmarshal(req.([]uint8), &ticker)
	return ticker
}
