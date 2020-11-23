package model

import (
	"time"

	"gorm.io/gorm"
)

type UserVendor struct {
	ID            uint32         `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt     time.Time      `json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserID        uint64         `json:"user_id"`
	SubdistrictID int            `gorm:"type:int;not null" json:"subdistrict_id"`
	Address       string         `gorm:"type:text;not null" json:"address"`
	Status        uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"status"`
}

func (UserVendor) TableName() string {
	return "users_vendors"
}
