package productlibrary

import (
	"distrodakwah_backend/app/services/model/inventorymodel"
	"distrodakwah_backend/app/services/model/productmodel"
	"fmt"
)

type ProductStock struct {
	Name            string
	ItemInventoryID uint
	Sku             string
	Kind            uint
	Stock           int
}

func SaveProductStocks(stocks []ProductStock, userID int) error {
	var err error
	spiIntenvories := []inventorymodel.SPIInventory{}
	vpiInventories := []inventorymodel.VPIInventory{}

	for _, v := range stocks {
		if v.Kind == productmodel.ProductKindSingle {
			spiIntenvories = append(spiIntenvories, inventorymodel.SPIInventory{
				ID:    v.ItemInventoryID,
				Stock: v.Stock,
			})
		} else if v.Kind == productmodel.ProductKindVariant {
			vpiInventories = append(vpiInventories, inventorymodel.VPIInventory{
				ID:    v.ItemInventoryID,
				Stock: v.Stock,
			})
		}
	}

	if len(spiIntenvories) > 0 {
		err = UpdateSPItemStock(spiIntenvories, userID)
		if err != nil {
			fmt.Println("error updating product")
			return err
		}
	}

	if len(vpiInventories) > 0 {
		err = UpdateVPItemStock(vpiInventories, userID)
		if err != nil {
			fmt.Println("error updating product")
			return err
		}
	}

	return nil
}
