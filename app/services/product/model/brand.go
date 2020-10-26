package model

import (
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	ImageID int    `gorm:"type:BIGINT;not null" json:"image_id"`
}
