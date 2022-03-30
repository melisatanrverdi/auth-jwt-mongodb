package routes

import (
	controller "github.com/melisatanrverdi/auth-jwt-mongodb/controllers"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.Signup)
	incomingRoutes.POST("/users/login", controller.Login)
}
