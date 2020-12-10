package productmodel

import "distrodakwah_backend/app/services/model/inventorymodel"

type Item struct {
	ID            uint64                        `gorm:"primaryKey;autoIncrement;not null"`
	ProductID     uint64                        `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Product       *Product                      `gorm:"foreignKey:ProductID" json:"product"`
	ItemInventory *inventorymodel.ItemInventory `gorm:"foreignKey:ItemID" json:"item"`
	Weight        int                           `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	Sku           string                        `gorm:"type:varchar(255);not null" json:"sku"`
	Prices        []ItemPrice                   `gorm:"foreignKey:ItemID" json:"prices"`
	Options       []Option                      `gorm:"foreignKey:ItemID" json:"options"`
}

type ItemForPriceExport struct {
	ID        uint64      `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64      `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku       string      `gorm:"type:varchar(255);not null" json:"sku"`
	Prices    []ItemPrice `gorm:"foreignKey:ItemID" json:"prices"`
}

func (ItemForPriceExport) TableName() string {
	return "items"
}
