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
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/repository"
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
		Message: httphelper.StatusOkMessage,
		Data:    data,
	}
	return c.JSON(res.Status, res)
}
