package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httphelper"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/repository"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/request"
)

type InventoryController struct {
	InventoryRepository *repository.InventoryRepository
}

func (ic *InventoryController) GetProductStocks(c echo.Context) error {
	pageReq, err := strconv.Atoi(c.QueryParam("p_page"))
	limitReq, err := strconv.Atoi(c.QueryParam("p_limit"))
	preloadReq := c.QueryParam("preload")
	productIDArrReq := c.QueryParam("product_id_arr")

	request := &request.FetchAllReq{
		Preload:      []string{},
		ProductIDArr: []int{},
		SPIDArr:      []int{},
		VPIDArr:      []int{},
		Metadata: &pagination.Metadata{
			Page:  pageReq,
			Limit: limitReq,
		},
	}

	if preloadReq != "" {
		err = json.NewDecoder(strings.NewReader(preloadReq)).Decode(&request.Preload)
	}
	if productIDArrReq != "" {
		err = json.NewDecoder(strings.NewReader(productIDArrReq)).Decode(&request.ProductIDArr)
	}

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "page_and_limit_empty")
	}
	data, err := ic.InventoryRepository.FetchAll(request)
	res := &httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    data,
	}
	return c.JSON(res.Status, res)

}

func (ic *InventoryController) GenerateExportInventoryTemplate(c echo.Context) error {
	data, err := ic.InventoryRepository.ExportInventory()
	if err != nil {
		c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}
	res := &httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    data,
	}
	return c.JSON(res.Status, res)

}
