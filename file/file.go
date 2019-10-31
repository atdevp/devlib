package file

import (
	"os"
)

// 判断是否为目录
func IsDir(p string) bool {
	fi, err := os.Stat(p)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func IsFile(p string) bool {
	if IsDir(p) {
		return false
	}
	return true
}

func IsExists(filename string) bool {
	os.Chmod
}
