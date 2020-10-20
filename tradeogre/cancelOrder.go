package tradeogre

import (
	"encoding/json"
)

type CancelOrderStruct struct {
	Success      bool   `json:"success"`
}

func (to *Keys) CancelOrder(UUID string) CancelOrderStruct {
	params := map[string]string{
		"uuid": UUID,
	}
	req := to.PrivateRequest("https://tradeogre.com/api/v1/order/cancel", params, "")
	var res CancelOrderStruct
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
