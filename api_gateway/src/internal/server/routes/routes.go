package routes

import (
	"carsales/internal/controller"
	"carsales/middleware/cors"

	"github.com/gin-gonic/gin"
)

func Get() *gin.Engine {
	r := gin.Default()

	corsCfg := cors.Get()
	r.Use(corsCfg)

	authController := new(controller.AuthController)
	auth := r.Group("api/auth")
	{
		auth.POST("/signUp", authController.SignUp)
		auth.POST("/signIn", authController.SignIn)
	}

	return r
}