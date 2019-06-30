package googleAPI

import (
	"CloudStorage/cloudLib"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type JsonGoogleCloud struct {
	PriceInfo PricingInfo `json:"skus"`
}

type PricingInfo struct {
	PriceExpression PricingExpression `json:"pricingInfo"`
}

type PricingExpression struct {

}

type Google struct {
}

func (Google) Price(size float64) float64 {
	return 3.1234
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

	return file, nil
}

func (Google) List(path string) (files []cloudLib.CloudFile, err error) {

	return files, nil
}
