package productlibrary

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/model/inventorymodel"
)

type Subdistrict struct {
	ID     int    `json:"id"`
	CityID int    `json:"city_id"`
	Name   string `json:"name"`
}

func GetSubdistrictsByCityID(id int) ([]Subdistrict, error) {
	var err error

	subdistricts := []Subdistrict{}
	err = database.DB.Model(&inventorymodel.Subdistrict{}).
		Where("city_id = ?", id).
		Find(&subdistricts).Error
	if err != nil {
		return nil, err
	}

	return subdistricts, nil
}
