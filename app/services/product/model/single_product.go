package model

type SingleProduct struct {
	ID                   uint64                  `gorm:"primaryKey;autoIncrement;not null"`
	ProductID            uint64                  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Weight               int                     `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	SingleProductsPrices []*SingleProductsPrices `gorm:"foreignKey:SingleProductID;references:ID" json:"prices"`
}

type SingleProductWithParent struct {
	ID        uint64   `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64   `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Weight    int      `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
}

func (SingleProductWithParent) TableName() string {
	return "single_products"
}
