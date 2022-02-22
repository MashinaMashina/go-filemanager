package viewer

import (
	"file-manager/internal/errorpage"
	"file-manager/internal/filesystem"
	"file-manager/internal/viewer/dir"
	"file-manager/internal/viewer/file"
	"file-manager/internal/viewer/general"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func openPath(c *gin.Context) {
	path, err := filepath.Abs(c.DefaultQuery("path", "./"))

	if err != nil {
		errorpage.Page(err.Error(), c)
		return
	}

	isDir, err := filesystem.IsDir(path)

	if err != nil {
		errorpage.Page("Не удается открыть "+path, c)
		return
	}

	if _, remove := c.GetQuery("remove"); remove {
		general.Remove(path, c)
	} else if _, rename := c.GetQuery("rename"); rename {
		general.Rename(path, c)
	} else if isDir {
		if _, fileCreate := c.GetQuery("create_file"); fileCreate {
			file.Create(path, c)
		} else if _, dirCreate := c.GetQuery("create_dir"); dirCreate {
			dir.Create(path, c)
		} else {
			dir.View(path, c)
		}
	} else {
		file.View(path, c)
	}
}
