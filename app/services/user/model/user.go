package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password"`
	RoleID    uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"role_id"`
}

type PublicUser struct {
	Email string
}

type LoginCredetials struct {
	Email    string
	Password string
}
