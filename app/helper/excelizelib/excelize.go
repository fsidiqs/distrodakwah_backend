package excelizelib

import (
	"bytes"

	"distrodakwah_backend/app/services/model/productmodel"
)

var Sheetname = "inventory"
var RowReadStart = 1

func PopulateXLSX(data []productmodel.Item) (*bytes.Buffer, error) {
	// xlsx := excelize.NewFile()

	// xlsx.NewSheet(Sheetname)
	// xlsx.SetCellValue(Sheetname, "A1", "ID")
	// xlsx.SetCellValue(Sheetname, "D1", "Item ID")
	// xlsx.SetCellValue(Sheetname, "B1", "Product Kind")
	// xlsx.SetCellValue(Sheetname, "C1", "Product ID")
	// xlsx.SetCellValue(Sheetname, "E1", "SKU")
	// xlsx.SetCellValue(Sheetname, "F1", "Stock")
	// xlsx.SetCellValue(Sheetname, "G1", "Keep")
	// xlsx.SetCellValue(Sheetname, "H1", "Location")

	// rowNum := 2
	// for _, elem := range data {
	// 	for _, itemInventory := range elem.ItemInventory {
	// 		fmt.Printf("inventory :%+v\n", itemInventory.ItemInventoryDetail.Subdistrict)
	// 		xlsx.SetCellValue(Sheetname, fmt.Sprintf("A%d", rowNum), itemInventory.ID)
	// 		xlsx.SetCellValue(Sheetname, fmt.Sprintf("B%d", rowNum), elem.ID)
	// 		xlsx.SetCellValue(Sheetname, fmt.Sprintf("C%d", rowNum), elem.Product.ProductKindID)
	// 		xlsx.SetCellValue(Sheetname, fmt.Sprintf("D%d", rowNum), elem.ProductID)
	// 		xlsx.SetCellValue(Sheetname, fmt.Sprintf("E%d", rowNum), elem.Sku)
	// 		xlsx.SetCellValue(Sheetname, fmt.Sprintf("F%d", rowNum), itemInventory.Stock)
	// 		xlsx.SetCellValue(Sheetname, fmt.Sprintf("G%d", rowNum), itemInventory.Keep)
	// 		xlsx.SetCellValue(Sheetname, fmt.Sprintf("H%d", rowNum), itemInventory.ItemInventoryDetail.Subdistrict.Name)
	// 		rowNum++
	// 	}
	// }

	// var b bytes.Buffer
	// writr := bufio.NewWriter(&b)
	// err := xlsx.Write(writr)
	// if err != nil {
	// 	return nil, err
	// }
	// return &b, nil

	return nil, nil
}
