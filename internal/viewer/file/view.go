package file

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func View(path string, c *gin.Context) {
	ext := filepath.Ext(path)

	if containExt(ext, getTextExtensions()) {
		viewTextFile(path, c)
	} else if containExt(ext, getPictureExtensions()) {
		viewPictureFile(path, c)
	} else {
		viewUnknownFile(path, c)
	}
}
