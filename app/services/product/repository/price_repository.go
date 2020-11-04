package repository

import (
	"fmt"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"
)

func (r *ProductRepository) SavePricesTx(prices []*model.Price, tx *gorm.DB) error {
	err := tx.Model(&model.Price{}).Create(&prices).Error
	if err != nil {
		fmt.Printf("error creating prices \n %+v \n", err)
		tx.Rollback()
		return err
	}

	return nil
}
