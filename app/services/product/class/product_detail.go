package class

type EditProductDetailVPReq struct {
	VariantProducts string `json:"variant_products"`
	Options         string `json:"options"`
	Variants        string `json:"variants"`
}

// func (EditProductDetailVPReq) ProductDetailDecoder() ProductDetailReq {
// 	return ProductDetailReq{}
// }

type EditProductDetailSPReq struct {
	SingleProducts string `json:"single_product"`
	// Options         string `json:"options"`
	// Variants        string `json:"variants"`
}

// func (EditProductDetailSPReq) ProductDetailDecoder() ProductDetailReq {
// 	return ProductDetailReq{}

// }

type ProductDetailDecoder interface {
	ProductDetailDecoder() ProductDetailJSON
}

type ProductDetailReq string

type ProductDetailJSON struct {
	Details string `json:"details"`
	// for variants
	Options  string `json:"options"`
	Variants string `json:"variants"`
}

type ProductDetail struct {
	ID        uint64
	ProductID uint64 `json:"product_id"`
	Sku       string
	Weight    int
}
