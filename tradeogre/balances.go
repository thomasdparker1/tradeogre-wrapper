package tradeogre

import (
	"encoding/json"
)

type Balances struct {
	Success  bool              `json:"success"`
	Balances map[string]string `json:"balances"`
}

func (to *Keys) Balances() Balances {
	req := to.PrivateRequest("https://tradeogre.com/api/v1/account/balances", nil, "")
	var res Balances
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
