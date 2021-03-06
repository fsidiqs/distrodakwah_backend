package productlibrary

import (
	"distrodakwah_backend/app/services/handler/producthandler"
	"encoding/json"
	"strings"
)

type SingleProductItem struct {
	ID             uint           `gorm:"primaryKey;autoIncrement;not null"`
	SPID           uint           `gorm:"column:SP_id" json:"single_product_id"`
	SingleProduct  *SingleProduct `gorm:"foreignKey:SPID" json:"single_product,omitempty"`
	Weight         int            `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	Sku            string         `gorm:"type:varchar(255);not null" json:"sku"`
	SPIPrices      []SPItemPrice  `gorm:"foreignKey:SPItemID;references:ID" json:"single_product_item_prices"`
	SPIInventories []SPIInventory `gorm:"foreignKey:SPItemID" json:"SPIInventories"`
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

	SPIInventories, err := NewSPIInventory(itemreq.SubdistrictIDs)
	item.SPIPrices = []SPItemPrice{
		NewSPItemRetailPrice(itemreq.Price),
	}
	item.SPIInventories = SPIInventories

	return item, nil
}

// func (s SingleProductItem) GetItemPriceableByProductID()
