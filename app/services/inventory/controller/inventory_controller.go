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
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/middleware"
	inventoryLibrary "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/library"
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
	itemIDArrReq := c.QueryParam("item_id_arr")

	request := request.FetchAllReq{
		Preload:      []string{},
		ProductIDArr: []int{},
		ItemIDArr:    []int{},
		Metadata: pagination.Metadata{
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
	if itemIDArrReq != "" {
		err = json.NewDecoder(strings.NewReader(itemIDArrReq)).Decode(&request.ItemIDArr)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
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
	itemInventoryID, err := strconv.ParseUint(c.Param("item_inventory_id"), 10, 64)
	preloadReq := c.QueryParam("preload")

	if err != nil {
		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}
	req := repository.FindReq{
		ItemInventoryID: itemInventoryID,
		Preload:         httphelper.Preload{},
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

func (ic *InventoryController) ExportStocks(c echo.Context) error {
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
	userContext := c.(*middleware.UserContext)
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

	stockTempl := []inventoryLibrary.ItemInventoryXlsx{}
	rows := xlsx.GetRows(excelizelib.Sheetname)
	if err != nil {
		return err
	}
	rowsLen := len(rows)

	if rowsLen > 0 {
		for currRow := excelizelib.RowReadStart; currRow < rowsLen; currRow++ {

			ID, _ := strconv.ParseUint(rows[currRow][0], 10, 64)
			stock, _ := strconv.ParseInt(rows[currRow][5], 10, 32)
			keep, _ := strconv.ParseInt(rows[currRow][6], 10, 32)

			stockTempl = append(
				stockTempl,
				inventoryLibrary.ItemInventoryXlsx{
					ID:    uint64(ID),
					Stock: int(stock),
					Keep:  int(keep),
				},
			)
		}
	}

	err = ic.InventoryRepository.PerformInventoryUpdate(stockTempl, userContext.User.ID)
	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)
}
