package instance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Peers []string

func GetPeers(c *gin.Context) {
	peers := Peers{}
	c.JSON(http.StatusOK, peers)
}
