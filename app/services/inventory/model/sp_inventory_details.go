package model

type SPInventoryDetail struct {
	SPInventoryID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"sp_inventory_id"`
	VendorID      int32  `gorm:"type:INT;UNSIGNED;NOT NULL" json:"vendor_id"`
}

func (SPInventoryDetail) TableName() string {
	return "sp_inventory_details"
}
