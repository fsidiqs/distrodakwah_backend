package productlibrary

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/model/productmodel"
	"fmt"

	"gorm.io/gorm"
)

const RetailPriceName = "retail price"

type SPItemPrice struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	SPItemID uint   `gorm:"column:SP_item_id" json:"sp_item_id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Value    int    `json:"value"`
}

func (SPItemPrice) TableName() string {
	return "SP_item_prices"
}

func NewSPItemRetailPrice(price int) SPItemPrice {
	return SPItemPrice{
		Name:  "retail price",
		Value: price,
	}
}

func (itemPrice SPItemPrice) GetByProductIDs(SPItemIDs []int) ([]ItemPrice, error) {
	var err error
	var DB *gorm.DB = database.DB

	itemPricesDB := []productmodel.SPItemPrice{}
	err = DB.Model(&productmodel.SPItemPrice{}).
		Joins("INNER JOIN SP_items on SP_items.id = SP_item_prices.SP_item_id").
		Where("SP_items.product_id IN (?)", SPItemIDs).
		Find(&itemPricesDB).Error

	if err != nil {
		fmt.Println("error fetchign prices")
		return nil, err
	}

	itemPrices := make([]ItemPrice, len(itemPricesDB))
	for i, itemPriceDB := range itemPricesDB {
		itemPrices[i] = ItemPrice{
			ID:            itemPriceDB.ID,
			ItemableID:    itemPriceDB.SPItemID,
			ProductKindID: productmodel.ProductKindSingle,
			Name:          itemPriceDB.Name,
		}
	}
	return itemPrices, nil
}
