package file

import (
	"file-manager/internal/filesystem"
	"file-manager/internal/viewer/breadcrumb"
	"file-manager/internal/viewer/links"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
)

type TextDTO struct {
	Dir          string
	Name         string
	Content      string
	CSRF         string
	Message      string
	Exists       bool
	Breadcrumb   breadcrumb.Breadcrumb
	DownloadLink *url.URL
	RenameLink   *url.URL
}

func viewTextFile(path string, c *gin.Context) {
	dto := TextDTO{
		CSRF:         csrf.GetToken(c),
		Exists:       true,
		DownloadLink: links.DownloadFile(path),
		RenameLink:   links.Rename(path),
	}

	f, err := filesystem.NewFileFromPath(path)

	if err != nil {
		return
	}

	if content, exists := c.GetPostForm("content"); exists {
		err = ioutil.WriteFile(path, []byte(content), f.Mode())
		if err != nil {
			dto.Message = err.Error()
		} else {
			dto.Message = "Сохранено"
		}
	}

	dto.Dir, dto.Name = filepath.Split(path)
	dto.Breadcrumb = breadcrumb.NewBreadcrumb(dto.Dir)

	content, err := ioutil.ReadFile(path)

	if err != nil {
		dto.Message = err.Error()
	}

	dto.Content = string(content)

	c.HTML(http.StatusOK, "view-text.html", dto)
}
