package productrepository

import (
	"fmt"

	"distrodakwah_backend/app/services/model/productmodel"

	"gorm.io/gorm"
)

func (r *ProductRepository) SaveProductProductImageTx(productsProductImage []*productmodel.ProductsProductImage, tx *gorm.DB) error {
	err := tx.Model(&productmodel.ProductsProductImage{}).Create(&productsProductImage).Error
	if err != nil {
		fmt.Printf("error creating ProductsProductImages\n %+v \n", err)
		tx.Rollback()
		return err
	}
	return nil
}
