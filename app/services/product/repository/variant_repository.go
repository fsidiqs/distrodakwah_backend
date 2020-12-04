package repository

import (
	"fmt"

	productModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"
)

func (r ProductRepository) TxUpdateVariants(tx *gorm.DB, variantArrReq []productModel.Variant) (*gorm.DB, error) {
	var err error
	fmt.Printf("items%+v \n", variantArrReq)
	for _, variant := range variantArrReq {
		err = tx.Model(&productModel.Variant{}).Where("id = ?", variant.ID).Updates(variant).Error
		if err != nil {
			fmt.Println("could not update item")
			return nil, err
		}

	}
	return tx, nil
}
