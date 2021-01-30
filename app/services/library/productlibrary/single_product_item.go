package productlibrary

import (
	"distrodakwah_backend/app/services/handler/producthandler"
	"distrodakwah_backend/app/services/library/inventorylibrary"
	"encoding/json"
	"strings"
)

type SingleProductItem struct {
	ID             uint                            `gorm:"primaryKey;autoIncrement;not null"`
	SPID           uint                            `gorm:"column:SP_id;type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	SingleProduct  *SingleProduct                  `gorm:"foreignKey:SPID" json:"single_product,omitempty"`
	Weight         int                             `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	Sku            string                          `gorm:"type:varchar(255);not null" json:"sku"`
	SPIPrices      []SPItemPrice                   `gorm:"foreignKey:SPItemID;references:ID" json:"single_product_item_prices"`
	SPIInventories []inventorylibrary.SPIInventory `gorm:"foreignKey:SP_item_id" json:"SPIInventories"`
}

func (SingleProductItem) TableName() string {
	return "SP_items"
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
	item.SPIPrices = []SPItemPrice{
		NewSPItemRetailPrice(itemreq.Price),
	}
	item.SPIInventories = SPIInventories

	return item, nil
}

// func (s SingleProductItem) GetItemPriceableByProductID()
