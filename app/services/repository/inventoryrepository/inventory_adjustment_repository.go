package inventoryrepository

import (
	"distrodakwah_backend/app/services/library/inventorylibrary"
	"distrodakwah_backend/app/services/model/inventorymodel"
)

func (r *InventoryRepository) PerformInventoryUpdate(itemInventoryArr []inventorylibrary.ItemInventoryXlsx, userID uint64) error {
	var err error
	tx := r.DB.Begin()
	for _, itemInventory := range itemInventoryArr {

		// save previous stock to a variable
		itemBefore := inventorymodel.ItemInventory{}
		err = tx.Model(&inventorymodel.ItemInventory{}).
			First(&itemBefore, itemInventory.ID).Error
		if err != nil {

			tx.Rollback()
			return err
		}

		// create inventoryadjustment
		err = tx.Model(&inventorymodel.ItemInventoryAdjustment{}).
			Create(&inventorymodel.ItemInventoryAdjustment{
				UserID:          userID,
				ItemInventoryID: itemInventory.ID,
				StockBefore:     itemBefore.Stock,
				StockAfter:      itemInventory.Stock,
			}).Error
		if err != nil {

			tx.Rollback()
			return err
		}
		// update ItemInventory Stock
		err = tx.Model(&inventorymodel.ItemInventory{}).
			Where("id = ?", itemInventory.ID).
			Updates(map[string]interface{}{"stock": itemInventory.Stock}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// func (r *InventoryRepository) PerformStockAdjustment(stocksTempl []*invModelAux.ExcelStockFormat) error {
// 	var err error
// 	// STEP make single_product_inventory_adjustment

// 	// spiAdj := []*model.SPInventoryAdjustment{}
// 	// STEP Extract the id array for querying spProductDB by id
// 	spInvIDArr := []uint64{}
// 	vpInvIDArr := []uint64{}
// 	for _, stock := range stocksTempl {
// 		if stock.ProductKindID == prodModel.ProductKindSingle {
// 			spInvIDArr = append(
// 				spInvIDArr,
// 				stock.RelatedProductID,
// 			)
// 		} else if stock.ProductKindID == prodModel.ProductKindVariant {
// 			vpInvIDArr = append(
// 				vpInvIDArr,
// 				stock.RelatedProductID,
// 			)
// 		}
// 	}
// 	// query InventoryDB by id array

// 	spInventories := []*model.SPInventory{}
// 	err = r.DB.Model(&invModel.SPInventory{}).
// 		Where("single_product_id in (?)", spInvIDArr).
// 		Find(&spInventories).Error
// 	if err != nil {
// 		return err
// 	}
// 	vpInventories := []*model.VPInventory{}
// 	err = r.DB.Model(&invModel.VPInventory{}).
// 		Where("variant_product_id in (?)", vpInvIDArr).
// 		Find(&vpInventories).Error

// 	// STEP Begin Transaction
// 	tx := r.DB.Begin()

// 	// STEP perform single product adjustment and updating stock
// 	spiAdj := []*invModel.SPInventoryAdjustment{}
// 	vpiAdj := []*invModel.VPInventoryAdjustment{}
// 	spInventoriesLen := len(spInventories)
// 	vpInventoriesLen := len(vpInventories)
// 	for _, stock := range stocksTempl {
// 		if stock.ProductKindID == prodModel.ProductKindSingle && spInventoriesLen > 0 {
// 			// STEP find single product inventory DB by current StockRelatedProductID
// spIndex, ok := spInvFindByID(spInventories, stock.RelatedProductID)
// 			theBeforeStock := 0
// 			if ok {
// 				theBeforeStock = spInventories[spIndex].Stock
// 			}
// 			spiAdj = append(
// 				spiAdj,
// 				&invModel.SPInventoryAdjustment{
// 					SPInventoryID: stock.RelatedProductID,
// 					UserID:        1,
// 					StockBefore:   theBeforeStock,
// 					StockAfter:    stock.Stock,
// 				},
// 			)
//STEP updating stock
// spInventory := &invModel.SPInventory{
// 	ID:              spInventories[singleProductID].ID, // qwa
// 	SingleProductID: stock.RelatedProductID,
// 	Stock:           stock.Stock,
// }
// create if not exists or update if conflict, conflict created by //qwa above
// err = tx.Model(&invModel.SPInventory{}).Clauses(clause.OnConflict{
// 	Columns:   []clause.Column{{Name: "ID"}},
// 	DoUpdates: clause.Assignments(map[string]interface{}{"stock": stock.Stock}),
// }).Create(&spInventory).Error
// 			err = tx.Model(&invModel.SPInventory{}).Where("single_product_id = ?", stock.RelatedProductID).Update("stock", stock.Stock).Error
// 		} else if stock.ProductKindID == prodModel.ProductKindVariant && vpInventoriesLen > 0 {

// 			vpIndex, ok := vpInvFindByID(vpInventories, stock.RelatedProductID)
// 			theBeforeStock := 0
// 			if ok {
// 				theBeforeStock = vpInventories[vpIndex].Stock
// 			}

// 			vpiAdj = append(
// 				vpiAdj,
// 				&invModel.VPInventoryAdjustment{
// 					VPInventoryID: stock.RelatedProductID,
// 					UserID:        1,
// 					StockBefore:   theBeforeStock,
// 					StockAfter:    stock.Stock,
// 				},
// 			)

// 			err = tx.Model(&invModel.VPInventory{}).Where("variant_product_id = ?", stock.RelatedProductID).Update("stock", stock.Stock).Error

// 		}
// 	}

// 	if len(spiAdj) > 0 {
// 		err = tx.Model(&invModel.SPInventoryAdjustment{}).Create(&spiAdj).Error
// 	}
// 	if len(vpiAdj) > 0 {

// 		err = tx.Model(&invModel.VPInventoryAdjustment{}).Create(&vpiAdj).Error
// 	}

// 	return tx.Commit().Error
// }
