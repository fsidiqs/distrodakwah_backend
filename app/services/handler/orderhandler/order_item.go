package orderhandler

type OrderItemReq struct {
	ItemID uint64 `json:"item_id"`
	Qty    int    `json:"qty"`
}
