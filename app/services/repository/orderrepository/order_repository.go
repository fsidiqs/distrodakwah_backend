package orderrepository

// import (
// 	"time"

// 	invModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"
// 	orderModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/order/model"
// 	prodModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
// )

// func (r *OrderRepository) SaveOrder(orderReq *orderModel.OrderReq) error {
// 	var err error

// 	order := &orderModel.OrderClass{
// 		UserID:              orderReq.UserID,
// 		OrderStatusID:       1,
// 		UniqueCode:          1,
// 		StatusID1Expires:    time.Now().Add(time.Hour * 24),
// 		OrderCustomerDetail: &orderModel.OrderCustomerDetail{},
// 		OrderShippings:      &orderModel.OrderShippings{},
// 		OrderItems:          []orderModel.OrderItemI{},
// 	}

// 	orderItemSingleProductArr := &orderModel.OrderItemSingleProductArr{}
// 	orderItemVariantProductArr := &orderModel.OrderItemVariantProductArr{}
// 	for _, oi := range *orderReq.OrderItems {
// 		if oi.KindID == prodModel.ProductKindSingle {
// 			*orderItemSingleProductArr = append(
// 				*orderItemSingleProductArr,
// 				&orderModel.OrderItemSingleProduct{
// 					SingleProductID: oi.ItemID,
// 					Qty:             oi.Qty,
// 					SPInventory:     &invModel.SPInventory{},
// 					Prices:          &orderModel.OrderItemSPPriceArr{},
// 				},
// 			)
// 		}
// 		if oi.KindID == prodModel.ProductKindVariant {
// 			*orderItemVariantProductArr = append(
// 				*orderItemVariantProductArr,
// 				&orderModel.OrderItemVariantProduct{
// 					VariantProductID: oi.ItemID,
// 					Qty:              oi.Qty,
// 					VPInventory:      &invModel.VPInventory{},
// 					Prices:           &orderModel.OrderItemVPPriceArr{},
// 				},
// 			)
// 		}
// 	}

// 	orderItemSingleProductArr.PopulateDataArr()
// 	orderItemVariantProductArr.PopulateDataArr()
// 	// orderItemLen := len(*orderItemSingleProductArr) + len(*orderItemVariantProductArr)
// 	// orderItemSPArrLen := len(*orderItemSingleProductArr)
// 	// orderItemVPArrLen := len(*orderItemVariantProductArr)

// 	// OrderItemIArr := []orderModel.OrderItemI{}
// 	// asd := []orderModel.OrderItemI{}
// 	for _, oi := range *orderItemSingleProductArr {

// 		order.OrderItems = append(order.OrderItems,
// 			oi,
// 		)
// 	}
// 	for _, oi := range *orderItemVariantProductArr {
// 		order.OrderItems = append(order.OrderItems,
// 			oi,
// 		)
// 	}
// 	// fmt.Printf("fajar sidiq salviro[%+v] %+v \n", test, interf.GetWeight())

// 	// for i := 0; i < orderItemSPArrLen; i++ {
// 	// 	OrderItemIArr[i] = (*orderItemSingleProductArr)[i]
// 	// }
// 	// for i := 0; i < orderItemVPArrLen; i++ {
// 	// 	OrderItemIArr[i] = (*orderItemVariantProductArr)[i]
// 	// }

// 	tx := r.DB.Begin()

// 	err = tx.Model(&orderModel.Order{}).Create(&order).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	orderShippingsObj := &orderModel.OrderShippings{}
// 	order.OrderShippings, err = orderShippingsObj.PrepareShippings(order.ID, order.OrderItems)

// 	err = tx.Model(&orderModel.OrderShipping{}).Create(&order.OrderShippings).Error

// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	for _, oi := range order.OrderItems {
// 		// get ordershipping id
// 		originID := oi.GetVendorSubdistrictID()

// 		ordershippingID, ok := order.OrderShippings.FindOrderShippingIndexByOriginID(originID)
// 		// for _, elem := range *order.OrderShippings {

// 		// }
// 		if !ok {
// 			tx.Rollback()
// 			return err
// 		}

// 		oi.SetOrderShippingID((*order.OrderShippings)[ordershippingID].ID)
// 		productKind, prepareQuery, err := oi.ReturnCreate(order)
// 		if err != nil {
// 			tx.Rollback()
// 			return err
// 		}
// 		if productKind == prodModel.ProductKindSingle {
// 			err = tx.Model(&orderModel.OrderItemSingleProduct{}).Create(prepareQuery).Error
// 		} else if productKind == prodModel.ProductKindVariant {
// 			err = tx.Model(&orderModel.OrderItemVariantProduct{}).Create(prepareQuery).Error

// 		}
// 		if err != nil {
// 			tx.Rollback()
// 			return err
// 		}

// 	}
// 	return tx.Commit().Error
// }

// func findOrderShippingIndexByOriginIDField(slices []*orderModel.OrderShipping, originID int) (int, bool) {
// 	for i := range slices {
// 		if slices[i].SubdistrictIDOrigin == originID {
// 			return i, true
// 		}
// 	}
// 	return 0, false
// }
