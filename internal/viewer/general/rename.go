package general

import (
	"file-manager/internal/viewer/breadcrumb"
	"file-manager/internal/viewer/links"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"os"
	"path/filepath"
)

type RenameDTO struct {
	Name       string
	CSRF       string
	Message    string
	Breadcrumb breadcrumb.Breadcrumb
}

func Rename(path string, c *gin.Context) {
	dto := RenameDTO{
		CSRF: csrf.GetToken(c),
	}

	var dir string
	dir, dto.Name = filepath.Split(path)
	dto.Breadcrumb = breadcrumb.NewBreadcrumb(dir)

	if name, exists := c.GetPostForm("name"); exists {
		if err := os.Rename(path, dir+string(os.PathSeparator)+name); err != nil {
			dto.Message = err.Error()
		} else {
			c.Redirect(http.StatusFound, links.View(dir).String())
			return
		}
	}

	c.HTML(http.StatusOK, "rename.html", dto)
}
