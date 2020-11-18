package excelizelib

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	prodModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
)

var Sheetname = "inventory"
var RowReadStart = 1

func PopulateXLSX(data []*prodModel.ProductInventory) (*bytes.Buffer, error) {
	xlsx := excelize.NewFile()

	xlsx.NewSheet(Sheetname)
	xlsx.SetCellValue(Sheetname, "A1", "Product Kind")
	xlsx.SetCellValue(Sheetname, "B1", "Related Product ID")
	xlsx.SetCellValue(Sheetname, "C1", "SKU")
	xlsx.SetCellValue(Sheetname, "D1", "Stock")
	xlsx.SetCellValue(Sheetname, "E1", "Keep Stock")

	rowNum := 2
	for _, product := range data {

		if product.ProductKindID == prodModel.ProductKindSingle {
			xlsx.SetCellValue(Sheetname, fmt.Sprintf("A%d", rowNum), product.ProductKindID)
			xlsx.SetCellValue(Sheetname, fmt.Sprintf("B%d", rowNum), product.SingleProduct.ID)
			xlsx.SetCellValue(Sheetname, fmt.Sprintf("C%d", rowNum), product.Sku)
			if product.SingleProduct.SPInventory == nil {
				xlsx.SetCellValue(Sheetname, fmt.Sprintf("D%d", rowNum), "undefined")
				xlsx.SetCellValue(Sheetname, fmt.Sprintf("E%d", rowNum), "undefined")
			} else {
				xlsx.SetCellValue(Sheetname, fmt.Sprintf("D%d", rowNum), product.SingleProduct.SPInventory.Stock)
				xlsx.SetCellValue(Sheetname, fmt.Sprintf("E%d", rowNum), product.SingleProduct.SPInventory.Keep)
			}
			rowNum++
		} else if product.ProductKindID == prodModel.ProductKindVariant {
			for _, variant := range product.VariantProducts {
				xlsx.SetCellValue(Sheetname, fmt.Sprintf("A%d", rowNum), product.ProductKindID)
				xlsx.SetCellValue(Sheetname, fmt.Sprintf("B%d", rowNum), variant.ID)
				xlsx.SetCellValue(Sheetname, fmt.Sprintf("C%d", rowNum), variant.Sku)
				if variant.VPInventory == nil {
					xlsx.SetCellValue(Sheetname, fmt.Sprintf("D%d", rowNum), "undefined")
					xlsx.SetCellValue(Sheetname, fmt.Sprintf("E%d", rowNum), "undefined")
				} else {
					xlsx.SetCellValue(Sheetname, fmt.Sprintf("D%d", rowNum), variant.VPInventory.Stock)
					xlsx.SetCellValue(Sheetname, fmt.Sprintf("E%d", rowNum), variant.VPInventory.Keep)
				}
				rowNum++
			}

		}

	}

	var b bytes.Buffer
	writr := bufio.NewWriter(&b)
	err := xlsx.Write(writr)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
