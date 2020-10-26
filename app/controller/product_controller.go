package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httpHelper"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/repository"
)

type ProductController struct {
	ProductRepository *repository.ProductRepository
}

func (pc *ProductController) Gets(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("p_page"))
	limit, err := strconv.Atoi(c.QueryParam("p_limit"))

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "page_and_limit_empty")
	}

	meta := &pagination.Metadata{
		Page:  page,
		Limit: limit,
	}

	data, err := pc.ProductRepository.FetchAll(meta)
	res := &httpHelper.Response{
		Status:  http.StatusOK,
		Message: httpHelper.SuccessMessage,
		Data:    data,
	}
	return c.JSON(res.Status, res)
}
