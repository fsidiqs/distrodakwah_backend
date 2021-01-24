package productlibrary

import (
	"distrodakwah_backend/app/services/model/productmodel"
	"fmt"
	"mime/multipart"

	"gorm.io/gorm"
)

type ProductImageUrl string

type ProductImageURL struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}
type ProductImage struct {
	FileName string
	Content  multipart.File
}

// func ConstructProductImages(imageUrls string) ([]ProductImageUrl, error) {
// 	productImageUrlsReqs := []ProductImageUrl{}
// 	err := json.NewDecoder(strings.NewReader(imageUrls)).Decode(&productImageUrlsReqs)
// 	if err != nil {
// 		return nil, err
// 	}
// 	productImageUrls := make([]ProductImageUrl, len(productImageUrlsReqs))
// 	for i, productImageUrlsReq := range productImageUrlsReqs {
// 		productImageUrls[i] = productImageUrlsReq
// 	}
// 	return productImageUrls, nil
// }

func SaveImagesTx(productImages []productmodel.ProductImage, tx *gorm.DB) error {

	err := tx.Model(&productmodel.ProductImage{}).Create(&productImages).Error

	if err != nil {
		fmt.Printf("error creating images\n %+v \n", err)
		tx.Rollback()
		return err
	}
	return nil
}
