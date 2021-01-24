package usermodel

import (
	"time"

	"gorm.io/gorm"
)

type UserVendor struct {
	ID           uint32         `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserID       int            `json:"user_id"`
	LocationType string         `json:"location_type"`
	LocationID   int            `gorm:"type:int;not null" json:"location_id"`
	Address      string         `gorm:"type:text;not null" json:"address"`
	PostalCode   string         `json:"postal_code"`
	Status       string         `gorm:"type:VARCHAR;UNSIGNED;NOT NULL;default:A" json:"status"`
}
