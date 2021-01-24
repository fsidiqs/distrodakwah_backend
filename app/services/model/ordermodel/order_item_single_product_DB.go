package ordermodel

// type OrderItemSingleProductDB struct {
// 	ID              int                `gorm:"primaryKey;autoIncrement;not null"`
// 	OrderID         int                `json:"order_id"`
// 	SingleProductID int                `json:"single_product_id"`
// 	SPInventory     *invModel.SPInventory `gorm:"foreignKey:SingleProductID" json:"sp_inventory"`
// 	Qty             int                   `json:"qty"`
// 	UnitWeight      int                   `json:"unit_weight"`
// 	// DropshipperItemPrice float64               `gorm:"type:decimal(19,2);not null;default:0.0" json:"dropshipper_item_price"`
// 	// RetailItemPrice      float64               `gorm:"type:decimal(19,2);not null;default:0.0" json:"retail_item_price"`
// 	Prices          *OrderItemSPPriceArr `gorm:"-" json:"prices"`
// 	OrderShippingID int               `json:"order_shipping_id"`
// }
