package aux

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"

type InventoryResponse struct {
	ProductInventory *model.ProductInventory `json:"product_inventory,omitempty"`
}
