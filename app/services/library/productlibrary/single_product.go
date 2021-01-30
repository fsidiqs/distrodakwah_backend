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
	ID                uint                 `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt         time.Time            `json:"created_at"`
	UpdatedAt         time.Time            `json:"updated_at"`
	DeletedAt         gorm.DeletedAt       `gorm:"index" json:"deleted_at"`
	BrandID           uint                 `json:"brand_id"`
	CategoryID        uint                 `json:"category_id"`
	ProductTypeID     uint8                `gorm:"type:INT;UNSIGNED;NOT NULL" json:"product_type_id"`
	Name              string               `gorm:"type:varchar(255);not null" json:"name"`
	Description       string               `gorm:"type:text;not null" json:"description"`
	Status            string               `gorm:"type:NOT NULL;default:0" json:"status"`
	ProductImages     []SingleProductImage `gorm:"foreignKey:SPID;references:ID" json:"single_product_images"`
	SingleProductItem *SingleProductItem   `gorm:"foreignKey:SPID;references:ID" json:"single_product_item"`
}

func (p SingleProduct) SaveProduct() error {
	var DB *gorm.DB = database.DB
	var err error
	tx := DB.Begin()
	productImages := []productmodel.SingleProductImage{}
	for _, productimage := range p.ProductImages {
		productImages = append(productImages, productmodel.SingleProductImage{
			URL: productimage.URL,
		})
	}

	var productModel *productmodel.SingleProduct
	productModel = &productmodel.SingleProduct{
		ProductImages: productImages,
		BrandID:       p.BrandID,
		CategoryID:    p.CategoryID,
		Status:        p.Status,
		Name:          p.Name,
		Description:   p.Description,
	}
	err = tx.Model(&productmodel.SingleProduct{}).Create(&productModel).Error

	if err != nil {
		fmt.Printf("error creating product \n %+v \n", err)
		tx.Rollback()
		return nil
	}

	singleProductItem := &productmodel.SingleProductItem{
		SPID:   productModel.ID,
		Weight: p.SingleProductItem.Weight,
		Sku:    p.SingleProductItem.Sku,
	}
	err = tx.Model(&productmodel.SingleProductItem{}).Create(&singleProductItem).Error

	if err != nil {
		fmt.Printf("error creating item \n %+v \n", err)
		tx.Rollback()
		return nil
	}

	SPItemPrice := &productmodel.SPItemPrice{
		SPItemID: singleProductItem.ID,
		Name:     p.SingleProductItem.SPIPrices[0].Name,
		Value:    p.SingleProductItem.SPIPrices[0].Value,
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
			SubdistrictID:     SPIInventory.SPIInventoryDetail.SubdistrictID,
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
		SPID:      p.ID,
		Weight:    singleProductItemDB.Weight,
		Sku:       singleProductItemDB.Sku,
		SPIPrices: itemPricesDB,
	}
	return nil
}

func (p SingleProduct) GetProductable() *Product {

	return &Product{
		ID:            p.ID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
		DeletedAt:     sql.NullTime{p.DeletedAt.Time, p.DeletedAt.Valid},
		BrandID:       p.BrandID,
		CategoryID:    p.CategoryID,
		ProductTypeID: p.ProductTypeID,
		Name:          p.Name,
		Description:   p.Description,
		Status:        p.Status,
		// ProductImages:     p.ProductImages,
		SingleProductItem: p.SingleProductItem,
	}
}

func (p *SingleProduct) SetProductImages(urls []string) {
	for _, url := range urls {

		p.ProductImages = append(p.ProductImages, SingleProductImage{
			URL: url,
		})
	}
}

func NewSingleProduct(productReqParsed producthandler.ProductJSONParsed) (*SingleProduct, error) {

	singleProduct := &SingleProduct{
		BrandID:       productReqParsed.BrandID,
		CategoryID:    productReqParsed.CategoryID,
		ProductTypeID: productReqParsed.ProductTypeID,
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

func GetSingleProductByID(id uint) (*SingleProduct, error) {
	var err error
	var DB *gorm.DB = database.DB
	singleProduct := &SingleProduct{}
	err = DB.Model(&SingleProduct{}).
		Preload("ProductImages").
		Preload("SingleProductItem.SPIPrices").
		Preload("SingleProductItem.SPIInventories.SPIInventoryDetail").
		Where("id = ?", id).
		First(&singleProduct).Error
	if err != nil {
		return nil, err
	}
	return singleProduct, nil
}
