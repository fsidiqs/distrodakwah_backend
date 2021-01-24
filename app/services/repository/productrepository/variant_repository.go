package productrepository

import (
	"fmt"

	"distrodakwah_backend/app/services/model/productmodel"

	"gorm.io/gorm"
)

func (r ProductRepository) TxUpdateVariants(tx *gorm.DB, variantArrReq []productmodel.VariantProductVariant) (*gorm.DB, error) {
	var err error
	fmt.Printf("items%+v \n", variantArrReq)
	for _, variant := range variantArrReq {
		err = tx.Model(&productmodel.VariantProductVariant{}).Where("id = ?", variant.ID).Updates(variant).Error
		if err != nil {
			fmt.Println("could not update item")
			return nil, err
		}

	}
	return tx, nil
}
