package productlibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"distrodakwah_backend/app/services/handler/producthandler"
	"distrodakwah_backend/app/services/model/productmodel"
)

type ProductI interface {
	GetProductable() Product
	// GetProductableItems() []Item
	SetProductImages([]string)
	ProductSaveable
}

type ProductSaveable interface {
	SaveProduct() error
}
type Product struct {
	ID                  uint                 `json:"id"`
	CreatedAt           time.Time            `json:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at"`
	DeletedAt           sql.NullTime         `json:"deleted_at"`
	BrandID             uint                 `json:"brand_id"`
	CategoryID          uint                 `json:"category_id"`
	ProductTypeID       uint8                `json:"product_type_id"`
	ProductKindID       uint8                `json:"product_kind_id"`
	Name                string               `json:"name"`
	Description         string               `json:"description"`
	Status              string               `json:"status"`
	ProductImages       []string             `json:"product_images"`
	SingleProductItem   *SingleProductItem   `json:"single_product_item"`
	VariantProductItems []VariantProductItem `json:"variant_product_item"`
	// ItemPrices    []productmodel.ItemPrice `json:"item_prices"`
	// Variants      []productmodel.Variant   `json:"variants"`
	// Options       []productmodel.Option    `json:"options"`
}

type EditProductReq struct {
	ID            uint   `json:"id"`
	BrandID       uint   `json:"brand_id"`
	DeletedAt     bool   `gorm:"index" json:"deleted_at"`
	CategoryID    uint   `json:"category_id"`
	ProductTypeID uint8  `json:"product_type_id"`
	ProductKindID uint8  `json:"product_kind_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	Items         string `json:"items"`
	ItemPrices    string `json:"item_prices"`
	Variants      string `json:"variants"`
	Options       string `json:"options"`
}

func ProductDecoder(productReqJSON string) (*Product, error) {
	var err error
	editProductReq := &EditProductReq{}

	err = json.NewDecoder(strings.NewReader(productReqJSON)).Decode(&editProductReq)
	if err != nil {
		fmt.Printf("error decodeing1 %+v\n", err)
		return nil, err
	}

	// prepare editProduct
	editProduct := &Product{
		ID:            editProductReq.ID,
		BrandID:       editProductReq.BrandID,
		CategoryID:    editProductReq.CategoryID,
		ProductTypeID: editProductReq.ProductTypeID,
		ProductKindID: editProductReq.ProductKindID,
		Name:          editProductReq.Name,
		Description:   editProductReq.Description,
		Status:        editProductReq.Status,
		// Items:         []productmodel.Item{},
		// Variants:      []productmodel.Variant{},
		// Options:       []productmodel.Option{},
		// ItemPrices:    []productmodel.ItemPrice{},
	}

	// err = json.NewDecoder(strings.NewReader(editProductReq.Items)).Decode(&editProduct.Items)
	// err = json.NewDecoder(strings.NewReader(editProductReq.ItemPrices)).Decode(&editProduct.ItemPrices)

	// if err != nil {
	// 	fmt.Printf("error decodeing items %+v\n", err)

	// 	return nil, err
	// }

	// if editProductReq.ProductKindID == productmodel.ProductKindVariant {
	// 	err = json.NewDecoder(strings.NewReader(editProductReq.Variants)).Decode(&editProduct.Variants)
	// 	err = json.NewDecoder(strings.NewReader(editProductReq.Options)).Decode(&editProduct.Options)

	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return editProduct, nil

}

func ConstructProductModel(productreq *producthandler.ProductJSONParsed) (*productmodel.Product, error) {
	product := &productmodel.Product{
		BrandID:       productreq.BrandID,
		CategoryID:    productreq.CategoryID,
		ProductTypeID: productreq.ProductTypeID,
		ProductKindID: productreq.ProductKindID,
		Status:        productreq.Status,
		Name:          productreq.Name,
		Description:   productreq.Description,
	}
	return product, nil
}

func ConstructProduct(productReqJson string) (ProductI, error) {
	var err error
	var productable ProductI

	productReqParsed := producthandler.ProductJSONParsed{}
	err = json.NewDecoder(strings.NewReader(productReqJson)).Decode(&productReqParsed)
	if err != nil {
		return nil, err
	}

	if productReqParsed.ProductKindID == productmodel.ProductKindSingle {
		productable, err = NewSingleProduct(productReqParsed)
	} else if productReqParsed.ProductKindID == productmodel.ProductKindVariant {
		productable, err = NewVariantProduct(productReqParsed)
	}
	return productable, nil
}
