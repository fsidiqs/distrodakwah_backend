package class

import (
	"encoding/json"
	"fmt"
	"strings"

	productModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
)

type EditProduct struct {
	ID            uint64                   `json:"id"`
	BrandID       uint                     `json:"brand_id"`
	DeletedAt     bool                     `gorm:"index" json:"deleted_at"`
	CategoryID    uint                     `json:"category_id"`
	ProductTypeID uint8                    `json:"product_type_id"`
	ProductKindID uint8                    `json:"product_kind_id"`
	Name          string                   `json:"name"`
	Description   string                   `json:"description"`
	Status        uint8                    `json:"status"`
	Items         []productModel.Item      `json:"items"`
	ItemPrices    []productModel.ItemPrice `json:"item_prices"`
	Variants      []productModel.Variant   `json:"variants"`
	Options       []productModel.Option    `json:"options"`
}

type EditProductReq struct {
	ID            uint64 `json:"id"`
	BrandID       uint   `json:"brand_id"`
	DeletedAt     bool   `gorm:"index" json:"deleted_at"`
	CategoryID    uint   `json:"category_id"`
	ProductTypeID uint8  `json:"product_type_id"`
	ProductKindID uint8  `json:"product_kind_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        uint8  `json:"status"`
	Items         string `json:"items"`
	ItemPrices    string `json:"item_prices"`
	Variants      string `json:"variants"`
	Options       string `json:"options"`
}

func ProductDecoder(productReqJSON string) (*EditProduct, error) {
	var err error
	editProductReq := &EditProductReq{}

	err = json.NewDecoder(strings.NewReader(productReqJSON)).Decode(&editProductReq)
	if err != nil {
		fmt.Printf("error decodeing1 %+v\n", err)
		return nil, err
	}

	// prepare editProduct
	editProduct := &EditProduct{
		ID:            editProductReq.ID,
		BrandID:       editProductReq.BrandID,
		CategoryID:    editProductReq.CategoryID,
		ProductTypeID: editProductReq.ProductTypeID,
		ProductKindID: editProductReq.ProductKindID,
		Name:          editProductReq.Name,
		Description:   editProductReq.Description,
		Status:        editProductReq.Status,
		Items:         []productModel.Item{},
		Variants:      []productModel.Variant{},
		Options:       []productModel.Option{},
		ItemPrices:    []productModel.ItemPrice{},
	}

	err = json.NewDecoder(strings.NewReader(editProductReq.Items)).Decode(&editProduct.Items)
	err = json.NewDecoder(strings.NewReader(editProductReq.ItemPrices)).Decode(&editProduct.ItemPrices)

	if err != nil {
		fmt.Printf("error decodeing items %+v\n", err)

		return nil, err
	}

	if editProductReq.ProductKindID == productModel.ProductKindVariant {
		err = json.NewDecoder(strings.NewReader(editProductReq.Variants)).Decode(&editProduct.Variants)
		err = json.NewDecoder(strings.NewReader(editProductReq.Options)).Decode(&editProduct.Options)

		if err != nil {
			return nil, err
		}
	}

	return editProduct, nil

}

// type ProductDetailReader interface {
// 	ProductDetail() ProductDetail
// }
