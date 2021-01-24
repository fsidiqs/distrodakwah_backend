package inventorylibrary

type ItemInventoryXlsx struct {
	ID    int
	Stock int `gorm:"type:INT;NOT NULL" json:"stock"`
	Keep  int `gorm:"type:INT;NOT NULL" json:"keep"`
}

func performStockAdjustment(itemArrXlsx []ItemInventoryXlsx) error {
	// var err error

	return nil
}
