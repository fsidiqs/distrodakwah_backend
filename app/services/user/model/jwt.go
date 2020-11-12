package model

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	User *PublicUser `json:"user"`
	jwt.StandardClaims
}
