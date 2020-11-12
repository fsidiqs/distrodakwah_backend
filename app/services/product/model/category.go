package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	SubdepartmentID uint           `gorm:"type:int;UNSIGNED;not null" json:"subdepartment_id"`
	Subdepartment   *Subdepartment `gorm:"foreignKey:SubdepartmentID" json:"subdepartment,omitempty"`
	ParentID        uint           `gorm:"type:int;not null" json:"parent_id"`
	Name            string         `gorm:"type:varchar(255);not null" json:"name"`
	ImageID         uint64         `gorm:"type:bigint" json:"image_id"`
}
