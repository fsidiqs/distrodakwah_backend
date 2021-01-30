package productlibrary

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/library/query/productquery"
	"distrodakwah_backend/app/services/model/productmodel"
	"fmt"

	"gorm.io/gorm"
)

func GetAllProductStocks(prodIDArr []uint) ([]ProductStock, error) {
	var err error
	var DB *gorm.DB = database.DB
	allProducts := []AllProduct{}
	err = DB.Raw(productquery.SEL_ALL_PRODUCTS).Scan(&allProducts).Error

	if err != nil {
		fmt.Println("error fetching all products")
		return nil, err
	}
	// filter to each kind
	spIDs := []uint{}
	vpIDs := []uint{}
	for _, ap := range allProducts {
		if ap.Kind == productmodel.ProductKindSingle {
			spIDs = append(spIDs, ap.ID)
		} else if ap.Kind == productmodel.ProductKindVariant {
			vpIDs = append(vpIDs, ap.ID)
		}
	}

	productStocks := []ProductStock{}
	err = DB.Raw(productquery.SEL_PRODUCT_STOCKS_BY_ID, spIDs, vpIDs).Find(&productStocks).Error
	return productStocks, nil

}
