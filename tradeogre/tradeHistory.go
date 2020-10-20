package tradeogre

import (
	"encoding/json"
)

type TradeHistoryStruct []struct {
	Date     int    `json:"date"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

func (to *Keys) TradeHistory(market string) TradeHistoryStruct {
	req := to.PrivateRequest("https://tradeogre.com/api/v1/history", nil,"/"+market)
	var tradeHistory TradeHistoryStruct
	_ = json.Unmarshal(req.([]uint8), &tradeHistory)
	return tradeHistory
}