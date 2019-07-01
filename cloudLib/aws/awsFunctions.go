package aws

import (
	"CloudStorage/cloudLib"
	"CloudStorage/cloudLib/aws/awsLib"
)

type AwsFunctions struct {
}

func (AwsFunctions) SendFile(base64File string, fileName string, path string) (createdFile cloudLib.CloudFile, err error) {
	aws := awsLib.Aws{}
	return aws.SendFile(base64File, fileName, path)
}

func (AwsFunctions) GetFile(fileName string, path string) (base64File string, err error) {
	aws := awsLib.Aws{}
	return aws.GetFile(fileName, path)
}

func (AwsFunctions) List(path string) (files []cloudLib.CloudFile, err error) {
	aws := awsLib.Aws{}
	return aws.List(path)
}

func (AwsFunctions) Price(size float64) (price float64, err error) {
	// Todo usar cache dos dados
	// Todo criar função para preencher o cache
	aws := awsLib.Aws{}

	return aws.Price(size)
}

func (AwsFunctions) Availability() (available bool, err error) {
	aws := awsLib.Aws{}

	return aws.Availability()
}
