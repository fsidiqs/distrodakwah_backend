package repository

import (
	"fmt"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"
)

func (r *ProductRepository) SaveSingleProductTx(singleProduct *model.SingleProduct, tx *gorm.DB) error {
	err := tx.Model(&model.SingleProduct{}).Create(&singleProduct).Error
	if err != nil {
		fmt.Printf("error creating SingleProduct\n %+v \n", err)
		tx.Rollback()
		return err
	}
	return nil
}
