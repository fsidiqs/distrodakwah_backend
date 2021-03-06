package productmodel

import (
	"errors"
)

// type ProductImage {}
type SingleProductImage struct {
	ID   int    `gorm:"primaryKey;autoIncrement;not null"`
	SPID uint   `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	URL  string `gorm:"type:varchar(255);not null" json:"url"`
}

type VariantProductImage struct {
	ID   int    `gorm:"primaryKey;autoIncrement;not null"`
	VPID uint   `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"variant_product_id"`
	URL  string `gorm:"type:varchar(255);not null" json:"url"`
}

var (
	ErrProductImageEmptyURL = errors.New("ProductImage is empty")
)
