package model

type SingleProduct struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Weight    int    `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
}

type SingleProductWithPrices struct {
	ID                   uint64                 `gorm:"primaryKey;autoIncrement;not null"`
	ProductID            uint64                 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Weight               int                    `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	SingleProductsPrices SingleProductsPriceArr `gorm:"foreignKey:SingleProductID;references:ID" json:"prices"`
}

func (SingleProductWithPrices) TableName() string {
	return "single_products"
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

// type SingleProductStock struct {
// 	ID          uint64             `gorm:"primaryKey;autoIncrement;not null"`
// 	ProductID   uint64             `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
// 	Product     *ProductSimpleInfo `gorm:"foreignKey:ProductID" json:"product,omitempty"`
// 	Weight      int                `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
// 	SPInventory *model.SPInventory `gorm:"foreignKey:SingleProductID;references:ID" json:"inventory"`
// }

// func (SingleProductStock) TableName() string {
// 	return "single_products"
// }
