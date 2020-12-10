package productlibrary

type ProductDetailsJSON struct {
	Details string `json:"details"`
	// for variants
	Options  string `json:"options"`
	Variants string `json:"variants"`
}

// type ProductDetail struct {
// 	ID        uint64
// 	ProductID uint64 `json:"product_id"`
// 	Sku       string
// 	Weight    int
// }

type ProductDetail struct {
	SingleProductDetail  string
	VariantProductDetail VariantProductDetailReq
}
