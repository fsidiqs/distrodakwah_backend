package model

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"

type InventoryResponse struct {
	SingleProduct  *model.SingleProductStock  `json:"sp_inventory,omitempty"`
	VariantProduct *model.VariantProductStock `json:"vp_inventory,omitempty"`
}

type InventoryExportData struct {
	SingleProducts  []*model.SingleProductStock
	VariantProducts []*model.VariantProductStock
}
