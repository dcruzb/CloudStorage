package aws

import (
	"CloudStorage/cloudLib"
	"CloudStorage/cloudLib/aws/awsLib"
	"os"
)

type AwsFunctions struct {
}

func (AwsFunctions) SendFile(file *os.File, path string) (createdFile cloudLib.CloudFile, err error) {
	aws := awsLib.Aws{}
	return aws.SendFile(file, path)
}

func (AwsFunctions) GetFile(fileName string, path string) (file *os.File, err error) {
	aws := awsLib.Aws{}
	return aws.GetFile(fileName, path)
}

func (AwsFunctions) List(path string) (files []cloudLib.CloudFile, err error) {
	panic("implement me")
}

func (AwsFunctions) Price(size float64) (price float64, err error) {
	// Todo usar cache dos dados
	// Todo criar função para preencher o cache
	// Todo obter arquivo data.json diretamente da AWS

	aws := awsLib.Aws{}

	return aws.Price(size), nil
}

func (AwsFunctions) Availability() (available bool, err error) {
	return true, nil
}
