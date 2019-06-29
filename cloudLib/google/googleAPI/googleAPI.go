package googleAPI

import (
	"CloudStorage/cloudLib"
	"os"
)

type Google struct {
}

func (Google) Price(size float64) float64 {
	return 3.1234
}

func (Google) SendFile(file *os.File, path string) (createdFile cloudLib.CloudFile, err error) {

	return createdFile, nil
}

func (Google) GetFile(fileName string, path string) (file *os.File, err error) {

	return file, nil
}

func (Google) List(path string) (files []cloudLib.CloudFile, err error) {

	return files, nil
}
