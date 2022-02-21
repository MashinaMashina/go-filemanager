package errorpage

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Page(error string, c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", gin.H{
		"Message": error,
	})
}
