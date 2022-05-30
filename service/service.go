package service

import (
	"github.com/gin-gonic/gin"
	"github.com/namle133/hash_password1.git/hash_password/domain"
)

type IProduct interface {
	SignUp(creds *domain.User)
	SignIn(creds *domain.User, c *gin.Context)
	Welcome(c *gin.Context) *domain.Claims
	Refresh(c *gin.Context)
	NewPasswordToken(us string, pw string, c *gin.Context)
}
