package productlibrary

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/model/inventorymodel"
)

type City struct {
	ID         int    `json:"id"`
	ProvinceID int    `json:"province_id"`
	Name       string `json:"name"`
}

func GetCitiesByProvinceID(id int) ([]City, error) {
	var err error

	cities := []City{}
	err = database.DB.Model(&inventorymodel.City{}).
		Where("province_id = ?", id).
		Find(&cities).Error
	if err != nil {
		return nil, err
	}

	return cities, nil
}
