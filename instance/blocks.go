package instance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Blocks []Block

type Block struct {
	Domain   string `json:"domain"`
	Digest   string `json:"digest"`
	Severity string `json:"severity"`
	Comment  string `json:"comment"`
}

func getBlocks(c *gin.Context) {
	blocks := Blocks{Block{}}
	c.JSON(http.StatusOK, blocks)
}
