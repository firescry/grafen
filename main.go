package main

import (
	"github.com/firescry/grafen/webfinger"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/.well-known/webfinger", webfinger.GetWebfinger)
	router.Run(":8080")
}
