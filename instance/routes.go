package instance

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/api/v1/instance", getInstance)
	r.GET("/api/v1/instance/activity", getActivity)
	r.GET("/api/v1/instance/domain_blocks", getBlocks)
	r.GET("/api/v1/instance/extended_description", getDescription)
	r.GET("/api/v1/instance/peers", getPeers)
	r.GET("/api/v1/instance/rules", getRules)
}
