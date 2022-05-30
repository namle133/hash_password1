package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/namle133/hash_password1.git/hash_password/domain"
	"github.com/namle133/hash_password1.git/hash_password/hash_password"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Product struct {
	Db *gorm.DB
}

func (p *Product) SignUp(creds *domain.User) {
	creds1 := &domain.Credentials{Username: creds.Username, Password: hash_password.Hash(creds.Password)}
	p.Db.Create(&creds1)
}

func (p *Product) SignIn(creds *domain.User, c *gin.Context) {
	var creds1 *domain.Credentials
	e := p.Db.First(&creds1, "username=?", creds.Username).Scan(&creds1).Error
	if e != nil {
		c.String(http.StatusBadRequest, "%v", e)
		return
	}
	er := bcrypt.CompareHashAndPassword(creds1.Password, []byte(creds.Password))
	if er != nil {
		c.String(http.StatusBadRequest, "%v", e)
		return
	} else {
		fmt.Println("password are equal")
	}
	expirationTime := time.Now().Add(3 * time.Minute)
	claims := domain.Claims{
		Username: creds.Username,
		Password: creds.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(domain.JwtKey)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}

	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
}

func (p *Product) Welcome(c *gin.Context) *domain.Claims {
	ck, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.String(http.StatusUnauthorized, "%v", err)
			return nil
		}
		c.String(http.StatusBadRequest, "%v", err)
		return nil
	}
	fmt.Println(ck)
	claims := &domain.Claims{}
	tknStr := ck

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return domain.JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.String(http.StatusUnauthorized, "%v", err)
			return nil
		}
		c.String(http.StatusBadRequest, "%v", err)
		return nil
	}

	if !tkn.Valid {
		c.String(http.StatusUnauthorized, "%v", tkn)
		return nil
	}
	return claims
}

func (p *Product) Refresh(c *gin.Context) {
	ck, err := c.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			c.String(http.StatusUnauthorized, "%v", err)
			return
		}
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	claims := &domain.Claims{}
	tknStr := ck

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return domain.JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.String(http.StatusUnauthorized, "%v", err)
			return
		}
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	if !tkn.Valid {
		c.String(http.StatusUnauthorized, "%v", tkn)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Minute {
		c.String(http.StatusBadRequest, "%v", 400)
		return
	}

	expirationTime := time.Now().Add(3 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(domain.JwtKey)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}

	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
}

func (p *Product) NewPasswordToken(us string, pw string, c *gin.Context) {

	ck, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.String(http.StatusUnauthorized, "%v", err)

			return
		}
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	claims := &domain.Claims{}
	tknStr := ck

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return domain.JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.String(http.StatusUnauthorized, "%v", err)
			return
		}
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	if !tkn.Valid {
		c.String(http.StatusUnauthorized, "%v", tkn)
		return
	}
	expirationTime := time.Now().Add(3 * time.Minute)
	var item *domain.Credentials
	er := p.Db.Model(&item).Where("username = ?", us).Update("password", hash_password.Hash(pw)).Error
	if er != nil {
		c.String(http.StatusBadRequest, "%v", er)
		return
	}
	claims.Password = pw
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(domain.JwtKey)
	if err != nil {
		c.String(http.StatusInternalServerError, "%v", err)
		return
	}

	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)

}
