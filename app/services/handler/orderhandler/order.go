package orderhandler

type OrderReq struct {
	OrderItemReq []OrderItemReq `json:"order_items"`
	// ItemOrigin        uint8          `json:"item_origin"`
	ShippingCompanyID uint   `json:"shipping_company_id"`
	CustomerID        uint64 `json:"customer_id"`
}
