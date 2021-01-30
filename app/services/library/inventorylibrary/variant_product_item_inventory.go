package inventorylibrary

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/model/inventorymodel"
	"encoding/json"
	"strings"

	"gorm.io/gorm"
)

type VPIInventory struct {
	ID                 uint                `gorm:"primaryKey;autoIncrement;not null"`
	VPItemID           uint                `gorm:"column:VP_item_id;type:BIGINT;UNSIGNED;NOT NULL" json:"vp_item_id"`
	Stock              int                 `gorm:"type:INT;NOT NULL" json:"stock"`
	Keep               int                 `gorm:"type:INT;NOT NULL" json:"keep"`
	VPIInventoryDetail *VPIInventoryDetail `gorm:"foreignKey:VPItemInventoryID" json:"item_inventory_detail"`
}

func (VPIInventory) TableName() string {
	return "VP_item_inventories"
}
func NewVPIInventory(jsonReq string) ([]VPIInventory, error) {
	// contains slice of subdistrict id
	parsedReq := []int{}

	err := json.NewDecoder(strings.NewReader(jsonReq)).Decode(&parsedReq)

	if err != nil {
		return nil, err
	}

	VPIInventories := make([]VPIInventory, len(parsedReq))
	for i, subdistrictID := range parsedReq {
		VPIInventories[i] = VPIInventory{
			VPIInventoryDetail: &VPIInventoryDetail{
				SubdistrictID: subdistrictID,
			},
		}
	}

	return VPIInventories, nil
}

func UpdateVPItemStock(spiIntenvories []inventorymodel.VPIInventory, userID int) error {
	var err error
	var DB *gorm.DB = database.DB
	tx := DB.Begin()

	for _, itemInventory := range spiIntenvories {

		// save previous stock to a variable
		itemBefore := inventorymodel.VPIInventory{}
		err = tx.Model(&inventorymodel.VPIInventory{}).
			First(&itemBefore, itemInventory.ID).Error
		if err != nil {

			tx.Rollback()
			return err
		}

		// create inventoryadjustment
		err = tx.Model(&inventorymodel.VPIInventoryAdjustment{}).
			Create(&inventorymodel.VPIInventoryAdjustment{
				UserID:         userID,
				VPIInventoryID: itemInventory.ID,
				StockBefore:    itemBefore.Stock,
				StockAfter:     itemInventory.Stock,
			}).Error
		if err != nil {

			tx.Rollback()
			return err
		}
		// update ItemInventory Stock
		err = tx.Model(&inventorymodel.VPIInventory{}).
			Where("id = ?", itemInventory.ID).
			Updates(map[string]interface{}{"stock": itemInventory.Stock}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
