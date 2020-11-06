package model

type SingleProductPriceTemplate struct {
	SingleProductID uint64  `json:"single_product_id"`
	Sku             string  `json:"sku"`
	PriceName       string  `json:"nama_harga"`
	PriceValue      float64 `json:"nilai_harga"`
}

type VariantProductPriceTemplate struct {
	VariantProductID uint64  `json:"variant_product_id"`
	Sku              string  `json:"sku"`
	PriceName        string  `json:"nama_harga"`
	PriceValue       float64 `json:"nilai_harga"`
}
