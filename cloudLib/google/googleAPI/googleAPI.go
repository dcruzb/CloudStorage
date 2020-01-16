package googleAPI

import (
	"bufio"
	"cloud.google.com/go/storage"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dcbCIn/CloudStorage/cloudLib"
	"github.com/dcbCIn/CloudStorage/shared"
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

	url := "https://cloudbilling.googleapis.com/v1/services/95FF-2EF5-5EA1/skus?key=" + shared.GOOGLE_KEY_CODE

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

func (Google) SendFile(base64File string, fileName string, remotePath string) (createdFile cloudLib.CloudFile, err error) {
	dtStart := time.Now()
	/*dec, err := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.Write(dec); err != nil {
		panic(err)
	}*/
	/*if err := file.Sync(); err != nil {
		panic(err)
	}*/

	ctx := context.Background()
	//projectID := "gifted-vigil-245219"

	// este arquivo autentica aplicação no serviço do storage da google cloud storage
	absPath, _ := filepath.Abs("./cloudLib/google/googleAPI/My First Project-41269a52f4a2.json")
	client, err1 := storage.NewClient(ctx, option.WithCredentialsFile(absPath)) //"./CloudStorage/cloudLib/google/googleAPI/My First Project-41269a52f4a2.json"))
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
	lib.PrintlnInfo("bucket", attrs.Name, " created at ", attrs.Created, ", is located in ", attrs.Location, " with storage class ", attrs.StorageClass)

	obj := bkt.Object(remotePath + fileName)

	w := obj.NewWriter(ctx)
	//	objAttrs, _ := obj.Attrs(ctx)
	//	objAttrs.ContentType = "image/jpg"
	//obj.Update(ctx, attr)
	//objAttrs := storage.ObjectAttrsToUpdate{ContentType: "image/jpg"}
	//obj.Attrs := objAttrs
	//obj.Update(ctx, objAttrs)

	/*_, err = io.Copy(w, file)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = w.Close()
	if err != nil {
		fmt.Println(err)
		return
	}*/

	//	buffer := make([]byte, 1024)//fileInfo.Size())

	decFile, err := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		panic(err)
	}

	w.Write(decFile)

	/*for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return createdFile, err
		}
		if n == 0 {
			break
		}

		if _, err := w.Write(buffer[:n]); err != nil {
			fmt.Println(err)
			return createdFile, err
		}
	}*/

	if err4 := w.Close(); err != nil {
		log.Fatalln(err4)
	}
	shared.LogEvent(shared.LOG, "googleAPI", "SendFile", "decodeAndWrite", "finished", "none", dtStart, time.Since(dtStart))

	//fileInfo, _ := decFile.Stat()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	createdFile.Id = fileName
	createdFile.Cloud = "Google Cloud Platform"
	createdFile.Path = remotePath + fileName
	size := (float64)(len(decFile)) / 1024 / 1024 // Convert to mb
	createdFile.Size = fmt.Sprintf("%f", size)    //strconv.FormatInt( fileInfo.Size(), 10)
	createdFile.Created = time.Now()
	createdFile.LastChecked = time.Now()

	return createdFile, nil
}

func (Google) GetFile(fileName string, path string) (base64File string, err error) {

	ctx := context.Background()

	// este arquivo autentica aplicação no serviço do storage da google cloud storage
	absPath, _ := filepath.Abs("./cloudLib/google/googleAPI/My First Project-41269a52f4a2.json")
	fmt.Println(absPath)
	client, err1 := storage.NewClient(ctx, option.WithCredentialsFile(absPath)) //"./CloudStorage/cloudLib/google/googleAPI/My First Project-41269a52f4a2.json"))
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

	file, err2 := os.Create("./temp/" + fileName)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	defer file.Close()

	file.Write(data)
	//fmt.Print(file.Stat())

	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	base64File = base64.StdEncoding.EncodeToString([]byte(content))
	fmt.Print(base64File)

	return base64File, err
}

func (Google) List(path string) (files []cloudLib.CloudFile, err error) {

	ctx := context.Background()

	// este arquivo autentica aplicação no serviço do storage da google cloud storage
	absPath, _ := filepath.Abs("./cloudLib/google/googleAPI/My First Project-41269a52f4a2.json")
	fmt.Println(absPath)
	client, err1 := storage.NewClient(ctx, option.WithCredentialsFile(absPath)) //"./CloudStorage/cloudLib/google/googleAPI/My First Project-41269a52f4a2.json"))
	if err1 != nil {
		log.Fatalln(err1)
	}

	bucketName := "midd_cloud"

	it := client.Bucket(bucketName).Objects(ctx, nil)

	i := 0
	filesT := [1000]cloudLib.CloudFile{}
	for {
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
