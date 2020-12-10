package inventoryrepository

import "gorm.io/gorm"

type InventoryRepository struct {
	DB *gorm.DB
}
