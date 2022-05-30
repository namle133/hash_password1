package encode

import (
	"github.com/gin-gonic/gin"
	"github.com/namle133/hash_password1.git/hash_password/domain"
	"net/http"
)

func SignUpResponse(c *gin.Context) {
	c.String(http.StatusOK, "%s", "SignUp Successfully!")
}

func WelcomeResponse(c *gin.Context, claims *domain.Claims) {
	c.String(http.StatusOK, "Welcome to  %v", claims.Username)
}
