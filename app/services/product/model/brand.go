package model

import (
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null" json:"name"`
	UserVendorID uint32 `gorm:"type:INT;UNSIGNED;not null" json:"vendor_id"`
	ImageID      uint64 `gorm:"type:BIGINT;UNSIGNED;not null" json:"image_id"`
}
