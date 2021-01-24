package inventorymodel

import "distrodakwah_backend/app/services/model/shippingmodel"

type VPIInventoryDetail struct {
	ID                int                        `gorm:"primaryKey;autoIncrement;not null"`
	VPItemInventoryID int                        `gorm:"column:VPI_inventory_id;type:BIGINT;UNSIGNED;NOT NULL" json:"item_inventory_id"`
	SubdistrictID     int                        `gorm:"type:INT;UNSIGNED;NOT NULL" json:"subdistrict_id"`
	Subdistrict       *shippingmodel.Subdistrict `gorm:"foreignKey:SubdistrictID;references:ID" json:"subdistrict"`
}

func (VPIInventoryDetail) TableName() string {
	return "VPI_inventory_details"
}
