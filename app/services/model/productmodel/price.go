package productmodel

type SingleProductPriceTemplate struct {
	SingleProductID int     `json:"single_product_id"`
	Sku             string  `json:"sku"`
	PriceName       string  `json:"nama_harga"`
	PriceValue      float64 `json:"nilai_harga"`
}

type VariantProductPriceTemplate struct {
	VariantProductID int     `json:"variant_product_id"`
	Sku              string  `json:"sku"`
	PriceName        string  `json:"nama_harga"`
	PriceValue       float64 `json:"nilai_harga"`
}
