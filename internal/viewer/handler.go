package viewer

import "github.com/gin-gonic/gin"

func RegisterHandler(r *gin.Engine) {
	r.GET("/", openPath)
	r.POST("/", openPath)
}
