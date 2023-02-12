package main

import (
	"github.com/firescry/grafen/instance"
	"github.com/firescry/grafen/webfinger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"content-type"},
	}
	router.Use(cors.New(config))

	webfinger.RegisterRoutes(router)
	instance.RegisterRoutes(router)

	router.Run("localhost:8080")
}
