package inventorylibrary

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/model/inventorymodel"
)

type Province struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetProvinces() ([]Province, error) {
	var err error

	provinces := []Province{}
	err = database.DB.Model(&inventorymodel.Province{}).Find(&provinces).Error
	if err != nil {
		return nil, err
	}

	return provinces, nil
}
