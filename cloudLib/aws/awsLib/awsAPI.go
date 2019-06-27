package awsLib

import (
	"CloudStorage/cloudLib"
	"encoding/json"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/minio/minio-go"
	"io/ioutil"
	"log"

	"net/http"
	"os"

	"strconv"
)

type JsonAWS struct {
	Nuvem AWS `json:"regions"`
}

type AWS struct {
	Regions Regions `json:"South America (Sao Paulo)"`
}

type Regions struct {
	Br_sp Region `json:"Tags Storage per Tag Mo"`
}

type Region struct {
	Price string `json:"price"`
}

type Aws struct {
}

func (Aws) Price(size float64) float64 {
	//jsonFile, err := os.Open("data.json")

	url := "https://b0.p.awsstatic.com/pricing/2.0/meteredUnitMaps/s3/USD/current/s3.json"

	response, erro := http.Get(url)

	if erro != nil {
		//Caso tenha tido erro, ele é apresentado na tela
		lib.PrintlnError("Erro ao abrir json. Erro", erro)
	}

	//defer jsonFile.Close()

	// lendo o json do response do http request
	responseJson, erro := ioutil.ReadAll(response.Body)

	aws := JsonAWS{}

	erro = json.Unmarshal(responseJson, &aws)
	if erro != nil {
		lib.PrintlnError("Erro ao realizar unmarshal. Erro:", erro)
	}

	floatvalue, _ := strconv.ParseFloat(aws.Nuvem.Regions.Br_sp.Price, 64)

	price := size * floatvalue

	return price
}

func (Aws) SendFile(file *os.File) (createdFile cloudLib.CloudFile, err error) {

	endpoint := "s3.amazonaws.com"
	accessKeyID := "AKIA4HA7C4NR5EBWXSHS"
	secretAccessKey := "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"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup

	// Make a new bucket called mymusic.
	bucketName := "porra"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("Já existe %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Criado com sucesso %s\n", bucketName)
	}


	//file.Name()

	// Upload the zip file
	objectName := "mid-cloud.zip"
	filePath := "mid-cloud.zip"
	contentType := "application/zip"

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType:contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

	return createdFile, nil
}
