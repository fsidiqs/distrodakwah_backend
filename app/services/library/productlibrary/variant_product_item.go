package productlibrary

import (
	"distrodakwah_backend/app/services/handler/producthandler"
	"distrodakwah_backend/app/services/library/inventorylibrary"
	"encoding/json"
	"strings"
)

type VariantProductItem struct {
	ID                    uint                            `json:"id"`
	ProductID             uint                            `json:"product_id"`
	VariantProductOptions []string                        `json:"variant_product_options"`
	Weight                int                             `json:"weight"`
	Sku                   string                          `json:"sku"`
	VPIInventories        []inventorylibrary.VPIInventory `json:"VPIInventory"`
}

func NewVariantProductItem(itemReqJson string) ([]VariantProductItem, error) {
	//item
	itemReqs := []producthandler.ItemCreateBasicProduct{}

	err := json.NewDecoder(strings.NewReader(itemReqJson)).Decode(&itemReqs)

	if err != nil {
		return nil, err
	}
	variantProductItems := make([]VariantProductItem, len(itemReqs))

	for i, itemReq := range itemReqs {
		// options
		options := []string{}
		err := json.NewDecoder(strings.NewReader(itemReq.Options)).Decode(&options)
		if err != nil {
			return nil, err
		}

		// options
		VPIInventories, err := inventorylibrary.NewVPIInventory(itemReq.SubdistrictIDs)
		variantProductItems[i] = VariantProductItem{
			Weight:                itemReq.Weight,
			Sku:                   itemReq.Sku,
			VariantProductOptions: options,
			VPIInventories:        VPIInventories,
		}

	}

	return variantProductItems, nil

}
