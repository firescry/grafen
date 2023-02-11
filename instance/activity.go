package instance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Activity []WeekActivity

type WeekActivity struct {
	Week          string `json:"week"`
	Statuses      string `json:"statuses"`
	Logins        string `json:"logins"`
	Registrations string `json:"registrations"`
}

func GetActivity(c *gin.Context) {
	activity := Activity{WeekActivity{}}
	c.JSON(http.StatusOK, activity)
}
