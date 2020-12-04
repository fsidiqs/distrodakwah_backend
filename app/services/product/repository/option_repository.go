package repository

import (
	"fmt"

	productModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"
)

func (r ProductRepository) TxUpdateOptions(tx *gorm.DB, optionArrReq []productModel.Option) (*gorm.DB, error) {
	var err error
	fmt.Printf("items%+v \n", optionArrReq)
	for _, option := range optionArrReq {
		err = tx.Model(&productModel.Option{}).Where("id = ?", option.ID).Updates(option).Error
		if err != nil {
			fmt.Println("could not update item")
			return nil, err
		}

	}
	return tx, nil
}
