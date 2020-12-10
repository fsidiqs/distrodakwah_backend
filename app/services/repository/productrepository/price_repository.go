package productrepository

import (
	"fmt"

	"distrodakwah_backend/app/services/model/productmodel"
	"distrodakwah_backend/app/services/model/productmodel/productauxmodel"

	"gorm.io/gorm"
)

type M map[string]interface{}

func (r *ProductRepository) GeneratePriceTemplate(productIDArr []int) (productauxmodel.ItemPriceArrExport, error) {
	items := []productmodel.ItemForPriceExport{}

	err := r.DB.Model(&productmodel.ItemForPriceExport{}).
		Preload("Prices").
		Find(&items).Error

	if err != nil {
		fmt.Printf("error while fetching items")
		return nil, err
	}

	priceXlsx := productauxmodel.ItemPriceArrExport{}
	for _, item := range items {
		for _, price := range item.Prices {
			priceXlsx = append(
				priceXlsx,
				productauxmodel.ItemPriceExport{
					ID:      price.ID,
					ItemID:  price.ItemID,
					ItemSku: item.Sku,
					Name:    price.Name,
					Value:   price.Value,
				},
			)
		}
	}
	return priceXlsx, nil
}

func (r *ProductRepository) ImportPrices(itemPriceArrReq []productmodel.ItemPrice) error {
	var err error
	tx := r.DB.Begin()

	// itemPriceArr := []productmodel.ItemPrice{}
	// for _, itemPriceReq := range itemPriceArrReq {
	// 	itemPriceArr = append(
	// 		itemPriceArr,
	// 		productmodel.ItemPrice{
	// 			ItemID: itemPriceReq.ItemID,
	// 			Name:   itemPriceReq.Name,
	// 			Value:  itemPriceReq.Value,
	// 		},
	// 	)

	// } // 	// add validation here
	fmt.Printf("itemPrices %+v \n", itemPriceArrReq)
	err = tx.Model(&productmodel.ItemPrice{}).Create(&itemPriceArrReq).Error

	if err != nil {
		fmt.Printf("error creating prices\n %+v \n", err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r ProductRepository) TxUpdateItemPrices(tx *gorm.DB, itemPriceArrReq []productmodel.ItemPrice) (*gorm.DB, error) {
	var err error
	for _, itemPrice := range itemPriceArrReq {
		err = tx.Model(&productmodel.ItemPrice{}).Where("id = ?", itemPrice.ID).Updates(itemPrice).Error
		if err != nil {
			fmt.Println("could not update item price")
			return nil, err
		}

	}
	return tx, nil
}
