package filesystem

import (
	"github.com/dustin/go-humanize"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type File interface {
	Name() string
	FullPath() string
	Size() int64
	HumanSize() string
	Mode() fs.FileMode
	ModTime() time.Time
	HumanModTime() string
	IsDir() bool
	Sys() interface{}
	IsHidden() bool
}

type file struct {
	fs.FileInfo
	dir string
}

func NewFile(info fs.FileInfo, dir string) File {
	return file{info, filepath.Clean(dir)}
}

func NewFileFromPath(path string) (File, error) {
	lstat, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}

	dir, _ := filepath.Split(path)
	return file{lstat, filepath.Clean(dir)}, nil
}

func (f file) FullPath() string {
	return f.dir + string(os.PathSeparator) + f.Name()
}

func (f file) HumanModTime() string {
	return f.ModTime().Format("2006.01.02 15:04")
}

func (f file) HumanSize() string {
	return humanize.Bytes(uint64(f.Size()))
}
