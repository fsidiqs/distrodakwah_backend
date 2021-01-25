package productlibrary

import (
	"database/sql"
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/handler/producthandler"
	"distrodakwah_backend/app/services/model/inventorymodel"
	"distrodakwah_backend/app/services/model/productmodel"
	"fmt"

	"time"

	"gorm.io/gorm"
)

type SingleProductForCreate struct {
	ID                uint               `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	DeletedAt         sql.NullTime       `json:"deleted_at"`
	BrandID           uint               `json:"brand_id"`
	CategoryID        uint               `json:"category_id"`
	ProductTypeID     uint8              `json:"product_type_id"`
	ProductKindID     uint8              `json:"product_kind_id"`
	Name              string             `json:"name"`
	Description       string             `json:"description"`
	Status            string             `json:"status"`
	ProductImages     []string           `json:"product_images"`
	SingleProductItem *SingleProductItem `json:"single_product_item"`
}

type SingleProduct struct {
	ID                uint                        `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt         time.Time                   `json:"created_at"`
	UpdatedAt         time.Time                   `json:"updated_at"`
	DeletedAt         sql.NullTime                `json:"deleted_at"`
	BrandID           uint                        `json:"brand_id"`
	CategoryID        uint                        `json:"category_id"`
	ProductTypeID     uint8                       `json:"product_type_id"`
	ProductKindID     uint8                       `json:"product_kind_id"`
	Name              string                      `json:"name"`
	Description       string                      `json:"description"`
	Status            string                      `json:"status"`
	ProductImages     []productmodel.ProductImage `gorm:"-" json:"product_images"`
	SingleProductItem *SingleProductItem          `json:"single_product_item"`
}

func (p SingleProduct) SaveProduct() error {
	var DB *gorm.DB = database.DB
	var err error
	tx := DB.Begin()
	productImages := []productmodel.ProductImage{}
	for _, productimage := range p.ProductImages {
		productImages = append(productImages, productmodel.ProductImage{
			URL: productimage.URL,
		})
	}

	var productModel *productmodel.Product
	productModel = &productmodel.Product{
		ProductImages: productImages,
		BrandID:       p.BrandID,
		CategoryID:    p.CategoryID,
		ProductKindID: p.ProductKindID,
		Status:        p.Status,
		Name:          p.Name,
		Description:   p.Description,
	}
	err = tx.Model(&productmodel.Product{}).Create(&productModel).Error

	if err != nil {
		fmt.Printf("error creating product \n %+v \n", err)
		tx.Rollback()
		return nil
	}

	singleProductItem := &productmodel.SingleProductItem{
		ProductID: productModel.ID,
		Weight:    p.SingleProductItem.Weight,
		Sku:       p.SingleProductItem.Sku,
	}
	err = tx.Model(&productmodel.SingleProductItem{}).Create(&singleProductItem).Error

	if err != nil {
		fmt.Printf("error creating item \n %+v \n", err)
		tx.Rollback()
		return nil
	}

	SPItemPrice := &productmodel.SPItemPrice{
		SPItemID: singleProductItem.ID,
		Name:     "retail price",
		Value:    0,
	}

	err = tx.Model(&productmodel.SPItemPrice{}).Create(&SPItemPrice).Error

	if err != nil {
		fmt.Printf("error creating SPItemPrice \n %+v \n", err)
		tx.Rollback()
		return nil
	}

	for _, SPIInventory := range p.SingleProductItem.SPIInventories {
		singleProductInventory := &inventorymodel.SPIInventory{
			SPItemID: singleProductItem.ID,
			Stock:    0,
			Keep:     0,
		}

		err = tx.Model(&inventorymodel.SPIInventory{}).Create(&singleProductInventory).Error
		if err != nil {
			fmt.Printf("error creating SPIInventory \n %+v \n", err)
			tx.Rollback()
			return nil
		}
		SPIInventoryDetail := &inventorymodel.SPIInventoryDetail{
			SPItemInventoryID: singleProductInventory.ID,
			SubdistrictID:     SPIInventory.SPIIDetail.SubdistrictID,
		}
		err = tx.Model(&inventorymodel.SPIInventoryDetail{}).Create(&SPIInventoryDetail).Error
		if err != nil {
			fmt.Printf("error creating SPIInventoryDetail \n %+v \n", err)
			tx.Rollback()
			return nil
		}

	}
	return tx.Commit().Error
}
func (p *SingleProduct) FetchProductable() error {
	var err error
	var DB *gorm.DB = database.DB

	singleProductItemDB := productmodel.SingleProductItem{}
	err = DB.Model(&productmodel.SingleProductItem{}).
		Where("product_id = ?", p.ID).
		Find(&singleProductItemDB).Error
	if err != nil {
		fmt.Println("error fetching products")
		return nil
	}

	itemPricesDB := []SPItemPrice{}
	err = DB.Model(&productmodel.SPItemPrice{}).
		Joins("INNER JOIN SP_items on SP_items.id = SP_item_prices.SP_item_id").
		Where("SP_items.id = ?", singleProductItemDB.ID).
		Find(&itemPricesDB).Error

	if err != nil {
		fmt.Println("fetching fetchign prices")
		return nil
	}

	p.SingleProductItem = &SingleProductItem{
		ProductID:    p.ID,
		Weight:       singleProductItemDB.Weight,
		Sku:          singleProductItemDB.Sku,
		SPItemPrices: itemPricesDB,
	}
	return nil
}

func (p SingleProduct) GetProductable() *Product {

	return &Product{
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
		// ProductImages:     p.ProductImages,
		SingleProductItem: p.SingleProductItem,
	}
}

func (p *SingleProduct) SetProductImages(urls []string) {
	for _, url := range urls {

		p.ProductImages = append(p.ProductImages, productmodel.ProductImage{
			URL: url,
		})
	}
}

func NewSingleProduct(productReqParsed producthandler.ProductJSONParsed) (*SingleProduct, error) {

	singleProduct := &SingleProduct{
		BrandID:       productReqParsed.BrandID,
		CategoryID:    productReqParsed.CategoryID,
		ProductTypeID: productReqParsed.ProductTypeID,
		ProductKindID: productReqParsed.ProductKindID,
		Name:          productReqParsed.Name,
		Description:   productReqParsed.Description,
		Status:        productReqParsed.Status,
	}

	item, err := NewSingleProductItem(productReqParsed.SingleProductItem)
	if err != nil {
		return nil, err
	}
	singleProduct.SingleProductItem = item
	return singleProduct, nil
}
