package productmodel

type VariantProductItem struct {
	ID                    uint                   `gorm:"primaryKey;autoIncrement;not null"`
	VPID                  uint                   `gorm:"column:VP_id;type:BIGINT;UNSIGNED;NOT NULL" json:"variant_product_id"`
	Weight                int                    `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	Sku                   string                 `gorm:"type:varchar(255);not null" json:"sku"`
	VariantProductOptions []VariantProductOption `gorm:"foreignKey:VPItemID;references:ID" json:"variant_product_options"`
	VPItemPrices          []VPItemPrice          `gorm:"foreignKey:VPItemID;references:ID" json:"variant_product_item_prices"`
}

func (VariantProductItem) TableName() string {
	return "VP_items"
}
