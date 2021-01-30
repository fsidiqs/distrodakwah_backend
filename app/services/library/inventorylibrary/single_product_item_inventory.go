package inventorylibrary

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/model/inventorymodel"
	"encoding/json"
	"strings"

	"gorm.io/gorm"
)

type SPIInventory struct {
	ID                 uint                `gorm:"primaryKey;autoIncrement;not null"`
	SPItemID           uint                `gorm:"column:SP_item_id;type:BIGINT;UNSIGNED;NOT NULL" json:"sp_item_id"`
	Stock              int                 `gorm:"type:INT;NOT NULL" json:"stock"`
	Keep               int                 `gorm:"type:INT;NOT NULL" json:"keep"`
	SPIInventoryDetail *SPIInventoryDetail `gorm:"foreignKey:SPItemInventoryID" json:"item_inventory_detail"`
}

func (SPIInventory) TableName() string {
	return "SP_item_inventories"
}

func NewSPIInventory(jsonReq string) ([]SPIInventory, error) {
	// contains slice of subdistrict id
	parsedReq := []int{}

	err := json.NewDecoder(strings.NewReader(jsonReq)).Decode(&parsedReq)

	if err != nil {
		return nil, err
	}

	SPIInventories := make([]SPIInventory, len(parsedReq))
	for i, subdistrictID := range parsedReq {
		SPIInventories[i] = SPIInventory{
			SPIInventoryDetail: &SPIInventoryDetail{
				SubdistrictID: subdistrictID,
			},
		}
	}

	return SPIInventories, nil
}

func UpdateSPItemStock(spiIntenvories []inventorymodel.SPIInventory, userID int) error {
	var err error
	var DB *gorm.DB = database.DB
	tx := DB.Begin()

	for _, itemInventory := range spiIntenvories {

		// save previous stock to a variable
		itemBefore := inventorymodel.SPIInventory{}
		err = tx.Model(&inventorymodel.SPIInventory{}).
			First(&itemBefore, itemInventory.ID).Error
		if err != nil {

			tx.Rollback()
			return err
		}

		// create inventoryadjustment
		err = tx.Model(&inventorymodel.SPIInventoryAdjustment{}).
			Create(&inventorymodel.SPIInventoryAdjustment{
				UserID:         userID,
				SPIInventoryID: itemInventory.ID,
				StockBefore:    itemBefore.Stock,
				StockAfter:     itemInventory.Stock,
			}).Error
		if err != nil {

			tx.Rollback()
			return err
		}
		// update ItemInventory Stock
		err = tx.Model(&inventorymodel.SPIInventory{}).
			Where("id = ?", itemInventory.ID).
			Updates(map[string]interface{}{"stock": itemInventory.Stock}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
