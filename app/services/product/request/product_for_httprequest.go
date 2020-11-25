package request

import "errors"

const HargaJualName = "jual"

var (
	ErrHargaJualEmpty = errors.New("harga 'jual' is empty")
)

type SingleProductDetailReq struct {
	Price  float64 `json:"harga_jual"`
	Weight int     `json:"weight"`
}

func (p *SingleProductDetailReq) Validate() error {
	if p.Price == 0 {
		return ErrHargaJualEmpty
	}
	return nil
}

type VariantProductDetailReq struct {
	Sku          string                      `json:"sku"`
	SellingPrice float64                     `json:"harga_jual"`
	Weight       int                         `json:"weight"`
	Variants     []*VariantProductVariantReq `json:"variants"`
}

type VariantProductVariantReq struct {
	VariantValue string `json:"variant_value"`
	OptionValue  string `json:"option_value"`
}
