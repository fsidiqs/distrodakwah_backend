package producthandler

import (
	"errors"
)

const RetailPriceName = "retail"

var (
	ErrRetailPriceEmpty = errors.New("harga 'retail' is empty")
)

type ProductImage struct {
	ID  int    `gorm:"primaryKey;autoIncrement;not null"`
	URL string `gorm:"type:varchar(255);not null" json:"url"`
}

type ProductImages []ProductImage

type ProductJSONParsed struct {
	BrandID       uint   `json:"brand_id"`
	CategoryID    uint   `json:"category_id"`
	ProductTypeID uint8  `json:"product_type_id"`
	ProductKindID uint8  `json:"product_kind_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	// ProductImages ProductImages
	SingleProductItem   string `json:"single_product_item"`
	VariantProductItems string `json:"variant_product_items"`
	//for variant
	Variants string `json:"variants"`
}

type ItemCreateBasicProduct struct {
	ID             int    `gorm:"primaryKey;autoIncrement;not null"`
	ProductID      int    `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku            string `json:"sku"`
	Price          int    `json:"retail_price"`
	Weight         int    `json:"weight"`
	SubdistrictIDs string `json:"subdistrict_ids"`
	Options        string `json:"options"`
}

type ItemInventoryRequestCreateBasicProduct struct {
	ID            int `gorm:"primaryKey;autoIncrement;not null"`
	ItemID        int `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"item_id"`
	SubdistrictID int `json:"subdistrict_id"`
}

type OptionReq struct {
	Option string `json:"option"`
}
