package transfer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}