package repository

import (
	"fmt"

	productModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	productAuxModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model/aux"
	"gorm.io/gorm"
)

type M map[string]interface{}

func (r *ProductRepository) GeneratePriceTemplate(productIDArr []int) (productAuxModel.ItemPriceArrExport, error) {
	items := []productModel.ItemForPriceExport{}

	err := r.DB.Model(&productModel.ItemForPriceExport{}).
		Preload("Prices").
		Find(&items).Error

	if err != nil {
		fmt.Printf("error while fetching items")
		return nil, err
	}

	priceXlsx := productAuxModel.ItemPriceArrExport{}
	for _, item := range items {
		for _, price := range item.Prices {
			priceXlsx = append(
				priceXlsx,
				productAuxModel.ItemPriceExport{
					ID:      price.ID,
					ItemID:  price.ItemID,
					ItemSku: item.Sku,
					Name:    price.Name,
					Value:   price.Value,
				},
			)
		}
	}
	return priceXlsx, nil
}

func (r *ProductRepository) ImportPrices(itemPriceArrReq []productModel.ItemPrice) error {
	var err error
	tx := r.DB.Begin()

	// itemPriceArr := []productModel.ItemPrice{}
	// for _, itemPriceReq := range itemPriceArrReq {
	// 	itemPriceArr = append(
	// 		itemPriceArr,
	// 		productModel.ItemPrice{
	// 			ItemID: itemPriceReq.ItemID,
	// 			Name:   itemPriceReq.Name,
	// 			Value:  itemPriceReq.Value,
	// 		},
	// 	)

	// } // 	// add validation here
	fmt.Printf("itemPrices %+v \n", itemPriceArrReq)
	err = tx.Model(&productModel.ItemPrice{}).Create(&itemPriceArrReq).Error

	if err != nil {
		fmt.Printf("error creating prices\n %+v \n", err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r ProductRepository) TxUpdateItemPrices(tx *gorm.DB, itemPriceArrReq []productModel.ItemPrice) (*gorm.DB, error) {
	var err error
	for _, itemPrice := range itemPriceArrReq {
		err = tx.Model(&productModel.ItemPrice{}).Where("id = ?", itemPrice.ID).Updates(itemPrice).Error
		if err != nil {
			fmt.Println("could not update item price")
			return nil, err
		}

	}
	return tx, nil
}
