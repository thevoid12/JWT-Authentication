package main

import (
	"fmt"
	constant "jwtAuth/constants"
	routes "jwtAuth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//port := os.Getenv("PORT")
	port := constant.GetEnvConstFromViper("PORT")
	if port == "" {
		port = "8070"
	}
	fmt.Println(port)

	//creating a router
	router := gin.New()
	// Use attaches a global middleware to the router. i.e.
	// the middleware attached through Use() will be included in the handlers chain for every single request
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	// routes.UserRoutes(router)

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
