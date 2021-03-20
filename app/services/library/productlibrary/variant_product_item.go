package productlibrary

import (
	"distrodakwah_backend/app/services/handler/producthandler"
	"encoding/json"
	"strings"
)

type VariantProductItem struct {
	ID                    uint                   `gorm:"column:id;primaryKey;autoIncrement;not null"`
	VPID                  uint                   `gorm:"column:VP_id" json:"variant_product_id"`
	VariantProduct        *VariantProduct        `gorm:"foreignKey:VPID" json:"variant_product,omitempty"`
	Weight                int                    `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	Sku                   string                 `gorm:"type:varchar(255);not null" json:"sku"`
	VariantProductOptions []VariantProductOption `gorm:"foreignKey:VPItemID;references:ID" json:"variant_product_options"`
	VPIInventories        []VPIInventory         `gorm:"foreignKey:VP_item_id" json:"variant_product_item_inventories"`
	VPItemPrices          []VPItemPrice          `gorm:"foreignKey:VPItemID;references:ID" json:"variant_product_item_prices"`
}

func (VariantProductItem) TableName() string {
	return "VP_items"
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
		vpOptions := make([]VariantProductOption, len(options))
		for i, option := range options {
			vpOptions[i] = VariantProductOption{
				Name: option,
			}
		}
		// options
		VPIInventories, err := NewVPIInventory(itemReq.SubdistrictIDs)
		variantProductItems[i] = VariantProductItem{
			Weight:                itemReq.Weight,
			Sku:                   itemReq.Sku,
			VariantProductOptions: vpOptions,
			VPIInventories:        VPIInventories,
			VPItemPrices: []VPItemPrice{
				VPItemPrice{
					Name:  RetailPriceName,
					Value: itemReq.Price,
				},
			},
		}

	}

	return variantProductItems, nil

}
