package producthandler

import (
	"errors"

	"distrodakwah_backend/app/services/model/productmodel"
)

const RetailPriceName = "retail"

var (
	ErrRetailPriceEmpty = errors.New("harga 'retail' is empty")
)

type ProductFromRequestJSON struct {
	BrandID       uint   `json:"brand_id"`
	CategoryID    uint   `json:"category_id"`
	ProductTypeID uint8  `json:"product_type_id"`
	ProductKindID uint8  `json:"product_kind_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        uint8  `json:"status"`
	ProductImages productmodel.ProductImages
	Items         string `json:"items"`
	Variants      string `json:"variants"`
}

type ItemCreateBasicProduct struct {
	ID        uint64  `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku       string  `json:"sku"`
	Price     float64 `json:"retail_price"`
	Weight    int     `json:"weight"`
	Options   string  `json:"options"`
}

type OptionReq struct {
	Option string `json:"option"`
}
