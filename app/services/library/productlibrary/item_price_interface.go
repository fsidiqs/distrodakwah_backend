package productlibrary

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/model/productmodel"
	"fmt"

	"gorm.io/gorm"
)

type ItemPriceI interface {
	GetByProductIDs([]int) []ItemPrice
}

type ProductPriceI interface {
	GetPricesByProductID([]int)
}

type ItemPrice struct {
	ID            uint
	ProductKindID uint8
	ItemableID    uint
	Name          string
	Value         int
}

type ProductDetailItemPrice struct {
	Name string
	// MainSku    string
	ItemID uint
	Sku    string
	// PriceID    uint
	Kind       uint
	PriceName  string
	PriceValue int
}

type ProductDetailItemPriceArr []ProductDetailItemPrice

func SaveProductPrices(itemPriceArr ProductDetailItemPriceArr) error {
	var err error
	var DB *gorm.DB = database.DB
	tx := DB.Begin()
	spItemPrices := []productmodel.SPItemPrice{}
	vpItemPrices := []productmodel.VPItemPrice{}

	for _, v := range itemPriceArr {
		if v.Kind == productmodel.ProductKindSingle {
			spItemPrices = append(spItemPrices, productmodel.SPItemPrice{
				SPItemID: v.ItemID,
				Name:     v.PriceName,
				Value:    v.PriceValue,
			})
		} else if v.Kind == productmodel.ProductKindVariant {
			vpItemPrices = append(vpItemPrices, productmodel.VPItemPrice{
				VPItemID: v.ItemID,
				Name:     v.PriceName,
				Value:    v.PriceValue,
			})
		}

	}

	if len(spItemPrices) > 0 {
		err = tx.Model(&productmodel.SPItemPrice{}).Create(spItemPrices).Error
		if err != nil {
			fmt.Printf("error creating sp Item Price \n %+v \n", err)
			tx.Rollback()
			return nil
		}
	}

	if len(vpItemPrices) > 0 {
		err = tx.Model(&productmodel.VPItemPrice{}).Create(vpItemPrices).Error
		if err != nil {
			fmt.Printf("error creating vp Item Price \n %+v \n", err)
			tx.Rollback()
			return nil
		}
	}
	return tx.Commit().Error
}
