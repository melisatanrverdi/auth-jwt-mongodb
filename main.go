package main

import (
	middleware "github.com/melisatanrverdi/auth-jwt-mongodb/middleware"
	routes "github.com/melisatanrverdi/auth-jwt-mongodb/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	// API
	router.GET("/api", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api"})

	})

	//p := os.ExpandEnv("${PORT}")
	router.Run(":" + "8080")
}
