package repository

import "gorm.io/gorm"

type AuthRepository struct {
	DB *gorm.DB
}

const (
	DBCreateFailedErr = "error creating a record"
)
