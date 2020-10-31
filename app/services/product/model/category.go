package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	ParentID uint   `gorm:"type:int;not null" json:"parent_id"`
	ImageID  uint64 `gorm:"type:bigint" json:"image_id"`
}
