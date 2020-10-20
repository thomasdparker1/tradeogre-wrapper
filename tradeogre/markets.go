package tradeogre

import (
	"encoding/json"
)

func (to *Keys) Markets() map[string]map[string]string {
	req := to.PrivateRequest("https://tradeogre.com/api/v1/markets", nil, "")
	var marketResInterface []map[string]map[string]string
	sortedRes := make(map[string]map[string]string)
	_ = json.Unmarshal(req.([]uint8), &marketResInterface)
	for k := range marketResInterface {
		for x, y := range marketResInterface[k] {
			sortedRes[x] = y
		}
	}
	return sortedRes
}
