package controller

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/excelizelib"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httphelper"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model/aux"
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

func (ic *InventoryController) GetProductStock(c echo.Context) error {
	relatedID, err := strconv.ParseUint(c.Param("related_id"), 10, 64)
	productKindID, err := strconv.ParseUint(c.QueryParam("product_kind_id"), 10, 8)
	preloadReq := c.QueryParam("preload")

	if err != nil {
		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}
	req := repository.FindReq{
		RelatedID:     relatedID,
		ProductKindID: uint8(productKindID),
		Preload:       httphelper.Preload{},
	}

	if preloadReq != "" {
		err = json.NewDecoder(strings.NewReader(preloadReq)).Decode(&req.Preload)
	}
	inventory, err := ic.InventoryRepository.Find(req)
	res := &httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    inventory,
	}
	return c.JSON(res.Status, res)

}

func (ic *InventoryController) GenerateExportInventoryTemplate(c echo.Context) error {
	data, err := ic.InventoryRepository.ExportInventory()
	if err != nil {
		c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}
	b, err := excelizelib.PopulateXLSX(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res := c.Response()
	header := res.Header()
	header.Set(echo.HeaderContentType, echo.MIMEOctetStream)
	header.Set(echo.HeaderContentDisposition, "attachment;filename=inventory.xlsx")
	header.Set("Content-Transfer-Encoding", "binary")
	header.Set("Expires", "0")
	res.WriteHeader(http.StatusOK)
	return c.Blob(http.StatusOK, echo.MIMEOctetStream, (*b).Bytes())

}

func (ic *InventoryController) ImportStocks(c echo.Context) error {
	form, err := c.MultipartForm()
	files := form.File["stocks_file"]
	if err != nil {
		fmt.Println(err)
		return err
	}

	var theFile multipart.File

	for _, file := range files {
		// Source

		src, err := file.Open()

		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		theFile, err = file.Open()

		if err != nil {
			return err
		}
		defer theFile.Close()

	}

	xlsx, err := excelize.OpenReader(theFile)
	if err != nil {
		return err
	}

	stockTempl := []*aux.ExcelStockFormat{}
	rows := xlsx.GetRows(excelizelib.Sheetname)
	if err != nil {
		return err
	}

	rowsLen := len(rows)

	if rowsLen > 0 {
		for currRow := excelizelib.RowReadStart; currRow < rowsLen; currRow++ {
			tempProdKindID, _ := strconv.ParseUint(rows[currRow][0], 10, 8)

			tempRelProdID, _ := strconv.ParseUint(rows[currRow][1], 10, 64)
			tempStock, _ := strconv.ParseInt(rows[currRow][3], 10, 32)
			tempKeep, _ := strconv.ParseInt(rows[currRow][4], 10, 32)

			stockTempl = append(
				stockTempl,
				&aux.ExcelStockFormat{
					ProductKindID:    uint8(tempProdKindID),
					RelatedProductID: tempRelProdID,
					Stock:            int(tempStock),
					Keep:             int(tempKeep),
				},
			)
		}
	}

	err = ic.InventoryRepository.PerformStockAdjustment(stockTempl)
	return nil
}
