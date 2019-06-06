package aws

import (
	"CloudStorage/cloudLib"
	"CloudStorage/cloudLib/aws/awsLib"
	"os"
)

type AwsFunctions struct {
}

func (AwsFunctions) SendFile(file *os.File) (createdFile cloudLib.CloudFile, err error) {
	panic("implement me")
}

func (AwsFunctions) GetFile() (file *os.File, err error) {
	panic("implement me")
}

func (AwsFunctions) List(path string) (files []cloudLib.CloudFile, err error) {
	panic("implement me")
}

func (AwsFunctions) Price(size float64) float64 {
	// Todo usar cache dos dados
	// Todo criar função para preencher o cache
	// Todo obter arquivo data.json diretamente da AWS

	aws := awsLib.Aws{}

	return aws.Price(size)
}

func (AwsFunctions) Availability() bool {
	return false
}
