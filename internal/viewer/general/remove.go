package general

import (
	"file-manager/internal/errorpage"
	"file-manager/internal/viewer/links"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func Remove(path string, c *gin.Context) {
	if c.Request.Method != "POST" {
		errorpage.Page("Неверный тип запроса", c)
		return
	}

	dir, _ := filepath.Split(path)

	err := os.RemoveAll(path)
	if err != nil {
		errorpage.Page("Не удается удалить "+path+". Ошибка: "+err.Error(), c)
		return
	}

	c.Redirect(http.StatusFound, links.View(dir).String())
}
