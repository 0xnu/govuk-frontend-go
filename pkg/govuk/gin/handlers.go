package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthz() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	}
}
