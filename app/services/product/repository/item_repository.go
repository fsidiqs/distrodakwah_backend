package repository

import (
	"fmt"

	productModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"
)

func (r ProductRepository) TxUpdateItems(tx *gorm.DB, itemArrReq []productModel.Item) (*gorm.DB, error) {
	var err error
	fmt.Printf("items%+v \n", itemArrReq)
	for _, item := range itemArrReq {
		err = tx.Model(&productModel.Item{}).Where("id = ?", item.ID).Updates(item).Error
		if err != nil {
			fmt.Println("could not update item")
			return nil, err
		}

	}
	return tx, nil
}
