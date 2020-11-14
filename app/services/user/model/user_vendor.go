package model

import (
	"time"

	"gorm.io/gorm"
)

type UserVendor struct {
	ID          int32          `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Province    string         `gorm:"type:varchar(255);not null" json:"province"`
	City        string         `gorm:"type:varchar(255);unique;not null" json:"city"`
	Subdistrict string         `gorm:"type:varchar(255);not null" json:"subdistrict"`
	Address     string         `gorm:"type:varchar(255);not null" json:"address"`
	Status      uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"status"`
}

func (UserVendor) TableName() string {
	return "users_vendors"
}
