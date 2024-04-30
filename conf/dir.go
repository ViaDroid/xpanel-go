package conf

import (
	"path/filepath"
	"runtime"
)

func ConfDir(filename string) string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), filename)
}
