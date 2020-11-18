package aux

type ExcelStockFormat struct {
	ProductKindID    uint8  `json:"product_kind_id"`
	RelatedProductID uint64 `json:"related_product_id"`
	Stock            int    `json:"stock"`
	Keep             int    `json:"keep"`
}
