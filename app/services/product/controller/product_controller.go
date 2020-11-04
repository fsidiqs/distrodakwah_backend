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
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/repository"
)

type ProductController struct {
	ProductRepository *repository.ProductRepository
}

func (pc *ProductController) Gets(c echo.Context) error {
	pageReq, err := strconv.Atoi(c.QueryParam("p_page"))
	limitReq, err := strconv.Atoi(c.QueryParam("p_limit"))
	preloadReq := c.QueryParam("preload")
	productIDArrReq := c.QueryParam("product_id_arr")

	request := &repository.FetchAllReq{
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

	// for _, prices := range product.ProductDetail.Prices {
	// 	fmt.Printf("%+v \n", *prices)

	// }
	// fmt.Printf("%+v \n", product)
	// fmt.Printf("%+v \n", product.ProductDetail.Pri)
	// req, err := ioutil.ReadAll(c.Request().Body)

	// err = json.Unmarshal(req, &product)
	// if err != nil {
	// 	fmt.Printf("error parsing %+v \n", err)
	// 	return nil
	// }
	// fmt.Printf("%+v \n", product)
	// for _, image := range product.ProductDetail {
	// 	fmt.Printf("%+v \n", *image)
	// }
	// fmt.Printf("%T \n", *product.ProductImages[0])
	// json.Unmarshal(&produ)
	// fmt.Printf("%+v \n", product)

	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	// }

	// product, err = pc.ProductRepository.SaveProduct(product)
	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)
}

func (pc *ProductController) CreateProductBasicStructure(c echo.Context) error {
	product := &model.ProductFromRequestJSON{}
	if err := c.Bind(&product); err != nil {
		fmt.Printf("Error: %+v", err)
	}

	err := pc.ProductRepository.SaveProductBasicStructure(product)
	if err != nil {
		fmt.Printf("error creating product: %+v", err)
		return c.JSON(http.StatusInternalServerError, httphelper.InternalServerErrorMessage)
	}

	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)
}
