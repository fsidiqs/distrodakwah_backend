package productauxmodel

type ItemPriceExport struct {
	ID      int     `json:"id"`
	ItemID  uint    `json:"item_id"`
	ItemSku string  `json:"sku"`
	Name    string  `json:"name"`
	Value   float64 `json:"value"`
}

type ItemPriceArrExport []ItemPriceExport
