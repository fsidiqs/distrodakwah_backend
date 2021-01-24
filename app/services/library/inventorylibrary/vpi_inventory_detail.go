package inventorylibrary

type VPIIDetail struct {
	ID             uint `json:"id"`
	VPIInventoryID uint `json:"VPI_inventory_id"`
	SubdistrictID  int  `json:"subdistrict_id"`
}
