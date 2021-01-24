package productlibrary

type Options []Option

type Option struct {
	ID        uint   `json:"id"`
	VariantID uint   `json:"variant_id"`
	ItemID    uint   `json:"variant_product_id"`
	Name      string `json:"name"`
}
