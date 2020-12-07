package usercontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/auth"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httphelper"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/handler/userhandler"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/model/usermodel"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/repository/userrepository"
)

type Controller struct {
	UserRepository *userrepository.Repository
}

func (uc *Controller) Login(c echo.Context) error {
	cred := auth.LoginCredetials{}
	var err error
	if err = c.Bind(&cred); err != nil {
		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}
	user := &usermodel.User{}
	user, _ = uc.UserRepository.GetUserByEmail(cred.Email)
	err = auth.VerifyPassword(user.Password, cred.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "login failed")
	}
	publicUser := &auth.CredUser{
		Email: user.Email,
	}
	token, err := auth.GenerateJWT(publicUser)
	if err != nil {
		log.Println("Error Creating JWT token", err)
		return c.JSON(http.StatusInternalServerError, "Something Went Wrong")
	}

	resp := httphelper.Response{
		Message: "you are logged in",
		Status:  http.StatusOK,
		Data:    token,
	}
	return c.JSON(http.StatusOK, resp)
}

func (uc *Controller) CreateUser(c echo.Context) error {
	userReq := userhandler.UserReq{}
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}
	userCreate, err := usermodel.NewUser(userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Printf("fajar: %+v \n", userCreate)
	return nil
	// user, err := uc.UserRepository.CreateUser()

}
