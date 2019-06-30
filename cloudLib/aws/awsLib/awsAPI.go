package awsLib

import (
	"CloudStorage/cloudLib"
	"CloudStorage/shared"
	"encoding/json"
	"fmt"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/minio/minio-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

func (Aws) Price(size float64) (price float64, err error) {
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

	price = size * floatvalue

	return price, nil
}

func (Aws) Availability() (available bool, err error) {
	// TODO implement Aws Availability
	return true, nil
}

func (Aws) SendFile(file *os.File, path string) (createdFile cloudLib.CloudFile, err error) {

	endpoint := "s3.amazonaws.com"
	accessKeyID := shared.AWS_ACCESS_KEY_ID
	secretAccessKey := shared.AWS_SECRET_ACCESS_KEY
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup

	// Make a new bucket called mymusic.
	bucketName := "ufpestorage"
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

	n, err := minioClient.PutObject(bucketName, path+filepath.Base(file.Name()), file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully uploaded bytes: ", n)

	createdFile.Id = file.Name()
	createdFile.Cloud = "AWS"
	createdFile.Path = path + filepath.Base(file.Name())
	createdFile.Size = strconv.FormatInt(fileStat.Size(), 10)
	createdFile.Created = time.Now()
	createdFile.LastChecked = time.Now()

	return createdFile, nil
}

func (Aws) GetFile(fileName string, path string) (file *os.File, err error) {
	endpoint := "s3.amazonaws.com"
	accessKeyID := shared.AWS_ACCESS_KEY_ID
	secretAccessKey := shared.AWS_SECRET_ACCESS_KEY
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	object, err := minioClient.GetObject("ufpestorage", path+fileName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err = os.Create("./files/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileInfo, _ := object.Stat()
	buffer := make([]byte, fileInfo.Size)
	object.Read(buffer)

	file.Write(buffer)
	fmt.Print(file.Stat())
	//
	//if _, err = io.Copy(file, buffer); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	return file, nil
}

func (Aws) List(path string) (files []cloudLib.CloudFile, err error) {

	endpoint := "s3.amazonaws.com"
	accessKeyID := shared.AWS_ACCESS_KEY_ID
	secretAccessKey := shared.AWS_SECRET_ACCESS_KEY
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a done channel to control 'ListObjectsV2' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := minioClient.ListObjectsV2("ufpestorage", "myprefix", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}

	return files, nil
}
