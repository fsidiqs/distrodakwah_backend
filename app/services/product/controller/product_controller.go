package controller

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	productClass "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/class"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httphelper"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model/aux"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/repository"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/request"
)

type ProductController struct {
	ProductRepository *repository.ProductRepository
}

func (pc *ProductController) GetProductsByColumn(c echo.Context) error {
	pageReq, err := strconv.Atoi(c.QueryParam("p_page"))
	limitReq, err := strconv.Atoi(c.QueryParam("p_limit"))
	preloadReq := c.QueryParam("preload")

	urlVal := c.QueryParams()
	request := &request.FetchByColumnReq{
		Preload:  []string{},
		PKindIDs: []int{},
		PTypeIDs: []int{},
		Metadata: &pagination.Metadata{
			Page:  pageReq,
			Limit: limitReq,
		},
	}

	if preloadReq != "" {
		err = json.NewDecoder(strings.NewReader(preloadReq)).Decode(&request.Preload)
	}

	if err := request.Mydecode(urlVal); err != nil {
		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}

	data, err := pc.ProductRepository.FetchByColumns(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httphelper.InternalServerErrorMessage)
	}
	res := &httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    data,
	}
	return c.JSON(res.Status, res)
}

func (pc *ProductController) Gets(c echo.Context) error {
	pageReq, err := strconv.Atoi(c.QueryParam("p_page"))
	limitReq, err := strconv.Atoi(c.QueryParam("p_limit"))
	preloadReq := c.QueryParam("preload")
	productIDArrReq := c.QueryParam("product_id_arr")

	request := &request.FetchAllReq{
		Preload:      []string{},
		ProductIDArr: []int{},
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
	data, err := pc.ProductRepository.FetchAll(request)
	res := &httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    data,
	}
	return c.JSON(res.Status, res)
}

func (pc *ProductController) Post(c echo.Context) error {

	product := &model.ProductFromRequestJSON{}
	// var product map[string]interface{}
	if err := c.Bind(&product); err != nil {
		fmt.Printf("error: %+v ", err)
		return err
	}

	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)
}

func (pc *ProductController) CreateProductBasicStructure(c echo.Context) (err error) {
	formReq, _ := c.MultipartForm()
	productReq := c.FormValue("product")
	files := formReq.File["product_images"]
	product := &model.ProductFromRequestJSON{}

	theFiles := make([]productClass.ProductImage, len(files))

	for i, file := range files {
		// Source

		src, err := file.Open()

		if err != nil {
			fmt.Println(err)
			return err
		}
		defer src.Close()

		// Destination
		theFiles[i].FileName = file.Filename
		theFiles[i].Content, err = file.Open()
		if err != nil {
			return err
		}
		defer theFiles[i].Content.Close()

	}

	// ! Update this
	// productImageURLs, err := digitalocean.UploadFiles(theFiles)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Image Upload failed")
	}
	if productReq != "" {
		err = json.NewDecoder(strings.NewReader(productReq)).Decode(&product)
	}
	productImageURLs := []string{"test.jpeg"}
	for _, imgURL := range productImageURLs {
		product.ProductImages = append(
			product.ProductImages,
			&model.ProductImage{
				URL: imgURL,
			},
		)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}

	err = pc.ProductRepository.SaveProductBasicStructure(product)
	if err != nil {
		fmt.Printf("error creating product: %+v", err)
		return c.JSON(http.StatusInternalServerError, httphelper.InternalServerErrorMessage)
	}
	fmt.Println("asd")
	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)
}

func (pc *ProductController) UpdateProduct(c echo.Context) (err error) {
	// formReq, _ := c.MultipartForm()
	productReq := c.FormValue("product")

	_, err = productClass.ProductDetailJSONDecoder(productReq)
	return nil
}

func (pc *ProductController) ImportPrices(c echo.Context) error {
	form, err := c.MultipartForm()
	files := form.File["prices_file"]
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
	pricesTemplate := &aux.ProductPriceTemplate{}

	rows := xlsx.GetRows("Single Product Prices")
	if err != nil {
		return err
	}

	rowsLen := len(rows)
	if rowsLen > 0 {

		for i := 1; i < rowsLen; i++ {
			tempSingleProductID, _ := strconv.ParseUint(rows[i][0], 10, 64)
			tempPriceValue, _ := strconv.ParseFloat(rows[i][3], 10)
			pricesTemplate.SingleProductPricesTemplate = append(
				pricesTemplate.SingleProductPricesTemplate,
				&model.SingleProductPriceTemplate{
					SingleProductID: tempSingleProductID,

					PriceName: rows[i][2],

					PriceValue: tempPriceValue,
				},
			)

		}
	}

	rows = xlsx.GetRows("Variant Product Prices")
	if err != nil {
		return err
	}

	rowsLen = len(rows)

	if rowsLen > 0 {
		for i := 1; i < rowsLen; i++ {
			tempVariantProductID, _ := strconv.ParseUint(rows[i][0], 10, 64)
			tempPriceValue, _ := strconv.ParseFloat(rows[i][3], 10)

			pricesTemplate.VariantProductPriceTemplate = append(
				pricesTemplate.VariantProductPriceTemplate,
				&model.VariantProductPriceTemplate{
					VariantProductID: tempVariantProductID,

					PriceName: rows[i][2],

					PriceValue: tempPriceValue,
				},
			)

		}
	}

	err = pc.ProductRepository.ImportPrices(pricesTemplate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)
}

func (pc *ProductController) GeneratePriceTemplate(c echo.Context) (err error) {
	productIDArrReq := c.QueryParam("product_id_arr")
	var productIDArr []int
	if productIDArrReq != "" {
		err = json.NewDecoder(strings.NewReader(productIDArrReq)).Decode(&productIDArr)
	}

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	data, err := pc.ProductRepository.GeneratePriceTemplate(productIDArr)

	xlsx := excelize.NewFile()
	xlsx.NewSheet("Single Product Prices")
	xlsx.SetCellValue("Single Product Prices", "A1", "Single Product ID")
	xlsx.SetCellValue("Single Product Prices", "B1", "SKU")
	xlsx.SetCellValue("Single Product Prices", "C1", "Nama Harga")
	xlsx.SetCellValue("Single Product Prices", "D1", "Nilai Harga")

	for index, singleProduct := range data.SingleProductPricesTemplate {
		xlsx.SetCellValue("Single Product Prices", fmt.Sprintf("A%d", index+2), singleProduct.SingleProductID)
		xlsx.SetCellValue("Single Product Prices", fmt.Sprintf("B%d", index+2), singleProduct.Sku)
		xlsx.SetCellValue("Single Product Prices", fmt.Sprintf("C%d", index+2), singleProduct.PriceName)
		xlsx.SetCellValue("Single Product Prices", fmt.Sprintf("D%d", index+2), singleProduct.PriceValue)

	}

	xlsx.NewSheet("Variant Product Prices")
	xlsx.SetCellValue("Variant Product Prices", "A1", "Variant Product ID")
	xlsx.SetCellValue("Variant Product Prices", "B1", "SKU")
	xlsx.SetCellValue("Variant Product Prices", "C1", "Nama Harga")
	xlsx.SetCellValue("Variant Product Prices", "D1", "Nilai Harga")

	for index, variantProduct := range data.VariantProductPriceTemplate {
		xlsx.SetCellValue("Variant Product Prices", fmt.Sprintf("A%d", index+2), variantProduct.VariantProductID)
		xlsx.SetCellValue("Variant Product Prices", fmt.Sprintf("B%d", index+2), variantProduct.Sku)
		xlsx.SetCellValue("Variant Product Prices", fmt.Sprintf("C%d", index+2), variantProduct.PriceName)
		xlsx.SetCellValue("Variant Product Prices", fmt.Sprintf("D%d", index+2), variantProduct.PriceValue)

	}

	var b bytes.Buffer
	writr := bufio.NewWriter(&b)

	if err := xlsx.Write(writr); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := c.Response()
	header := res.Header()
	header.Set(echo.HeaderContentType, echo.MIMEOctetStream)
	header.Set(echo.HeaderContentDisposition, "attachment;filename=price.xlsx")
	header.Set("Content-Transfer-Encoding", "binary")
	header.Set("Expires", "0")
	res.WriteHeader(http.StatusOK)
	return c.Blob(http.StatusOK, echo.MIMEOctetStream, b.Bytes())

}

func (pc *ProductController) GetProductKinds(c echo.Context) error {
	data, err := pc.ProductRepository.FetchAllKind()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res := &httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    data,
	}
	return c.JSON(res.Status, res)
}

func (pc *ProductController) GetProductTypes(c echo.Context) error {
	data, err := pc.ProductRepository.FetchAllType()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res := &httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    data,
	}
	return c.JSON(res.Status, res)
}
