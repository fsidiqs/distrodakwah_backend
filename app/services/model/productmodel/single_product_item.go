package productmodel

type SingleProductItem struct {
	ID        uint     `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint     `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Weight    int      `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	Sku       string   `gorm:"type:varchar(255);not null" json:"sku"`
}

func (SingleProductItem) TableName() string {
	return "SP_items"
}
