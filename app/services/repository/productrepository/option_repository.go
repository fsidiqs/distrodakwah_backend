package productrepository

import (
	"fmt"

	"distrodakwah_backend/app/services/model/productmodel"

	"gorm.io/gorm"
)

func (r ProductRepository) TxUpdateOptions(tx *gorm.DB, optionArrReq []productmodel.VariantProductOption) (*gorm.DB, error) {
	var err error
	fmt.Printf("items%+v \n", optionArrReq)
	for _, option := range optionArrReq {
		err = tx.Model(&productmodel.VariantProductOption{}).Where("id = ?", option.ID).Updates(option).Error
		if err != nil {
			fmt.Println("could not update item")
			return nil, err
		}

	}
	return tx, nil
}
