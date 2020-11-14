package model

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/user/model"

type VPInventoryDetail struct {
	VPInventoryID uint64            `gorm:"type:BIGINT;UNSIGNEND;NOT NULL" json:"vp_inventory_id"`
	VendorID      int32             `gorm:"type:INT;UNSIGNED;NOT NULL" json:"vendor_id"`
	UserVendor    *model.UserVendor `gorm:"foreignKey:VendorID" json:"vendor"`
}

func (VPInventoryDetail) TableName() string {
	return "vp_inventory_details"
}
