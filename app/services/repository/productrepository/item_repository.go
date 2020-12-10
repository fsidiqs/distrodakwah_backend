package productrepository

import (
	"fmt"

	"distrodakwah_backend/app/services/model/productmodel"

	"gorm.io/gorm"
)

func (r ProductRepository) TxUpdateItems(tx *gorm.DB, itemArrReq []productmodel.Item) (*gorm.DB, error) {
	var err error
	fmt.Printf("items%+v \n", itemArrReq)
	for _, item := range itemArrReq {
		err = tx.Model(&productmodel.Item{}).Where("id = ?", item.ID).Updates(item).Error
		if err != nil {
			fmt.Println("could not update item")
			return nil, err
		}

	}
	return tx, nil
}
