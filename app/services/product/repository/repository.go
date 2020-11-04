package repository

import "gorm.io/gorm"

type ProductRepository struct {
	DB *gorm.DB
}
