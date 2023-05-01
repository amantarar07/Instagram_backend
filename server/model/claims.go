package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Type string `json:"type"`
	Id   string
	Role string
	PhoneNumber string `json:"phoneNumber"`
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}
