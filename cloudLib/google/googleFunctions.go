package google

import (
	"CloudStorage/cloudLib"
	"CloudStorage/cloudLib/google/googleAPI"
	"os"
)

type GoogleFunctions struct {
}

func (gf GoogleFunctions) SendFile(base64File string, fileName string, path string) (createdFile cloudLib.CloudFile, err error) {
	google := googleAPI.Google{}
	return google.SendFile(base64File, fileName, path)
}

func (gf GoogleFunctions) GetFile(fileName string, path string) (file *os.File, err error) {
	google := googleAPI.Google{}
	return google.GetFile(fileName, path)
}

func (gf GoogleFunctions) List(path string) (files []cloudLib.CloudFile, err error) {
	google := googleAPI.Google{}
	return google.List(path)
}

func (gf GoogleFunctions) Price(size float64) (price float64, err error) {
	// Todo usar cache dos dados
	// Todo criar função para preencher o cache
	google := googleAPI.Google{}

	return google.Price(size)
}

func (gf GoogleFunctions) Availability() (available bool, err error) {
	google := googleAPI.Google{}

	return google.Availability()
}
