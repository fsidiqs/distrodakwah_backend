package inventorylibrary

type ItemInventoryXlsx struct {
	ID    uint64
	Stock int `gorm:"type:INT;NOT NULL" json:"stock"`
	Keep  int `gorm:"type:INT;NOT NULL" json:"keep"`
}

func performStockAdjustment(itemArrXlsx []ItemInventoryXlsx) error {
	// var err error

	return nil
}
