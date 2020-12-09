package productrepository

import "gorm.io/gorm"

type ProductRepository struct {
	DB *gorm.DB
}
