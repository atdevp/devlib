package file

import (
	"io/ioutil"
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
