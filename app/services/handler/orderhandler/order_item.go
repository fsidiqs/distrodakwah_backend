package orderhandler

type OrderItemReq struct {
	ItemID int `json:"item_id"`
	Qty    int `json:"qty"`
}
