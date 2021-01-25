package productlibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/handler/producthandler"
	"distrodakwah_backend/app/services/model/productmodel"

	"gorm.io/gorm"
)

type ProductI interface {
	SetProductImages([]string)
	ProductSaveI
}

type ProductSaveI interface {
	SaveProduct() error
}

type Product struct {
	ID                     uint                 `json:"id"`
	CreatedAt              time.Time            `json:"created_at"`
	UpdatedAt              time.Time            `json:"updated_at"`
	DeletedAt              sql.NullTime         `json:"deleted_at"`
	BrandID                uint                 `json:"brand_id"`
	CategoryID             uint                 `json:"category_id"`
	ProductTypeID          uint8                `json:"product_type_id"`
	ProductKindID          uint8                `json:"product_kind_id"`
	Name                   string               `json:"name"`
	Description            string               `json:"description"`
	Status                 string               `json:"status"`
	ProductImages          []ProductImageURL    `gorm:"-" json:"product_images"`
	SingleProductItem      *SingleProductItem   `gorm:"-" json:"single_product_item"`
	VariantProductItems    []VariantProductItem `gorm:"-" json:"variant_product_item"`
	VariantProductVariants []VPVariant          `gorm:"-" json:"variant_products"`
}

func (p Product) DBtoSingleProduct() SingleProduct {
	return SingleProduct{
		ID:            p.ID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
		DeletedAt:     p.DeletedAt,
		BrandID:       p.BrandID,
		CategoryID:    p.CategoryID,
		ProductTypeID: p.ProductTypeID,
		ProductKindID: p.ProductKindID,
		Name:          p.Name,
		Description:   p.Description,
		Status:        p.Status,
	}
}

func (p Product) DBtoVariantProduct() VariantProduct {
	return VariantProduct{
		ID:            p.ID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
		DeletedAt:     p.DeletedAt,
		BrandID:       p.BrandID,
		CategoryID:    p.CategoryID,
		ProductTypeID: p.ProductTypeID,
		ProductKindID: p.ProductKindID,
		Name:          p.Name,
		Description:   p.Description,
		Status:        p.Status,
	}
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

// func GetProductPrices(productids []int) (ItemPriceI, error) {
// 	var err error
// 	var DB *gorm.DB = database.DB
// 	var itemPriceable ItemPriceI

// 	productDBs := []productmodel.Product{}
// 	err = DB.Model(&productmodel.Product{}).Where("id IN (?)", productids).
// 		Find(&productDBs).Error

// 	if err != nil {
// 		fmt.Println("error fetchign Products")
// 		return nil, err
// 	}
// 	productables := make([]ProductI, len(productDBs))
// 	for i, productdb := range productDBs {
// 		if productdb.ProductKindID == productmodel.ProductKindSingle {
// 			productables[i] = &SingleProduct{
// 				ID: productdb.ID,
// 			}
// 		}
// 		//productable.getPricesByProductID()
// 	}
// }

func GetAllProducts() ([]ProductI, error) {
	var err error
	var DB *gorm.DB = database.DB

	productDBs := []Product{}
	err = DB.Model(&productmodel.Product{}).Find(&productDBs).Error
	if err != nil {
		fmt.Println("error fetching all products")
		return nil, err
	}
	productables := make([]ProductI, len(productDBs))
	// singleProducts := []SingleProduct{}
	// variantProducts := []VariantProduct{}

	for i, productdb := range productDBs {
		if productdb.ProductKindID == productmodel.ProductKindSingle {
			sp := productdb.DBtoSingleProduct()
			sp.FetchProductable()
			productables[i] = &sp
		} else if productdb.ProductKindID == productmodel.ProductKindVariant {
			vp := productdb.DBtoVariantProduct()
			vp.FetchProductable()
			productables[i] = &vp
		}
	}

	return productables, nil
}
