package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/user/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/config"
	"golang.org/x/crypto/bcrypt"
)

// Hash make a password hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword verify the hashed password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateJWT(user *model.PublicUser) (string, error) {
	claim := model.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "distrodakwah.id",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(config.JWTSECRET)
}
