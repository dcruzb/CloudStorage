package googleAPI

import (
	"CloudStorage/cloudLib"
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dcbCIn/MidCloud/lib"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type JsonGoogleCloud struct {
	Sku []Sku `json:"skus"`
}

type Sku struct {
	Name        string        `json:"name"`
	PricingInfo []PricingInfo `json:"pricingInfo"`
}

type PricingInfo struct {
	PriceExpression PricingExpression `json:"pricingExpression"`
}

type PricingExpression struct {
	TieredRates []TieredRate `json:"tieredRates"`
}

type TieredRate struct {
	UnitPrice UnitPrice `json:"unitPrice"`
}

type UnitPrice struct {
	Units string `json:"units"`
	Nanos int    `json:"nanos"`
}

type Google struct {
}

func (Google) Price(size float64) (price float64, err error) {
	//return 3.1234

	//jsonFile, err := os.Open("data.json")

	url := "https://cloudbilling.googleapis.com/v1/services/95FF-2EF5-5EA1/skus?key=Key_Code"

	response, erro := http.Get(url)

	if erro != nil {
		//Caso tenha tido erro, ele é apresentado na tela
		lib.PrintlnError("Erro ao abrir json. Erro", erro)
	}

	//defer jsonFile.Close()

	// lendo o json do response do http request
	responseJson, erro := ioutil.ReadAll(response.Body)

	jsonGoogle := JsonGoogleCloud{}

	erro = json.Unmarshal(responseJson, &jsonGoogle)
	if erro != nil {
		lib.PrintlnError("Erro ao realizar unmarshal. Erro:", erro)
	}

	for _, sku := range jsonGoogle.Sku {
		if sku.Name == "services/95FF-2EF5-5EA1/skus/8A46-D6C4-859E" {
			floatvalue, _ := strconv.ParseFloat(sku.PricingInfo[0].PriceExpression.TieredRates[0].UnitPrice.Units+"."+
				strconv.Itoa(sku.PricingInfo[0].PriceExpression.TieredRates[0].UnitPrice.Nanos), 64)

			price := size * floatvalue

			return price, nil
		}
	}

	// if price not found
	return price, errors.New("Price not found. ")

}

func (Google) Availability() (available bool, err error) {
	// TODO implement Google Availability
	return true, nil
}

func (Google) SendFile(file *os.File, path string) (createdFile cloudLib.CloudFile, err error) {

	ctx := context.Background()
	//projectID := "gifted-vigil-245219"

	// este arquivo autentica aplicação no serviço do storage da google cloud storage
	client, err1 := storage.NewClient(ctx, option.WithCredentialsFile("C:/Users/CASA/go/src/CloudStorage/cloudLib/google/googleAPI/My First Project-41269a52f4a2.json"))
	if err1 != nil {
		log.Fatalln(err1)
	}

	bucketName := "midd_cloud"

	bkt := client.Bucket(bucketName)

	// O bucket midd_cloud já foi criado manualmente. O comando abaixo gera um novo bucket, por isso está comentado.
	//if err := bkt.Create(ctx, projectID, nil); err != nil {
	//	log.Fatalf("Failed to create bucket: %v", err)
	//}

	attrs, err3 := bkt.Attrs(ctx)
	if err3 != nil {
		log.Fatalln(err3)
	}
	fmt.Printf("bucket %s, created at %s, is located in %s with storage class %s\n", attrs.Name, attrs.Created, attrs.Location, attrs.StorageClass)

	obj := bkt.Object(path + filepath.Base(file.Name()))

	w := obj.NewWriter(ctx)

	fileInfo, _ := file.Stat()
	buffer := make([]byte, fileInfo.Size())

	w.Write(buffer)

	if err4 := w.Close(); err != nil {
		log.Fatalln(err4)
	}

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	createdFile.Id = file.Name()
	createdFile.Cloud = "Google Cloud Platform"
	createdFile.Path = path + filepath.Base(file.Name())
	createdFile.Size = strconv.FormatInt(fileStat.Size(), 10)
	createdFile.Created = time.Now()
	createdFile.LastChecked = time.Now()

	return createdFile, nil
}

func (Google) GetFile(fileName string, path string) (file *os.File, err error) {

	ctx := context.Background()

	// este arquivo autentica aplicação no serviço do storage da google cloud storage
	client, err1 := storage.NewClient(ctx, option.WithCredentialsFile("C:/Users/CASA/go/src/CloudStorage/cloudLib/google/googleAPI/My First Project-41269a52f4a2.json"))
	if err1 != nil {
		log.Fatalln(err1)
	}

	bucketName := "midd_cloud"

	rc, err2 := client.Bucket(bucketName).Object(path + fileName).NewReader(ctx)
	if err2 != nil {
		log.Fatalln(err2)
	}

	fmt.Print(rc)

	defer rc.Close()

	data, err3 := ioutil.ReadAll(rc)
	if err3 != nil {
		log.Fatalln(err3)
	}

	file, err = os.Create("C:/temp/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	file.Write(data)
	fmt.Print(file.Stat())

	return file, err
}

func (Google) List(path string) (files []cloudLib.CloudFile, err error) {

	ctx := context.Background()

	// este arquivo autentica aplicação no serviço do storage da google cloud storage
	client, err1 := storage.NewClient(ctx, option.WithCredentialsFile("C:/Users/CASA/go/src/CloudStorage/cloudLib/google/googleAPI/My First Project-41269a52f4a2.json"))
	if err1 != nil {
		log.Fatalln(err1)
	}

	bucketName := "midd_cloud"

	it := client.Bucket(bucketName).Objects(ctx, nil)

	i := 0
	filesT := [1000]cloudLib.CloudFile{}
	for  {
		attrs, err2 := it.Next()
		if err2 == iterator.Done {
			break
		}
		if err2 != nil {
			log.Fatalln(err2)
		}

		filesT[i].Id = filepath.Base(attrs.Name)
		filesT[i].Cloud = "Google Cloud Platform"
		filesT[i].Path = path + filepath.Base(attrs.Name)
		filesT[i].Size = strconv.FormatInt(attrs.Size, 10)
		filesT[i].Created = attrs.Created
		filesT[i].LastChecked = attrs.Created
		i++
	}

	fmt.Print(filesT[0].Id)

	return files, nil
}
