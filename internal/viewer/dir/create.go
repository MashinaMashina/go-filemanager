package dir

import (
	"errors"
	"file-manager/internal/viewer/breadcrumb"
	"file-manager/internal/viewer/links"
	"fmt"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"os"
)

type CreateDTO struct {
	CSRF       string
	Message    string
	Name       string
	Breadcrumb breadcrumb.Breadcrumb
}

func Create(dir string, c *gin.Context) {
	dto := CreateDTO{
		CSRF:       csrf.GetToken(c),
		Breadcrumb: breadcrumb.NewBreadcrumb(dir),
	}

	if name := c.DefaultPostForm("name", ""); name != "" {
		path := dir + string(os.PathSeparator) + name

		dto.Name = name
		var fileErr error
		if _, fileErr = os.Stat(path); errors.Is(fileErr, os.ErrNotExist) {
			err := os.MkdirAll(path, 0760)
			if err != nil {
				dto.Message = "Ошибка создания папки: " + err.Error()
			} else {
				c.Redirect(http.StatusFound, links.View(dir).String())
				return
			}
		} else {
			if fileErr == nil {
				dto.Message = "Ошибка создания файла: Файл уже существует"
			} else {
				fmt.Println(fileErr.Error())
				dto.Message = fileErr.Error()
			}
		}
	}

	c.HTML(http.StatusOK, "create-dir.html", dto)
}
