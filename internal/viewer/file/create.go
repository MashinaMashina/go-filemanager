package file

import (
	"errors"
	"file-manager/internal/viewer/breadcrumb"
	"fmt"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"io/ioutil"
	"net/http"
	"os"
)

func Create(dir string, c *gin.Context) {
	dto := TextDTO{
		CSRF:       csrf.GetToken(c),
		Dir:        dir,
		Exists:     false,
		Breadcrumb: breadcrumb.NewBreadcrumb(dir),
	}

	if name := c.DefaultPostForm("name", ""); name != "" {
		path := dir + string(os.PathSeparator) + name
		content := c.DefaultPostForm("content", "")

		dto.Name = name
		dto.Content = content

		var fileErr error
		if _, fileErr = os.Stat(path); errors.Is(fileErr, os.ErrNotExist) {
			err := ioutil.WriteFile(path, []byte(content), 0760)
			if err != nil {
				dto.Message = "Ошибка создания файла: " + err.Error()
			} else {
				dto.Message = "Сохранено"
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

	c.HTML(http.StatusOK, "view-text.html", dto)
}
