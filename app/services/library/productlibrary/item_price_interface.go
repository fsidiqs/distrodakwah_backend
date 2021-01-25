package productlibrary

type ItemPriceI interface {
	GetByProductIDs([]int) []ItemPrice
}

type ProductPriceI interface {
	GetPricesByProductID([]int)
}

type ItemPrice struct {
	ID            uint
	ProductKindID uint8
	ItemableID    uint
	Name          string
	Value         int
}
