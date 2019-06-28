package awsLib

import (
	"CloudStorage/cloudLib"
	"encoding/json"
	"fmt"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/minio/minio-go"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"

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
	Br_sp Region `json:"Standard Storage First 50 TB per GB Mo"`
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

func (Aws) SendFile(file *os.File, path string) (createdFile cloudLib.CloudFile, err error) {

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
	bucketName := "cloudstorage1234"
	location := "sa-east-1"

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

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := minioClient.PutObject(bucketName, path + filepath.Base(file.Name()), file, fileStat.Size(), minio.PutObjectOptions{ContentType:"application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully uploaded bytes: ", n)

	return createdFile, nil
}

func (Aws) GetFile(fileName string, path string) (file *os.File, err error) {
	endpoint := "s3.amazonaws.com"
	accessKeyID := "AKIA4HA7C4NR5EBWXSHS"
	secretAccessKey := "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"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	object, err := minioClient.GetObject("cloudstorage1234" , path + fileName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err = os.Create("../files/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err = io.Copy(file, object); err != nil {
		fmt.Println(err)
		return
	}

	return file, nil
}
