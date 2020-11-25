package class

import (
	"encoding/json"
	"fmt"
	"strings"

	productModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
)

type EditProductReq struct {
	ID            string `json:"id"`
	BrandID       uint   `json:"brand_id"`
	CategoryID    uint   `json:"category_id"`
	ProductTypeID uint8  `json:"product_type_id"`
	ProductKindID uint8  `json:"product_kind_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        uint8  `json:"status"`
	Sku           string `json:"sku"`

	ProductDetail string `json:"product_detail"`
}

func (p EditProductReq) ProductDetailJSONDecoder() ProductDetailJSON {
	var err error
	ProductDetails := []ProductDetail{}
	productDetailJSON := &ProductDetailJSON{}

	err = json.NewDecoder(strings.NewReader(p.ProductDetail)).Decode(&productDetailJSON)
	err = json.NewDecoder(strings.NewReader(productDetailJSON.Details)).Decode(&ProductDetails)
	// json.NewDecoder(strings.NewReader(productDetailJSON.Options)).Decode(&Options)
	// json.NewDecoder(strings.NewReader(productDetailJSON.Variants)).Decode(&Variants)
	if p.ProductKindID == productModel.ProductKindVariant
	if err != nil {
		fmt.Printf("error decodeing %+v", err)
	}
	fmt.Printf("decoded %+v", ProductDetails)

	return *productDetailJSON

}

type ProductDetailReader interface {
	ProductDetail() ProductDetail
}
