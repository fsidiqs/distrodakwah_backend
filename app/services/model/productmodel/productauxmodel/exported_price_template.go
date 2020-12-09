package productauxmodel

type ItemPriceExport struct {
	ID      uint64  `json:"id"`
	ItemID  uint64  `json:"item_id"`
	ItemSku string  `json:"sku"`
	Name    string  `json:"name"`
	Value   float64 `json:"value"`
}

type ItemPriceArrExport []ItemPriceExport
