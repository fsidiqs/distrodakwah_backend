package productcontroller

import (
	"net/http"

	"distrodakwah_backend/app/helper/httphelper"
	"distrodakwah_backend/app/services/handler/producthandler"

	"github.com/labstack/echo"
)

func (pc *ProductController) GetCategories(c echo.Context) error {
	data, err := pc.ProductRepository.FetchAllCategory()
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

func (pc *ProductController) PostCategory(c echo.Context) error {
	categoryReq := &producthandler.CategoryReq{}
	if err := c.Bind(&categoryReq); err != nil {
		return err
	}

	err := pc.ProductRepository.CreateCategory(categoryReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)

}

func (pc *ProductController) GetDepartments(c echo.Context) error {
	data, err := pc.ProductRepository.FetchAllDepartments()

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

func (pc *ProductController) PostDepartment(c echo.Context) error {
	departmentReq := &producthandler.DepartmentReq{}
	if err := c.Bind(&departmentReq); err != nil {
		return err
	}

	err := pc.ProductRepository.CreateDepartment(departmentReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)
}

func (pc *ProductController) GetSubdepartments(c echo.Context) error {
	data, err := pc.ProductRepository.FetchAllSubdepartments()

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

func (pc *ProductController) PostSubdepartment(c echo.Context) error {
	subdepartmentReq := &producthandler.SubdepartmentReq{}
	if err := c.Bind(&subdepartmentReq); err != nil {
		return err
	}

	err := pc.ProductRepository.CreateSubdepartment(subdepartmentReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, httphelper.StatusOKMessage)
}
