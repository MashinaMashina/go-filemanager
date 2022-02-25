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

	action := c.DefaultQuery("action", "")

	switch action {
	case "remove":
		general.Remove(path, c)
	case "rename":
		general.Rename(path, c)
	default:
		if isDir {
			switch action {
			case "upload":
				file.Upload(path, c)
			case "create_file":
				file.Create(path, c)
			case "create_dir":
				dir.Create(path, c)
			default:
				dir.View(path, c)
			}
		} else {
			switch action {
			case "download":
				file.DownloadFile(path, c)
			default:
				file.View(path, c)
			}
		}
	}
}
