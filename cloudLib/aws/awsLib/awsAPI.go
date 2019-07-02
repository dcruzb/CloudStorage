package awsLib

import (
	"CloudStorage/cloudLib"
	"CloudStorage/shared"
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/minio/minio-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func (Aws) SendFile(base64File string, fileName string, remotePath string) (createdFile cloudLib.CloudFile, err error) {

	endpoint := "s3.amazonaws.com"
	accessKeyID := shared.AWS_ACCESS_KEY_ID
	secretAccessKey := shared.AWS_SECRET_ACCESS_KEY
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		lib.PrintlnError(err)
		return createdFile, *shared.NewRemoteError(err.Error())
	}

	lib.PrintlnDebug(minioClient) // minioClient is now setup

	// Make a new bucket called mymusic.
	bucketName := "ufpestorage"
	location := "sa-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			lib.PrintlnDebug("Já existe %s\n", bucketName)
		} else {
			lib.PrintlnError(err)
			return createdFile, *shared.NewRemoteError(err.Error())
		}
	} else {
		lib.PrintlnDebug("Criado com sucesso %s\n", bucketName)
	}

	decFile, err := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		lib.PrintlnError(err)
		return createdFile, *shared.NewRemoteError(err.Error())
	}

	file, err := os.Create("./temp/" + fileName)
	if err != nil {
		lib.PrintlnError(err)
		return createdFile, *shared.NewRemoteError(err.Error())
	}

	defer file.Close()

	if _, err := file.Write(decFile); err != nil {
		lib.PrintlnError(err)
		return createdFile, *shared.NewRemoteError(err.Error())
	}
	if err := file.Sync(); err != nil {
		lib.PrintlnError(err)
		return createdFile, *shared.NewRemoteError(err.Error())
	}

	fileStat, err := file.Stat()
	if err != nil {
		lib.PrintlnError(err)
		return createdFile, *shared.NewRemoteError(err.Error())
	}

	//fileTeste, err := os.Open("./temp/" + fileName)
	//if err != nil {
	//	lib.PrintlnError(err)
	//	return createdFile, *shared.NewRemoteError(err.Error())
	//}

	//defer fileTeste.Close()

	file.Seek(0, 0)

	n, err := minioClient.PutObject(bucketName, remotePath+fileName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		lib.PrintlnError(err)
		return createdFile, *shared.NewRemoteError(err.Error())
	}

	fmt.Println("Successfully uploaded bytes: ", n)

	createdFile.Id = fileName
	createdFile.Cloud = "AWS"
	createdFile.Path = remotePath + fileName
	size := (float64)(len(decFile)) / 1024 / 1024 // Convert to mb
	createdFile.Size = fmt.Sprintf("%f", size)    //strconv.FormatInt( fileInfo.Size(), 10)
	createdFile.Created = time.Now()
	createdFile.LastChecked = time.Now()

	// remover arquivo da pasta temp
	//err2 := os.Remove("./cloudLib/aws/awsLib/temp/" + filepath.Base(fileTeste.Name()))
	//if err2 != nil {
	//	fmt.Println(err2.Error())
	//}

	return createdFile, nil
}

func (Aws) GetFile(fileName string, path string) (base64File string, err error) {
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

	file, err2 := os.Create("./temp/" + fileName)
	if err2 != nil {
		fmt.Println(err2)
	}

	defer file.Close()

	fileInfo, _ := object.Stat()
	buffer := make([]byte, fileInfo.Size)
	object.Read(buffer)

	file.Write(buffer)

	fileTeste, err := os.Open("./temp/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fileTeste.Close()

	reader := bufio.NewReader(fileTeste)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	base64File = base64.StdEncoding.EncodeToString([]byte(content))

	// remover arquivo da pasta temp
	//err2 := os.Remove("./cloudLib/aws/awsLib/temp/" + filepath.Base(fileTeste.Name()))
	//if err2 != nil {
	//	fmt.Println(err2.Error())
	//}

	return base64File, nil
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
