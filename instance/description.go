package instance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Description struct {
	UpdatedAt string `json:"updated_at"`
	Content   string `json:"content"`
}

func getDescription(c *gin.Context) {
	description := Description{}
	c.JSON(http.StatusOK, description)
}
