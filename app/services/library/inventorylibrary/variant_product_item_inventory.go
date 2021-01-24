package inventorylibrary

import (
	"encoding/json"
	"strings"
)

type VPIInventory struct {
	ID         uint       `json:"id"`
	VPItemID   uint       `json:"SP_item_id"`
	stock      int        `json:"stock"`
	keep       int        `json:"keep"`
	VPIIDetail VPIIDetail `json:"VPII_detail"`
}

func NewVPIInventory(jsonReq string) ([]VPIInventory, error) {
	// contains slice of subdistrict id
	parsedReq := []int{}

	err := json.NewDecoder(strings.NewReader(jsonReq)).Decode(&parsedReq)

	if err != nil {
		return nil, err
	}

	VPIInventories := make([]VPIInventory, len(parsedReq))
	for i, subdistrictID := range parsedReq {
		VPIInventories[i] = VPIInventory{
			VPIIDetail: VPIIDetail{
				SubdistrictID: subdistrictID,
			},
		}
	}

	return VPIInventories, nil
}
