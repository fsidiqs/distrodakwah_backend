package productlibrary

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const InventorySheetName = "inventory"

var RowReadStart = 1

func PopulateProductStockXLSX(data []ProductStock) (*bytes.Buffer, error) {
	xlsx := excelize.NewFile()

	xlsx.NewSheet(InventorySheetName)
	xlsx.SetCellValue(InventorySheetName, "A1", "Produt Name")
	xlsx.SetCellValue(InventorySheetName, "B1", "ItemInventoryID")
	xlsx.SetCellValue(InventorySheetName, "C1", "Sku")
	xlsx.SetCellValue(InventorySheetName, "D1", "Kind")
	xlsx.SetCellValue(InventorySheetName, "E1", "Stock")

	rowNum := 2
	for _, elem := range data {
		xlsx.SetCellValue(InventorySheetName, fmt.Sprintf("A%d", rowNum), elem.Name)
		xlsx.SetCellValue(InventorySheetName, fmt.Sprintf("B%d", rowNum), elem.ItemInventoryID)
		xlsx.SetCellValue(InventorySheetName, fmt.Sprintf("C%d", rowNum), elem.Sku)
		xlsx.SetCellValue(InventorySheetName, fmt.Sprintf("D%d", rowNum), elem.Kind)
		xlsx.SetCellValue(InventorySheetName, fmt.Sprintf("E%d", rowNum), elem.Stock)

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
