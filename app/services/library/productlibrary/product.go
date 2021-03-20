package productlibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/helper/pagination"
	"distrodakwah_backend/app/services/handler/producthandler"
	"distrodakwah_backend/app/services/library/query/productquery"
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
	ID                     uint                    `json:"id"`
	CreatedAt              time.Time               `json:"created_at"`
	UpdatedAt              time.Time               `json:"updated_at"`
	DeletedAt              sql.NullTime            `json:"deleted_at"`
	BrandID                uint                    `json:"brand_id"`
	CategoryID             uint                    `json:"category_id"`
	ProductTypeID          uint8                   `json:"product_type_id"`
	ProductKindID          uint8                   `json:"product_kind_id"`
	Name                   string                  `json:"name"`
	Description            string                  `json:"description"`
	Status                 string                  `json:"status"`
	ProductImages          []ProductImageURL       `gorm:"-" json:"product_images"`
	SingleProductItem      *SingleProductItem      `gorm:"-" json:"single_product_item"`
	VariantProductItems    []VariantProductItem    `gorm:"-" json:"variant_product_item"`
	VariantProductVariants []VariantProductVariant `gorm:"-" json:"variant_products"`
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

func ConstructProductForAddNewProduct(productReqJson string) (ProductI, error) {
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

func GetProductByID(productID uint, kindID int) (ProductI, error) {
	var err error
	var productable ProductI

	if kindID == productmodel.ProductKindSingle {
		productable, err = GetSingleProductByID(productID)
	} else if kindID == productmodel.ProductKindVariant {
		productable, err = GetVariantProductByID(productID)
	}
	if err != nil {
		fmt.Printf("product error")
		return nil, err
	}
	return productable, nil
}

type AllProduct struct {
	CreatedAt time.Time    `json:"created_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Kind      uint         `json:"string"`
}

type ProductAllKinds struct {
	ID            uint         `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     sql.NullTime `json:"deleted_at"`
	BrandID       uint         `json:"brand_id"`
	CategoryID    uint         `json:"category_id"`
	ProductTypeID uint8        `json:"product_type_id"`
	ProductKindID uint8        `json:"product_kind_id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Status        string       `json:"status"`
	// helper
	Kind uint `json:"string"`
}

func GetAllProductPricesOnly(productIdArr []uint) []ProductDetailItemPrice {
	var err error
	var DB *gorm.DB = database.DB
	allProducts := []AllProduct{}
	err = DB.Raw(productquery.SEL_ALL_PRODUCTS).Scan(&allProducts).Error

	if err != nil {
		fmt.Println("error fetching all products")
		// return nil, err
	}
	// filter to each kind
	spIDs := []uint{}
	vpIDs := []uint{}
	for _, ap := range allProducts {
		if ap.Kind == productmodel.ProductKindSingle {
			spIDs = append(spIDs, ap.ID)
		} else if ap.Kind == productmodel.ProductKindVariant {
			vpIDs = append(vpIDs, ap.ID)
		}
	}

	productPrices := []ProductDetailItemPrice{}
	err = DB.Raw(productquery.SEL_PRODUCT_PRICES_BY_ID, spIDs, vpIDs).Find(&productPrices).Error
	// variantProducts :=
	return productPrices
}

func GetAllProducts(metadata pagination.Metadata) (*pagination.Pagination, error) {
	res := &pagination.Pagination{Metadata: metadata}
	var err error
	var DB *gorm.DB = database.DB

	allProducts := []ProductAllKinds{}
	err = DB.Raw(
		productquery.SEL_ALL_PRODUCTS_ALL_FIELDS,
		metadata.Offset,
		metadata.Limit,
	).Scan(&allProducts).Error

	spIDs := []uint{}
	// singleProducts := []SingleProduct{}
	vpIDs := []uint{}
	// variantProducts := []VariantProduct{}
	for _, product := range allProducts {
		if product.Kind == productmodel.ProductKindSingle {
			spIDs = append(spIDs, product.ID)
			// singleProducts = append(singleProducts, product.ToSingleProduct())
		} else if product.Kind == productmodel.ProductKindVariant {
			vpIDs = append(vpIDs, product.ID)
			// variantProducts = append(variantProducts, product.ToVariantProduct())
		}
	}
	if err != nil {
		fmt.Println("error fetching all products")
		return nil, err
	}

	singleProducts := []*SingleProductResponse{}
	err = DB.
		Raw(productquery.SEL_SP_BY_ID, spIDs).
		Preload("ProductImages").
		Preload("SingleProductItem.SPIPrices").
		// Preload("SingleProductItem.SPIInventories.SPIInventoryDetail").
		Find(&singleProducts).
		Error

	variantProducts := []*VariantProductResponse{}
	err = DB.
		Raw(productquery.SEL_VP_BY_ID, vpIDs).
		Preload("ProductImages").
		Preload("VariantProductItems.VPItemPrices").
		Preload("VariantProductItems.VPIInventories.VPIInventoryDetail").
		Preload("VariantProductItems").
		// Preload("VariantProductVariants.VariantProductOptions").
		Find(&variantProducts).Error

	var productResponseable []interface{}

	for _, sp := range singleProducts {

		productResponseable = append(productResponseable, sp)
	}
	for _, vp := range variantProducts {

		productResponseable = append(productResponseable, vp)
	}
	res.UpdateElements(productResponseable)
	return res, nil
}

// func NewProductPrices()
