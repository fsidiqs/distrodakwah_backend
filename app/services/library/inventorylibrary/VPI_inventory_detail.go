package inventorylibrary

import "distrodakwah_backend/app/services/model/shippingmodel"

type VPIInventoryDetail struct {
	ID                uint                       `gorm:"primaryKey;autoIncrement;not null"`
	VPItemInventoryID uint                       `gorm:"column:VPI_inventory_id;type:BIGINT;UNSIGNED;NOT NULL" json:"vpi_inventory_id"`
	SubdistrictID     int                        `gorm:"type:INT;UNSIGNED;NOT NULL" json:"subdistrict_id"`
	Subdistrict       *shippingmodel.Subdistrict `gorm:"foreignKey:SubdistrictID;references:ID" json:"subdistrict"`
}

func (VPIInventoryDetail) TableName() string {
	return "VPI_inventory_details"
}
