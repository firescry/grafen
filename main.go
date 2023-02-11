package main

import (
	"github.com/firescry/grafen/instance"
	"github.com/firescry/grafen/webfinger"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/.well-known/webfinger", webfinger.GetWebfinger)

	router.GET("/api/v1/instance/activity", instance.GetActivity)
	router.GET("/api/v1/instance/domain_blocks", instance.GetBlocks)
	router.GET("/api/v1/instance/extended_description", instance.GetDescription)
	router.GET("/api/v1/instance/peers", instance.GetPeers)
	router.GET("/api/v1/instance/rules", instance.GetRules)

	router.Run("localhost:8080")
}
