package main

import (
	"github.com/gin-gonic/gin"
	"github.com/namle133/hash_password1.git/hash_password/database"
	"github.com/namle133/hash_password1.git/hash_password/http/decode"
	"github.com/namle133/hash_password1.git/hash_password/http/encode"
	"github.com/namle133/hash_password1.git/hash_password/service"
)

func main() {
	r := gin.Default()
	p := &service.Product{Db: database.ConnectDatabase()}
	var i service.IProduct = p
	r.POST("/signup", func(c *gin.Context) {
		user := decode.SignRequest(c)
		i.SignUp(user)
		encode.SignUpResponse(c)
	})

	r.POST("/signin", func(c *gin.Context) {
		user := decode.SignRequest(c)
		i.SignIn(user, c)
	})

	r.GET("/welcome", func(c *gin.Context) {
		claims := i.Welcome(c)
		encode.WelcomeResponse(c, claims)
	})

	r.POST("/refresh", func(c *gin.Context) {
		i.Refresh(c)
	})

	r.PUT("/password", func(c *gin.Context) {
		us, pw := decode.NewPasswordRequest(c)
		i.NewPasswordToken(us, pw, c)
	})

	r.Run(":8000")
}
