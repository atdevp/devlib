package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func isExist(err error) bool {
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsDir(name string) bool {
	fi, err := Stat(name)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func IsFile(name string) bool {
	if IsDir(name) {
		return false
	}
	return true
}

func IsExist(name string) bool {
	_, e := Stat(name)
	if isExist(e) {
		return true
	}
	return false
}

func IsExec(name string) bool {

	fi, err := Stat(name)
	if !isExist(err) {
		return false
	}

	if fi.IsDir() {
		return false
	}

	m := fi.Mode().String()
	if strings.Index(m, "x") == 3 {
		return true
	}
	return false

}

func Size(name string) (int64, error) {
	fi, e := os.Stat(name)
	if fi.IsDir() {
		return 0, errors.New("invalid file type")
	}

	if !isExist(e) {
		return 0, e
	}
	return fi.Size(), nil
}

func Create(name string) (*os.File, error) {
	return os.Create(name)
}

func Delete(name string) error {
	return os.Remove(name)
}

func Abspath(name string) (string, error) {
	return filepath.Abs(filepath.Dir(name))
}
