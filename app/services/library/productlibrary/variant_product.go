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

type VariantProduct struct {
	ID                  uint                        `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt           time.Time                   `json:"created_at"`
	UpdatedAt           time.Time                   `json:"updated_at"`
	DeletedAt           sql.NullTime                `json:"deleted_at"`
	BrandID             uint                        `json:"brand_id"`
	CategoryID          uint                        `json:"category_id"`
	ProductTypeID       uint8                       `json:"product_type_id"`
	ProductKindID       uint8                       `json:"product_kind_id"`
	Name                string                      `json:"name"`
	Description         string                      `json:"description"`
	Status              string                      `json:"status"`
	ProductImages       []productmodel.ProductImage `gorm:"-" json:"product_images"`
	VariantProductItems []VariantProductItem        `json:"variant_product_item"`
	VPVariants          []VPVariant                 `json:"variant_products"`
}

func (p VariantProduct) SaveProduct() error {
	var DB *gorm.DB = database.DB

	var err error
	tx := DB.Begin()

	productModel := &productmodel.Product{
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

	productImagesModel := []productmodel.ProductImage{}
	for _, productImage := range p.ProductImages {
		productImagesModel = append(productImagesModel, productmodel.ProductImage{
			ProductID: productModel.ID,
			URL:       productImage.URL,
		})
	}

	err = tx.Model(&productmodel.ProductImage{}).Create(&productImagesModel).Error

	if err != nil {
		fmt.Printf("error creating image \n %+v \n", err)
		tx.Rollback()
		return nil
	}

	variantProductVariants := []productmodel.VariantProductVariant{}
	for _, variantProduct := range p.VPVariants {
		variantProductVariants = append(variantProductVariants, productmodel.VariantProductVariant{
			ProductID: productModel.ID,
			Name:      variantProduct.Name,
		})
	}
	err = tx.Model(&productmodel.VariantProductVariant{}).Create(&variantProductVariants).Error
	if err != nil {
		fmt.Printf("error creating VariantProductVariant \n %+v \n", err)
		tx.Rollback()
		return nil
	}
	// item variant product items loop
	for _, item := range p.VariantProductItems {
		variantProductItem := productmodel.VariantProductItem{
			ProductID: productModel.ID,
			Weight:    item.Weight,
			Sku:       item.Sku,
		}

		err = tx.Model(&productmodel.VariantProductItem{}).Create(&variantProductItem).Error
		if err != nil {
			fmt.Printf("error creating VariantProductItem \n %+v \n", err)
			tx.Rollback()
			return nil
		}

		vpOptions := []productmodel.VariantProductOption{}
		// assume total option count = total variant count, and the order is same aswell
		for i, vpOption := range item.VariantProductOptions {
			vpOptions = append(vpOptions, productmodel.VariantProductOption{
				VariantProductVariantID: variantProductVariants[i].ID,
				VariantProductItemID:    variantProductItem.ID,
				Name:                    vpOption.Name,
			})
		}
		err = tx.Model(&productmodel.VariantProductOption{}).Create(&vpOptions).Error
		if err != nil {
			fmt.Printf("error creating VariantProductOption \n %+v \n", err)
			tx.Rollback()
			return nil
		}
		// retail price create
		VPItemPrice := &productmodel.VPItemPrice{
			VPItemID: variantProductItem.ID,
			Name:     "retail price",
			Value:    0,
		}

		err = tx.Model(&productmodel.VPItemPrice{}).Create(&VPItemPrice).Error

		if err != nil {
			fmt.Printf("error creating 	VPItemPrice \n %+v \n", err)
			tx.Rollback()
			return nil
		}

		// retail price create
		// inventory create
		for _, VPIInventory := range item.VPIInventories {
			variantProductInventory := &inventorymodel.VPIInventory{
				VPItemID: variantProductItem.ID,
				Stock:    0,
				Keep:     0,
			}
			err = tx.Model(&inventorymodel.VPIInventory{}).Create(&variantProductInventory).Error
			if err != nil {
				fmt.Printf("error creating VPIInventory \n %+v \n", err)
				tx.Rollback()
				return nil
			}
			VPIInventoryDetail := &inventorymodel.VPIInventoryDetail{
				VPItemInventoryID: variantProductInventory.ID,
				SubdistrictID:     VPIInventory.VPIIDetail.SubdistrictID,
			}
			err = tx.Model(&inventorymodel.VPIInventoryDetail{}).Create(&VPIInventoryDetail).Error
			if err != nil {
				fmt.Printf("error creating VPIInventoryDetail \n %+v \n", err)
				tx.Rollback()
				return nil
			}
		}
		// inventory create
	}

	return tx.Commit().Error
}

// func (p VariantProduct) GetItems() []Item {
// 	return p.V
// }
func (p *VariantProduct) FetchProductable() error {
	var err error
	var DB *gorm.DB = database.DB
	vpVariants := []VPVariant{}
	err = DB.Model(&productmodel.VariantProductVariant{}).
		Where("product_id = ?", p.ID).
		Find(&vpVariants).Error

	variantProductItemDBs := []productmodel.VariantProductItem{}
	err = DB.Model(&productmodel.VariantProductItem{}).
		Where("product_id = ?", p.ID).
		Find(&variantProductItemDBs).Error
	if err != nil {
		fmt.Println("error fetching products")
		return nil
	}
	variantProductItems := make([]VariantProductItem, len(variantProductItemDBs))
	for j, itemDB := range variantProductItemDBs {

		itemPricesDB := []VPItemPrice{}
		err = DB.Model(&productmodel.VPItemPrice{}).
			Joins("INNER JOIN VP_items on VP_items.id = VP_item_prices.VP_item_id").
			Where("VP_items.id = ?", itemDB.ID).
			Find(&itemPricesDB).Error

		if err != nil {
			fmt.Println("fetching fetching prices")
			return nil
		}
		itemOptionsDB := []VPOption{}
		err = DB.Model(&productmodel.VariantProductOption{}).
			Where("VP_item_id = ?", itemDB.ID).
			Find(&itemOptionsDB).Error

		variantProductItems[j] = VariantProductItem{
			ID:        variantProductItemDBs[j].ID,
			ProductID: variantProductItemDBs[j].ProductID,
			// VariantProductOptions: variantProductItemDBs[j].VariantProductOptions,
			Weight:                variantProductItemDBs[j].Weight,
			Sku:                   variantProductItemDBs[j].Sku,
			VariantProductOptions: itemOptionsDB,
			VPItemPrices:          itemPricesDB,
		}
	}
	p.VPVariants = vpVariants
	p.VariantProductItems = variantProductItems
	return nil
}

func (p VariantProduct) GetProductable() *Product {
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
		// ProductImages:       p.ProductImages,
		VariantProductItems: p.VariantProductItems,
	}
}

func (p *VariantProduct) SetProductImages(urls []string) {
	for _, url := range urls {

		p.ProductImages = append(p.ProductImages, productmodel.ProductImage{
			URL: url,
		})
	}
}

func NewVariantProduct(productReqParsed producthandler.ProductJSONParsed) (*VariantProduct, error) {
	variantProduct := &VariantProduct{
		BrandID:       productReqParsed.BrandID,
		CategoryID:    productReqParsed.CategoryID,
		ProductTypeID: productReqParsed.ProductTypeID,
		ProductKindID: productReqParsed.ProductKindID,
		Name:          productReqParsed.Name,
		Description:   productReqParsed.Description,
		Status:        productReqParsed.Status,
	}

	variants, err := NewVariantProductVariant(productReqParsed.Variants)
	if err != nil {
		return nil, err
	}
	items, err := NewVariantProductItem(productReqParsed.VariantProductItems)

	if err != nil {
		return nil, err
	}
	variantProduct.VPVariants = variants
	variantProduct.VariantProductItems = items

	return variantProduct, nil
}
