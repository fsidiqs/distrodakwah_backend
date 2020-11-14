package model

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	User *CredUser `json:"user"`
	jwt.StandardClaims
}
