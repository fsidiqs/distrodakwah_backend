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
	VariantProductOptions []VPOption                      `gorm:"-" json:"variant_product_options"`
	Weight                int                             `json:"weight"`
	Sku                   string                          `json:"sku"`
	VPIInventories        []inventorylibrary.VPIInventory `gorm:"-" json:"VPIInventory"`
	VPItemPrices          []VPItemPrice                   `gorm:"-"`
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
		vpOptions := make([]VPOption, len(options))
		for i, option := range options {
			vpOptions[i] = VPOption{
				Name: option,
			}
		}
		// options
		VPIInventories, err := inventorylibrary.NewVPIInventory(itemReq.SubdistrictIDs)
		variantProductItems[i] = VariantProductItem{
			Weight:                itemReq.Weight,
			Sku:                   itemReq.Sku,
			VariantProductOptions: vpOptions,
			VPIInventories:        VPIInventories,
		}

	}

	return variantProductItems, nil

}
