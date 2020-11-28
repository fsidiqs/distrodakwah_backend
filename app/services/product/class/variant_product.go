package class

import (
	productModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
)

type ProductDetailVariant struct {
	Details  []VariantProduct
	Options  []Option
	Variants []Variant
}

type VariantProduct struct {
	ID        uint64
	ProductID uint64 `json:"product_id"`
	Sku       string
	Weight    int
}

// func (vp ProductDetailVariant) ProductDetail() ProductDetail {
// 	return ProductDetail{
// 		// ID:              vp.ID,
// 		// ProductDetailID: vp.ProductID,
// 		// Weight:          vp.Weight,
// 		// Sku:             vp.Sku,
// 	}
// }

type VariantProductArr []VariantProduct

type VariantProductDetailReq struct {
	VariantProductArr VariantProductArr
	Options           []productModel.Option
	Variants          []productModel.Variant
}
