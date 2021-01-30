package inventorycontroller

import (
	"distrodakwah_backend/app/helper/httphelper"
	"distrodakwah_backend/app/services/library/inventorylibrary"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (ic *InventoryController) GetProvinces(c echo.Context) error {
	provinces, err := inventorylibrary.GetProvinces()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resp := httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    provinces,
	}
	return c.JSON(resp.Status, resp)
}

func (ic *InventoryController) GetCitiesByProvinceID(c echo.Context) error {
	provinceIDStr := c.Param("province_id")
	provinceID, _ := strconv.Atoi(provinceIDStr)
	provinces, err := inventorylibrary.GetCitiesByProvinceID(provinceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resp := httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    provinces,
	}
	return c.JSON(resp.Status, resp)
}

func (ic *InventoryController) GetSubsByCityID(c echo.Context) error {
	CityIDStr := c.Param("city_id")
	CityID, _ := strconv.Atoi(CityIDStr)
	cities, err := inventorylibrary.GetSubdistrictsByCityID(CityID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resp := httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    cities,
	}
	return c.JSON(resp.Status, resp)
}
