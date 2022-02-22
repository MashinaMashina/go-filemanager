package file

import (
	"file-manager/internal/viewer/breadcrumb"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"path/filepath"
)

type PictureDTO struct {
	Dir        string
	Name       string
	Src        *url.URL
	Breadcrumb breadcrumb.Breadcrumb
}

func viewPictureFile(path string, c *gin.Context) {
	if _, exists := c.GetQuery("view"); exists {
		proxyFile(path, c)
		return
	}

	url := c.Request.URL
	q := url.Query()
	q.Add("view", "true")
	url.RawQuery = q.Encode()

	dto := PictureDTO{
		Src: url,
	}

	dto.Dir, dto.Name = filepath.Split(path)
	dto.Breadcrumb = breadcrumb.NewBreadcrumb(dto.Dir)

	c.HTML(http.StatusOK, "view-image.html", dto)
}
