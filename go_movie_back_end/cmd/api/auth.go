package main

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct{

	Issuer string
	Audience string
	Secrect string 
	TokenExpiry time.Duration
	RefreshExpiry time.Duration
	CookieDomain string
CookiePath string 
CookieName string

}

type jwtUser struct{
ID int `json:"id"`
FirstName string `json:"first_name"`
LastName string `json:"last_name"`
}

type TokenPairs struct{
	Token int `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	}

type Claims struct{
		jwt.RegisteredClaims
	}