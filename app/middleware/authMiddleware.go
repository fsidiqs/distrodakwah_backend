package middleware

import (
	"fmt"
	"net/http"

	"distrodakwah_backend/app/auth"
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/model/usermodel"
	"distrodakwah_backend/config"

	"github.com/dgrijalva/jwt-go"
	jwtReq "github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo"
)

type EmailContext struct {
	echo.Context
	Email string
}

type UserMiddleware struct {
	ID     uint64
	Email  string
	RoleID uint8
}

type UserContext struct {
	echo.Context
	User UserMiddleware
}

func performAuthCheking(c echo.Context) (*jwt.Token, error) {
	token, err := jwtReq.ParseFromRequestWithClaims(
		c.Request(),
		jwtReq.OAuth2Extractor,
		&auth.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return config.JWTSECRET, nil
		},
	)
	return token, err
}

func CheckAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := performAuthCheking(c)
		if err != nil {
			fmt.Printf("Failed to parse token")
			return c.JSON(http.StatusInternalServerError, "unauthenticated")
		}
		if !token.Valid {
			fmt.Printf("Invalid token")
			return c.JSON(http.StatusInternalServerError, "token is invalid")
		}

		emailContext := &EmailContext{Context: c, Email: token.Claims.(*auth.Claim).User.Email}

		return next(emailContext)
	}
}

const (
	isAdmin = 1 << iota
	isReseller
)

func AdminRoleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// emailContext := c.(*EmailContext)

		token, err := performAuthCheking(c)
		if err != nil {
			fmt.Printf("Failed to parse token")
			return c.JSON(http.StatusInternalServerError, "unauthenticated")
		}
		if !token.Valid {
			fmt.Printf("Invalid token")
			return c.JSON(http.StatusInternalServerError, "token is invalid")
		}

		emailContext := &EmailContext{Context: c, Email: token.Claims.(*auth.Claim).User.Email}

		userContext := &UserContext{Context: c, User: UserMiddleware{}}
		err = database.DB.Model(&usermodel.User{}).Where("email = ?", emailContext.Email).Find(&userContext.User).Error
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
