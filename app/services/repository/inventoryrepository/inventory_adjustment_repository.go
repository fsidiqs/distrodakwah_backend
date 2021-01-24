package inventoryrepository

import (
	"distrodakwah_backend/app/services/library/inventorylibrary"
)

func (r *InventoryRepository) PerformInventoryUpdate(itemInventoryArr []inventorylibrary.ItemInventoryXlsx, userID int) error {
	return nil
	// var err error
	// tx := r.DB.Begin()
	// for _, itemInventory := range itemInventoryArr {

	// 	// save previous stock to a variable
	// 	itemBefore := inventorymodel.ItemInventory{}
	// 	err = tx.Model(&inventorymodel.ItemInventory{}).
	// 		First(&itemBefore, itemInventory.ID).Error
	// 	if err != nil {

	// 		tx.Rollback()
	// 		return err
	// 	}

	// 	// create inventoryadjustment
	// 	err = tx.Model(&inventorymodel.ItemInventoryAdjustment{}).
	// 		Create(&inventorymodel.ItemInventoryAdjustment{
	// 			UserID:          userID,
	// 			ItemInventoryID: itemInventory.ID,
	// 			StockBefore:     itemBefore.Stock,
	// 			StockAfter:      itemInventory.Stock,
	// 		}).Error
	// 	if err != nil {

	// 		tx.Rollback()
	// 		return err
	// 	}
	// 	// update ItemInventory Stock
	// 	err = tx.Model(&inventorymodel.ItemInventory{}).
	// 		Where("id = ?", itemInventory.ID).
	// 		Updates(map[string]interface{}{"stock": itemInventory.Stock}).Error
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// }
	// return tx.Commit().Error
}
