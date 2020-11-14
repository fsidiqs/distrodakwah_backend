package model

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/user/model"

type SPInventoryDetail struct {
	SPInventoryID uint64            `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"sp_inventory_id"`
	VendorID      int32             `gorm:"type:INT;UNSIGNED;NOT NULL" json:"vendor_id"`
	UserVendor    *model.UserVendor `gorm:"foreignKey:VendorID" json:"vendor"`
}

func (SPInventoryDetail) TableName() string {
	return "sp_inventory_details"
}
