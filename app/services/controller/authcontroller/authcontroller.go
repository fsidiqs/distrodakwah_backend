package authcontroller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/auth"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httphelper"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/model/usermodel"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/repository/authrepository"
)

type Controller struct {
	AuthRepository *authrepository.Repository
}

func (uc *Controller) Login(c echo.Context) error {
	cred := auth.LoginCredetials{}
	var err error
	if err = c.Bind(&cred); err != nil {
		return c.JSON(http.StatusBadRequest, httphelper.BadRequestMessage)
	}
	user := &usermodel.User{}
	user, _ = uc.AuthRepository.GetUserByEmail(cred.Email)
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
