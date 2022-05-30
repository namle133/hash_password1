package domain

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Credentials struct {
	gorm.Model
	Username string `json:"username"`
	Password []byte `json:"password"`
}

var JwtKey = []byte("my-secrect-key")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
