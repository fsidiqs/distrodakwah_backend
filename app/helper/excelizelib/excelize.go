package excelizelib

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	productModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
)

var Sheetname = "inventory"
var RowReadStart = 1

func PopulateXLSX(data []productModel.Item) (*bytes.Buffer, error) {
	xlsx := excelize.NewFile()

	xlsx.NewSheet(Sheetname)
	xlsx.SetCellValue(Sheetname, "A1", "ID")
	xlsx.SetCellValue(Sheetname, "D1", "Item ID")
	xlsx.SetCellValue(Sheetname, "B1", "Product Kind")
	xlsx.SetCellValue(Sheetname, "C1", "Product ID")
	xlsx.SetCellValue(Sheetname, "E1", "SKU")
	xlsx.SetCellValue(Sheetname, "F1", "Stock")
	xlsx.SetCellValue(Sheetname, "G1", "Keep")

	rowNum := 2
	for _, elem := range data {
		xlsx.SetCellValue(Sheetname, fmt.Sprintf("A%d", rowNum), elem.ItemInventory.ID)
		xlsx.SetCellValue(Sheetname, fmt.Sprintf("B%d", rowNum), elem.ID)
		xlsx.SetCellValue(Sheetname, fmt.Sprintf("C%d", rowNum), elem.Product.ProductKindID)
		xlsx.SetCellValue(Sheetname, fmt.Sprintf("D%d", rowNum), elem.ProductID)
		xlsx.SetCellValue(Sheetname, fmt.Sprintf("E%d", rowNum), elem.Sku)
		xlsx.SetCellValue(Sheetname, fmt.Sprintf("F%d", rowNum), elem.ItemInventory.Stock)
		xlsx.SetCellValue(Sheetname, fmt.Sprintf("G%d", rowNum), elem.ItemInventory.Keep)
		rowNum++
	}

	var b bytes.Buffer
	writr := bufio.NewWriter(&b)
	err := xlsx.Write(writr)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
