package productcontroller

import (
	"net/http"

	"distrodakwah_backend/app/helper/httphelper"
	"distrodakwah_backend/app/services/handler/producthandler"

	"github.com/labstack/echo"
)

func (pc *ProductController) GetBrands(c echo.Context) error {
	data, err := pc.ProductRepository.FetchAllBrand()
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

func (pc *ProductController) PostBrand(c echo.Context) error {
	brandReq := &producthandler.BrandReq{}
	if err := c.Bind(&brandReq); err != nil {
		return err
	}

	err := pc.ProductRepository.CreateBrand(brandReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)

}
