package repository

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"

func vpInvFindByID(slice []*model.VPInventory, id uint64) (int, bool) {
	for i, v := range slice {
		if v.VariantProductID == id {
			return i, true
		}
	}

	return 0, false
}

func spInvFindByID(slice []*model.SPInventory, id uint64) (int, bool) {
	for i, v := range slice {
		if v.SingleProductID == id {
			return i, true
		}
	}

	return 0, false
}
