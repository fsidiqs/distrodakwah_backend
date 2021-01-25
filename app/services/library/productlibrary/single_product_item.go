package productlibrary

import (
	"distrodakwah_backend/app/services/handler/producthandler"
	"distrodakwah_backend/app/services/library/inventorylibrary"
	"encoding/json"
	"strings"
)

type SingleProductItem struct {
	ProductID      uint
	Weight         int                             `json:"weight"`
	Sku            string                          `json:"sku"`
	SPIInventories []inventorylibrary.SPIInventory `gorm:"-" json:"SPIInventory"`
	SPItemPrices   []SPItemPrice                   `gorm:"-"`
}

func NewSingleProductItem(itemReqJson string) (*SingleProductItem, error) {
	itemreq := producthandler.ItemCreateBasicProduct{}

	err := json.NewDecoder(strings.NewReader(itemReqJson)).Decode(&itemreq)
	if err != nil {
		return nil, err
	}
	item := &SingleProductItem{
		Weight: itemreq.Weight,
		Sku:    itemreq.Sku,
	}

	SPIInventories, err := inventorylibrary.NewSPIInventory(itemreq.SubdistrictIDs)
	item.SPItemPrices = []SPItemPrice{
		NewSPItemRetailPrice(),
	}
	item.SPIInventories = SPIInventories

	return item, nil
}

// func (s SingleProductItem) GetItemPriceableByProductID()
