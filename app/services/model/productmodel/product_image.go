package productmodel

import (
	"errors"
)

type ProductImage struct {
	ID        int    `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint   `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	URL       string `gorm:"type:varchar(255);not null" json:"url"`
}

// ProductImageURL used in basic create product
type ProductImageURL struct {
	URL string `gorm:"type:varchar(255);not null" json:"url"`
}

var (
	ErrProductImageEmptyURL = errors.New("ProductImage is empty")
)

type ProductImages []ProductImage

func (i *ProductImages) Validate() error {
	for _, image := range *i {
		if image.URL == "" {
			return ErrProductImageEmptyURL
		}
	}
	return nil
}
