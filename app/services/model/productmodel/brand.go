package productmodel

import (
	"distrodakwah_backend/app/services/model/usermodel"

	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	Name         string               `gorm:"type:varchar(255);not null" json:"name"`
	UserVendorID uint32               `gorm:"type:INT;UNSIGNED;not null" json:"user_vendor_id"`
	UserVendor   usermodel.UserVendor `gorm:"foreignKey:UserVendorID;references:ID" json:"user_vendor"`
	ImageID      uint64               `gorm:"type:BIGINT;UNSIGNED;not null" json:"image_id"`
}
