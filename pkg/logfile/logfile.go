package logfile

import (
	"os"
	"strings"
)

func InitLogFilePath(paths ...string) {
	for _, path := range paths {
		create(path)
	}
}

func create(path string) {
	parts := strings.Split(path, "/")
	dir := strings.Join(parts[:len(parts)-1], "/")
	os.MkdirAll(dir, 0755)
	os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
}
