package assets

import (
	"io/fs"
)

func Subdir(path string) (fs.FS, error) {
	return fs.Sub(FS, path)
}
