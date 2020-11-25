package class

type Options []Option

type Option struct {
	ID               uint64 `json:"id"`
	VariantID        uint64 `json:"variant_id"`
	VariantProductID uint64 `json:"variant_product_id"`
	Name             string `json:"name"`
}

type Variant struct {
	ID        uint64 `json:"id"`
	ProductID uint64 `json:"product_id"`
	Name      string `json:"name"`
}
