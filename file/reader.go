package file

import (
	"io/ioutil"
	"strings"
)

func ToBytes(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func ToString(file string) (string, error) {
	bs, err := ToBytes(file)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func ToTirmString(file string) (string, error) {
	str, err := ToString(file)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(str), nil
}
