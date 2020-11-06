package aux

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"

type ProductPriceTemplate struct {
	SingleProductPricesTemplate []*model.SingleProductPriceTemplate  `json:"single_product_prices"`
	VariantProductPriceTemplate []*model.VariantProductPriceTemplate `json:"variant_product_prices"`
}
