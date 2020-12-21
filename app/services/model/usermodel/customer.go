package usermodel

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Name           string       `gorm:"type:varchar(255);not null" json:"name"`
	Phone          string       `gorm:"type:varchar(255);not null" json:"phone"`
	Email          string       `gorm:"type:varchar(255);" json:"email"`
	SubdistrictID  int          `gorm:"type:INT;not null" json:"subdistrict_id"`
	Address        string       `gorm:"type:TEXT;not null" json:"address"`
	PostalCode     string       `gorm:"type:varchar(255);not null" json:"postal_code"`
	UserResellerID uint64       `json:"user_reseller_id"`
	UserReseller   UserReseller `gorm:"foreignKey:UserResellerID" json:"user_reseller"`
}
