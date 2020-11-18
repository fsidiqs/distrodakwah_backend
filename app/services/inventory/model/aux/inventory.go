package aux

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"

type InventoryResponse struct {
	SPInventory *model.SPInventory `json:"sp_inventory,omitempty"`
	VPInventory *model.VPInventory `json:"vp_inventory,omitempty"`
}
