package inventorylibrary

type SPIIDetail struct {
	ID             uint `json:"id"`
	SPIInventoryID uint `json:"SPI_inventory_id"`
	SubdistrictID  int  `json:"subdistrict_id"`
}
