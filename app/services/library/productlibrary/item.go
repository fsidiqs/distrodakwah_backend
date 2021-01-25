package productlibrary

type Item struct {
	ID        uint `json:"id"`
	ProductID uint `json:"product_id"`
	// Product   *Product `json:"product,omitempty"`
	// ItemInventory []inventorymodel.ItemInventory `gorm:"foreignKey:itemable_id" json:"item"`
	Options    []string    `json:"options"`
	Weight     int         `json:"weight"`
	Sku        string      `json:"sku"`
	ItemPrices []ItemPrice `json:"item_prices"`
}
