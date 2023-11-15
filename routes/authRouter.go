package routes

import (
	"khvnkay/golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/login", controllers.Login())
	incomingRoutes.POST("user/signup", controllers.SignUp())

}
