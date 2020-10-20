package tradeogre

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

type OrderBookStruct struct {
	Buy  map[string]string `json:"buy"`
	Sell map[string]string `json:"sell"`
}

func (to *Keys) OrderBook(market string) map[string][][]float64{
	req := to.PrivateRequest("https://tradeogre.com/api/v1/orders", nil,"/"+market)
	var toBook OrderBookStruct
	_ = json.Unmarshal(req.([]uint8), &toBook)
	b := make(chan [][]float64)
	a := make(chan [][]float64)
	go to.filterOrderBook(toBook.Buy, true, b)
	go to.filterOrderBook(toBook.Sell, false, a)
	return map[string][][]float64{"bids": <-b, "asks": <-a,}
}

func (to *Keys) filterOrderBook(TOBook map[string]string, reverse bool, c chan [][]float64) [][]float64 {
	var prices []float64
	for price, _ := range TOBook {
		p, _ := strconv.ParseFloat(price, 64)
		prices = append(prices, p)
	}
	sort.Float64s(prices)
	if reverse {
		var reversedPrices []float64
		for i := 1; i < len(prices); i++ {
			reversedPrices = append(reversedPrices, prices[len(prices)-i])
		}
		prices = reversedPrices
	}

	var items [][]float64
	for _, price := range prices {
		a, _ := TOBook[fmt.Sprintf("%.8f", price)]
		amount, _ := strconv.ParseFloat(a, 64)
		items = append(items, []float64{price, amount})
	}
	c <- items
	return items
}
