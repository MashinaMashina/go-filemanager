package file

import (
	"file-manager/internal/errorpage"
	"file-manager/internal/filesystem"
	"github.com/gin-gonic/gin"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strconv"
)

func proxyFile(path string, c *gin.Context) {
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

	cType := mime.TypeByExtension(filepath.Ext(path))

	c.Header("Content-Type", cType)
	c.Header("Content-Length", strconv.FormatInt(file.Size(), 10))

	io.CopyN(c.Writer, fReader, file.Size())
}
