package ordermodel

// import (
// 	"errors"

// 	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/sliceutils"
// )

// // func (o *OrderReq) PopulateOrderItemsData() error {
// // 	err := o.OrderItems.PopulateData()
// // 	if err != nil {
// // 		return err
// // 	}
// // 	return nil
// // }

// // func (o *OrderReq) PopulateData() error {
// // 	// populate orderItems
// // 	err := o.OrderItems.PopulateData()
// // 	if err != nil {
// // 		return err
// // 	}
// // 	// calculate order.Total()
// // 	o.CalculateTotal()
// // 	return nil
// // }

// // func (o *OrderReq) CalculateTotal() {
// // 	o.Total = o.OrderItems.SumDropshipperItemPrice()
// // }

// func (o *OrderReq) PrepareShippings() error {
// 	// var err error
// 	// prepare order shipping slices

// 	orderShippingOriginIDs := []int{}
// 	for _, oi := range *o.OrderItems {
// 		id, ok := oi.GetProductSubdistrictID()

// 		if !ok {
// 			return errors.New("Error interface")
// 		}
// 		orderShippingOriginIDs = append(orderShippingOriginIDs, id)
// 	}

// 	uniqueOrderShippingOriginIDs := sliceutils.UniqueInts(orderShippingOriginIDs)
// 	/// get SHipping company dirty way
// 	// shippings := []*shipModel.ShippingCompany{}
// 	// err = database.DB.Model(&shipModel.ShippingCompany{}).Find(&shippings).Error
// 	o.OrderShippings = make([]*OrderShipping, len(uniqueOrderShippingOriginIDs))
// 	for i, osid := range uniqueOrderShippingOriginIDs {
// 		// shipdetails := rajaongkir.ShippingDetails{
// 		// 	OriginSubID: osid,
// 		// 	// ShipName: ,
// 		// }
// 		// abc := rajaongkir.GetCost()
// 		o.OrderShippings[i] = &OrderShipping{
// 			SubdistrictIDOrigin: osid,
// 			TypeID:              1,
// 			// calculate shipping cost
// 			ShippingCost:        5,
// 			ShippingCompanyID:   1,
// 			ShippingServiceName: "asd",
// 			OrderID:             o.ID,
// 		}
// 	}

// 	for _, oi := range *o.OrderItems {
// 		// orderItemInterface = oi

// 		originID, ok := oi.GetProductSubdistrictID()
// 		if !ok {
// 			return errors.New("Error interface")
// 		}

// 		ordershippingID, ok := o.OrderShippings.FindOrderShippingIndexByOriginID(originID)

// 		o.OrderShippings[ordershippingID].TotalCost += float64(oi.Qty) * oi.DropshipperItemPrice
// 		o.OrderShippings[ordershippingID].Weight += (oi.Qty) * oi.UnitWeight

// 	}
// 	return nil
// }
