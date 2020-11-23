package repository

import "gorm.io/gorm"

type OrderRepository struct {
	DB *gorm.DB
}
