package file

import (
	"io/ioutil"
	"mime/multipart"
)

func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}
