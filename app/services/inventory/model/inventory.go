package model

type InventoryResponse struct {
	SingleProduct  *SPInventory `json:"sp_inventory,omitempty"`
	VariantProduct *VPInventory `json:"vp_inventory,omitempty"`
}
