package inventorymodel

import (
	"distrodakwah_backend/app/services/model/shippingmodel"
)

type ItemInventoryDetail struct {
	ID              uint64                     `gorm:"primaryKey;autoIncrement;not null"`
	ItemInventoryID uint64                     `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"item_inventory_id"`
	SubdistrictID   int                        `gorm:"type:INT;UNSIGNED;NOT NULL" json:"subdistrict_id"`
	Subdistrict     *shippingmodel.Subdistrict `gorm:"foreignKey:SubdistrictID;references:ID" json:"subdistrict"`
}
