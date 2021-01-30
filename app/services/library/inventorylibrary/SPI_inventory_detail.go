package inventorylibrary

import "distrodakwah_backend/app/services/model/shippingmodel"

type SPIInventoryDetail struct {
	ID                uint                       `gorm:"primaryKey;autoIncrement;not null"`
	SPItemInventoryID uint                       `gorm:"column:SPI_inventory_id;type:BIGINT;UNSIGNED;NOT NULL" json:"item_inventory_id"`
	SubdistrictID     int                        `gorm:"type:INT;UNSIGNED;NOT NULL" json:"subdistrict_id"`
	Subdistrict       *shippingmodel.Subdistrict `gorm:"foreignKey:SubdistrictID;references:ID" json:"subdistrict"`
}

func (SPIInventoryDetail) TableName() string {
	return "SPI_inventory_details"
}
