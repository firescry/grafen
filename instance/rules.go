package instance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Rules []Rule

type Rule struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func GetRules(c *gin.Context) {
	rules := Rules{Rule{}}
	c.JSON(http.StatusOK, rules)
}
