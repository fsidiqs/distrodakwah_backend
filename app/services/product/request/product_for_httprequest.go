package request

import "errors"

const HargaJualName = "Harga Jual"

var (
	ErrHargaJualEmpty = errors.New("HargaJual is empty")
)

type SingleProductDetailReq struct {
	VendorID uint32  `json:"vendor_id"`
	Sku      string  `json:"sku"`
	Price    float64 `json:"harga_jual"`
	Weight   int     `json:"weight"`
}

func (p *SingleProductDetailReq) Validate() error {
	if p.Price == 0 {
		return ErrHargaJualEmpty
	}
	return nil
}

type VariantProductDetailReq struct {
	VendorID     uint32                      `json:"vendor_id"`
	Sku          string                      `json:"sku"`
	SellingPrice float64                     `json:"harga_jual"`
	Weight       int                         `json:"weight"`
	Variants     []*VariantProductVariantReq `json:"variants"`
}

type VariantProductVariantReq struct {
	VariantValue string `json:"variant_value"`
	OptionValue  string `json:"option_value"`
}
