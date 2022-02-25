package file

import (
	"file-manager/internal/errorpage"
	"file-manager/internal/filesystem"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
)

func DownloadFile(path string, c *gin.Context) {
	file, err := filesystem.NewFileFromPath(path)

	if err != nil {
		errorpage.Page("Не удается открыть файл "+path, c)
		return
	}

	fReader, err := os.Open(path)

	if err != nil {
		errorpage.Page("Не удается открыть файл "+path, c)
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+file.Name())
	c.Header("Content-Length", strconv.FormatInt(file.Size(), 10))

	io.CopyN(c.Writer, fReader, file.Size())
}
