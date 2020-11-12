package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	jwtReq "github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/user/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/config"
)

type EmailContext struct {
	echo.Context
	Email string
}

type UserMiddleware struct {
	Email  string
	RoleID uint8
}

type UserContext struct {
	echo.Context
	User UserMiddleware
}

func CheckAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := jwtReq.ParseFromRequestWithClaims(
			c.Request(),
			jwtReq.OAuth2Extractor,
			&model.Claim{},
			func(t *jwt.Token) (interface{}, error) {
				return config.JWTSECRET, nil
			},
		)
		if err != nil {
			fmt.Printf("Failed to parse token")
			return c.JSON(http.StatusInternalServerError, "failed to parse token")
		}
		if !token.Valid {
			fmt.Printf("Invalid token")
			return c.JSON(http.StatusInternalServerError, "token is invalid")
		}
		emailContext := &EmailContext{Context: c, Email: token.Claims.(*model.Claim).User.Email}

		return next(emailContext)
	}
}

const (
	isAdmin = 1 << iota
	isReseller
)

func AdminRoleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		emailContext := c.(*EmailContext)
		userContext := &UserContext{Context: c, User: UserMiddleware{}}
		err := database.DB.Model(&model.User{}).Where("email = ?", emailContext.Email).Find(&userContext.User).Error
		if err != nil {
			return err
		}

		var roles byte = isAdmin

		if userContext.User.RoleID&roles != isAdmin {
			return c.JSON(http.StatusInternalServerError, "is not admin")
		}
		return next(userContext)
	}
}
