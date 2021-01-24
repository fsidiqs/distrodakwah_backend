package productmodel

import "gorm.io/gorm"

type Vendor struct {
	gorm.Model
	ImageID int    `gorm:"type:BIGINT;UNSIGNED;not null" json:"image_id"`
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
}
