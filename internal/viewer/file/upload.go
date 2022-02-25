package file

import (
	"file-manager/internal/viewer/breadcrumb"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"os"
)

type UploadDTO struct {
	CSRF       string
	Message    string
	Breadcrumb breadcrumb.Breadcrumb
}

func Upload(dir string, c *gin.Context) {
	dto := TextDTO{
		CSRF:       csrf.GetToken(c),
		Breadcrumb: breadcrumb.NewBreadcrumb(dir),
	}

	if form, err := c.MultipartForm(); err == nil {
		files := form.File["file[]"]

		for _, file := range files {
			dst := dir + string(os.PathSeparator) + file.Filename

			err = c.SaveUploadedFile(file, dst)
			if err != nil {
				dto.Message = err.Error()
			} else {
				dto.Message = "Успешно загружено"
			}
		}
	}

	c.HTML(http.StatusOK, "upload-file.html", dto)
}
