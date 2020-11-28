package class

import (
	"encoding/json"
	"fmt"
	"strings"

	productModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
)

type EditProduct struct {
	ID            uint64 `json:"id"`
	BrandID       uint   `json:"brand_id"`
	CategoryID    uint   `json:"category_id"`
	ProductTypeID uint8  `json:"product_type_id"`
	ProductKindID uint8  `json:"product_kind_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        uint8  `json:"status"`
	Sku           string `json:"sku"`

	ProductDetail ProductDetail `json:"product_detail"`
}

type EditProductReq struct {
	ID            uint64 `json:"id"`
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

func ProductDetailJSONDecoder(productReqJSON string) (*EditProduct, error) {
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
		Sku:           editProductReq.Sku,
		ProductDetail: ProductDetail{},
	}

	// singleProductDetail := ProductDetail{}

	// variantProductDetails :=
	// variantProducts := []VariantProduct{}
	productDetailsJSON := &ProductDetailsJSON{}

	err = json.NewDecoder(strings.NewReader(editProductReq.ProductDetail)).Decode(&productDetailsJSON)
	if err != nil {
		fmt.Printf("error decodeing2 %+v\n", err)

		return nil, err
	}

	// // json.NewDecoder(strings.NewReader(productDetailJSON.Options)).Decode(&Options)
	// // json.NewDecoder(strings.NewReader(productDetailJSON.Variants)).Decode(&Variants)
	if editProductReq.ProductKindID == productModel.ProductKindVariant {
		err = json.NewDecoder(strings.NewReader(productDetailsJSON.Details)).Decode(&editProduct.ProductDetail.VariantProductDetail.VariantProductArr)
		err = json.NewDecoder(strings.NewReader(productDetailsJSON.Options)).Decode(&editProduct.ProductDetail.VariantProductDetail.Options)
		err = json.NewDecoder(strings.NewReader(productDetailsJSON.Variants)).Decode(&editProduct.ProductDetail.VariantProductDetail.Variants)
		if err != nil {
			return nil, err

		}
	} else if editProductReq.ProductKindID == productModel.ProductKindSingle {
		// err = json.NewDecoder(strings.NewReader(productDetailsJSON.Details)).Decode(&singleProductDetail)
	}
	if err != nil {
		fmt.Printf("error decodeing %+v", err)
		return nil, err

	}
	// fmt.Printf("edit product %+v\n\n", editProduct)

	fmt.Printf("edit product detail options %+v\n", editProduct.ProductDetail.VariantProductDetail.Variants)

	return editProduct, nil

}

// type ProductDetailReader interface {
// 	ProductDetail() ProductDetail
// }
