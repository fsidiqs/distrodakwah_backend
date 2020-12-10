package usercontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"distrodakwah_backend/app/helper/httphelper"
	"distrodakwah_backend/app/services/handler/userhandler"
	"distrodakwah_backend/app/services/model/usermodel"
	"distrodakwah_backend/app/services/repository/userrepository"

	"github.com/labstack/echo"
)

type Controller struct {
	UserRepository *userrepository.Repository
}

func (uc *Controller) CreateUser(c echo.Context) error {
	var err error
	userReq := userhandler.UserReq{}
	err = json.NewDecoder(c.Request().Body).Decode(&userReq)
	if err != nil {

		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)

	}
	userCreate, err := usermodel.NewUser(userReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// create user using repo
	user, err := uc.UserRepository.CreateUser(*userCreate)
	resp := httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    user,
	}
	return c.JSON(resp.Status, resp)

}

func (uc *Controller) CreateUserReseller(c echo.Context) error {
	var err error

	userResellerReq := userhandler.UserResellerReq{}
	err = json.NewDecoder(c.Request().Body).Decode(&userResellerReq)
	if err != nil {
		fmt.Printf("test %+v \n", err)

		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}

	userResellerCreate, err := usermodel.NewUserReseller(userResellerReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	userReseller, err := uc.UserRepository.CreateUserReseller(*userResellerCreate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    userReseller,
	}

	return c.JSON(resp.Status, resp)
}
