package inventorylibrary

import (
	"encoding/json"
	"strings"
)

type SPIInventory struct {
	stock      int        `json:"stock"`
	keep       int        `json:"keep"`
	SPIIDetail SPIIDetail `json:"SPII_detail"`
}

func NewSPIInventory(jsonReq string) ([]SPIInventory, error) {
	// contains slice of subdistrict id
	parsedReq := []int{}

	err := json.NewDecoder(strings.NewReader(jsonReq)).Decode(&parsedReq)

	if err != nil {
		return nil, err
	}

	SPIInventories := make([]SPIInventory, len(parsedReq))
	for i, subdistrictID := range parsedReq {
		SPIInventories[i] = SPIInventory{
			SPIIDetail: SPIIDetail{
				SubdistrictID: subdistrictID,
			},
		}
	}

	return SPIInventories, nil
}
