package inventorycontroller

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"

	"distrodakwah_backend/app/helper/excelizelib"
	"distrodakwah_backend/app/helper/httphelper"
	"distrodakwah_backend/app/helper/pagination"
	"distrodakwah_backend/app/services/handler/inventoryhandler"
	"distrodakwah_backend/app/services/library/productlibrary"
	"distrodakwah_backend/app/services/repository/inventoryrepository"
)

type InventoryController struct {
	InventoryRepository *inventoryrepository.InventoryRepository
}

func (ic *InventoryController) GetProductStocks(c echo.Context) error {
	pageReq, err := strconv.Atoi(c.QueryParam("p_page"))
	limitReq, err := strconv.Atoi(c.QueryParam("p_limit"))
	preloadReq := c.QueryParam("preload")
	productIDArrReq := c.QueryParam("product_id_arr")
	itemIDArrReq := c.QueryParam("item_id_arr")

	request := inventoryhandler.FetchAllReq{
		Preload:      []string{},
		ProductIDArr: []uint{},
		ItemIDArr:    []uint{},
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
	itemInventoryID, err := strconv.Atoi(c.Param("item_inventory_id"))
	preloadReq := c.QueryParam("preload")

	if err != nil {
		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}
	req := inventoryrepository.FindReq{
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
	var err error
	productIDArrReq := c.QueryParam("product_id_arr")
	var productIDArr []uint
	if productIDArrReq != "" {
		err = json.NewDecoder(strings.NewReader(productIDArrReq)).Decode(&productIDArr)
	}

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// data, err := ic.InventoryRepository.ExportInventory()
	data, err := productlibrary.GetAllProductStocks(productIDArr)

	if err != nil {
		c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}
	b, err := productlibrary.PopulateProductStockXLSX(data)
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
	// userContext := c.(*middleware.UserContext)
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

	stockTempl := []productlibrary.ProductStock{}
	rows := xlsx.GetRows(excelizelib.Sheetname)
	if err != nil {
		return err
	}
	rowsLen := len(rows)

	if rowsLen > 0 {
		for currRow := excelizelib.RowReadStart; currRow < rowsLen; currRow++ {

			ID, _ := strconv.ParseUint(rows[currRow][1], 10, 64)
			kind, _ := strconv.ParseUint(rows[currRow][3], 10, 64)

			stock, _ := strconv.ParseInt(rows[currRow][4], 10, 32)

			stockTempl = append(
				stockTempl,
				productlibrary.ProductStock{
					ItemInventoryID: uint(ID),
					Kind:            uint(kind),
					Stock:           int(stock),
				},
			)
		}
	}
	//! user context must login
	// err = productlibrary.SaveProductStocks(stockTempl, userContext.User.ID)
	err = productlibrary.SaveProductStocks(stockTempl, 1)

	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)
}
