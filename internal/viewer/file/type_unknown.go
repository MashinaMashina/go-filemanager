package file

import (
	"github.com/gin-gonic/gin"
)

func viewUnknownFile(path string, c *gin.Context) {
	DownloadFile(path, c)
}
