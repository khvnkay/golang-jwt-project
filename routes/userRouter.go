package routes

import (
	"khvnkay/golang-jwt-project/controllers"
	"khvnkay/golang-jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/user", controllers.GetUsers())
	incomingRoutes.GET("/user/:user_id", controllers.GetUser())

}
