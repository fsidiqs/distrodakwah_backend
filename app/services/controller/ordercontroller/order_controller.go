package ordercontroller

import (
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/repository/orderrepository"
)

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/labstack/echo"
// 	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httphelper"
// 	orderModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/order/model"
// 	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/order/repository"
// )

type OrderController struct {
	OrderRepository *orderrepository.OrderRepository
}

func (oc *OrderController) PostOrder(c echo.Context) error {
	// 	test := &orderModel.OrderReq{
	// 		Invoice:          "testing2",
	// 		OrderStatusID:    1,
	// 		UniqueCode:       1,
	// 		StatusID1Expires: time.Now().Add(time.Hour * 24),
	// 		UserID:           1,
	// 		OrderItems: &orderModel.OrderItemReqArr{
	// 			&orderModel.OrderItemReq{
	// 				Qty:    3,
	// 				KindID: 1,
	// 				ItemID: 1,
	// 			},
	// 			&orderModel.OrderItemReq{
	// 				Qty:    4,
	// 				KindID: 2,
	// 				ItemID: 3,
	// 			},
	// 			&orderModel.OrderItemReq{
	// 				Qty:    7,
	// 				KindID: 2,
	// 				ItemID: 2,
	// 			},
	// 		},
	// 		OrderCustomerDetail: &orderModel.OrderCustomerDetail{
	// 			CustomerID: 1,
	// 		},
	// 	}

	// 	err := oc.OrderRepository.SaveOrder(test)
	// 	if err != nil {
	// 		fmt.Printf("error jar: %+v \n", err)
	// 		return c.JSON(http.StatusOK, httphelper.StatusOKMessage)

	// 	}
	// 	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)
	//!remove
	return nil

}
