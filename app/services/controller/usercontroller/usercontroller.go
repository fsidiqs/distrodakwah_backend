package usercontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"distrodakwah_backend/app/helper/httphelper"
	"distrodakwah_backend/app/helper/pagination"
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

// get user vendor using repo

func (uc *Controller) GetUserVendors(c echo.Context) error {
	var err error

	userVendorReq := usermodel.UserWithChild{}

	userVendor, err := uc.UserRepository.GetUserVendor(userVendorReq)

	fmt.Println("USERVENDOR", userVendor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	resp := httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    userVendor,
	}

	return c.JSON(resp.Status, resp)
}

// get user reseller using repo
func (uc *Controller) GetUserResellers(c echo.Context) error {
	var err error

	userResellerReq := []usermodel.UserReseller{}

	userReseller, err := uc.UserRepository.GetUserReseller(userResellerReq)

	fmt.Println("USERRESELLER", userReseller)
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

// get all user
// func (uc *Controller) GetAllUser(c echo.Context) error {
// 	var err error

// 	usersReq := []usermodel.User{}

// 	users, err := uc.UserRepository.GetUsers(usersReq)

// 	fmt.Println("USERS", users)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}
// 	resp := httphelper.Response{
// 		Status:  http.StatusOK,
// 		Message: httphelper.StatusOKMessage,
// 		Data:    users,
// 	}

// 	return c.JSON(resp.Status, resp)
// }

func (uc *Controller) GetAllUser(c echo.Context) error {
	pageReq, err := strconv.Atoi(c.QueryParam("u_page"))
	limitReq, err := strconv.Atoi(c.QueryParam("u_limit"))
	userIDarrReq := c.QueryParam("user_id_arr")

	request := &userhandler.FetchUsersReq{
		UserIDarr: []int{},
		Metadata: pagination.Metadata{
			Page:  pageReq,
			Limit: limitReq,
		},
	}

	if userIDarrReq != "" {
		err = json.NewDecoder(strings.NewReader(userIDarrReq)).Decode(&request.UserIDarr)
	}
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "page_and_limit_empty")
	}
	data, err := uc.UserRepository.GetUsers(request)
	res := &httphelper.Response{
		Status:  http.StatusOK,
		Message: httphelper.StatusOKMessage,
		Data:    data,
	}
	return c.JSON(res.Status, res)
}

// Delete User
func (uc *Controller) DeleteUsers(c echo.Context) error {

	id := c.Param("id")
	fmt.Println("ID", id)
	idConv, _ := strconv.Atoi(id)

	err := uc.UserRepository.DeleteUser(idConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	resp := httphelper.Response{
		Status:  http.StatusOK,
		Message: "user id" + " " + id + " " + "succes deleted!",
	}

	return c.JSON(resp.Status, resp)
}

//upgrade user reseller to ekslusif
func (uc *Controller) UpgradeAkunReseller(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("ID", id)
	idConv, _ := strconv.Atoi(id)

	err := uc.UserRepository.UpgradeAkun(idConv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := httphelper.Response{
		Status:  http.StatusOK,
		Message: "user id" + " " + id + " " + "succes upgrade to ekslusif",
	}

	return c.JSON(resp.Status, resp)

}
