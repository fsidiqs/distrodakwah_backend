package repository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httphelper"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"
)

func HandlePreload(query *gorm.DB, pre httphelper.Preload, prodKindID uint8) {
	for _, value := range pre {
		if value == "inventory_detail" {
			if prodKindID == model.ProductKindSingle {
				query = query.
					Preload("SPInventoryDetail")
			}
			if prodKindID == model.ProductKindVariant {
				query = query.
					Preload("VPInventoryDetail")
			}

		}
	}
}
