package repository

import (
	"fmt"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"
)

func (r *ProductRepository) SaveImagesTx(images []*model.ProductImage, tx *gorm.DB) error {
	for _, image := range images {
		fmt.Printf("images befores: %+v \n", image.ID)

	}

	err := tx.Model(&model.ProductImage{}).Create(&images).Error
	for _, image := range images {
		fmt.Printf("images afters: %+v \n", image.ID)

	}
	if err != nil {
		fmt.Printf("error creating images\n %+v \n", err)
		tx.Rollback()
		return err
	}
	return nil
}
