package filesystem

import (
	"syscall"
)

func (f file) IsHidden() bool {
	pointer, err := syscall.UTF16PtrFromString(f.FullPath())
	if err != nil {
		return false
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false
	}
	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0
}
