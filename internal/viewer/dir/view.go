package dir

import (
	"file-manager/internal/errorpage"
	"file-manager/internal/filesystem"
	"file-manager/internal/viewer/breadcrumb"
	"file-manager/internal/viewer/links"
	"fmt"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"io/ioutil"
	"net/http"
	"net/url"
)

type viewFile struct {
	filesystem.File
	csrf string
}

func (f viewFile) ViewLink() *url.URL {
	return links.View(f.FullPath())
}

func (f viewFile) DownloadLink() *url.URL {
	return links.DownloadFile(f.FullPath())
}

func (f viewFile) RenameLink() *url.URL {
	return links.Rename(f.FullPath())
}

func (f viewFile) RemoveLink() *url.URL {
	return links.Remove(f.FullPath(), f.csrf)
}

type DTO struct {
	Files         []viewFile
	Dir           string
	CreateFileUrl *url.URL
	CreateDirUrl  *url.URL
	Breadcrumb    breadcrumb.Breadcrumb
}

func View(path string, c *gin.Context) {
	fsFiles, err := getFiles(path)

	if err != nil {
		errorpage.Page("Не удается открыть папку "+path, c)
		return
	}

	files := make([]viewFile, 0)
	for _, f := range fsFiles {
		files = append(files, viewFile{f, csrf.GetToken(c)})
	}

	dto := DTO{
		Files:         files,
		Dir:           path,
		CreateFileUrl: links.CreateFile(path),
		CreateDirUrl:  links.CreateDir(path),
		Breadcrumb:    breadcrumb.NewBreadcrumb(path),
	}

	fmt.Println(fmt.Sprintf("%+v", dto.CreateFileUrl))

	c.HTML(http.StatusOK, "dir-list.html", dto)
}

func getFiles(dir string) ([]filesystem.File, error) {
	files := make([]filesystem.File, 0)

	fsFiles, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	for _, fsFile := range fsFiles {
		files = append(files, filesystem.NewFile(fsFile, dir))
	}

	return files, nil
}
