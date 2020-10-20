package tradeogre

import (
	"encoding/json"
)

type Balance struct {
	Success   bool   `json:"success"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
}
func (to *Keys) Balance(currency string) Balance {
	params := map[string]string{
		"currency": currency,
	}
	req := to.PrivateRequest("https://tradeogre.com/api/v1/account/balance", params, "")
	var res Balance
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
