package webfinger

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Users = []string{"test1", "test2"}

func GetWebfinger(c *gin.Context) {
	query := c.Request.URL.Query()
	resource, ok := query["resource"]

	if !ok || len(resource) > 1 || resource[0] == "" {
		// "resource" parameter is absent,
		// provided more than once
		// or provided without value
		c.String(http.StatusBadRequest, "")
		return
	}

	for _, user := range Users {
		if resource[0] == user {
			// resource found
			c.JSON(http.StatusOK, gin.H{"resource": "found"})
			return
		}
	}
	// resource not found
	c.String(http.StatusNotFound, "")
}
