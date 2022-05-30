package decode

import (
	"github.com/gin-gonic/gin"
	"github.com/namle133/hash_password1.git/hash_password/domain"
	"net/http"
)

func SignRequest(c *gin.Context) *domain.User {
	var creds *domain.User
	err := c.BindJSON(&creds)
	if err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return nil
	}
	return creds
}

func NewPasswordRequest(c *gin.Context) (string, string) {
	pw := c.Query("password")
	us := c.Query("username")
	return us, pw
}
