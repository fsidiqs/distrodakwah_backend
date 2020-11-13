package model

type VPInventoryDetail struct {
	VPInventoryID uint64 `gorm:"type:BIGINT;UNSIGNEND;NOT NULL" json:"vp_inventory_id"`
	VendorID      int32  `gorm:"type:INT;UNSIGNED;NOT NULL" json:"vendor_id"`
}

func (VPInventoryDetail) TableName() string {
	return "vp_inventory_details"
}
