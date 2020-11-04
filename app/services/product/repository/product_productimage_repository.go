package repository

import (
	"fmt"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"
)

func (r *ProductRepository) SaveProductProductImageTx(productsProductImage []*model.ProductsProductImage, tx *gorm.DB) error {
	err := tx.Model(&model.ProductsProductImage{}).Create(&productsProductImage).Error
	if err != nil {
		fmt.Printf("error creating ProductsProductImages\n %+v \n", err)
		fmt.Println("test")

		tx.Rollback()
		return err
	}
	return nil
}
