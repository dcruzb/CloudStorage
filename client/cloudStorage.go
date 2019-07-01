package main

import (
	"CloudStorage/cloudLib"
	"CloudStorage/shared"
	"bufio"
	"encoding/base64"
	"fmt"
	dist "github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	for i := 0; i <= shared.SAMPLE_SIZE; i++ {
		shared.LogEvent("cloudStorage", "main", "enviaArquivo", "iniciado", strconv.Itoa(i))
		enviaArquivo()
		shared.LogEvent("cloudStorage", "main", "enviaArquivo", "finalizado", strconv.Itoa(i))
	}
}

func enviaArquivo() {
	lib.PrintlnInfo("Initializing client CloudStorage")

	lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	cp, err := lp.Lookup("awsCloudFunctions") //"googleCloudFunctions") //
	lib.FailOnError(err, "Error at lookup.")
	err = lp.Close()
	lib.FailOnError(err, "Error at closing lookup")

	var sp cloudLib.StorageFunctionsProxy
	sp = *cloudLib.NewStorageFunctionsProxy(cp.Ip, cp.Port, cp.ObjectId)
	defer sp.Close()

	fileTeste, err := os.Open("C:/Users/dcruz/OneDrive/Documents/Mestrado/Download artigos para Fagner/preview.mini.jpg") //p426-hilton.pdf") //
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileTeste.Close()

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(fileTeste)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	cloudfile, err := sp.SendFile(encoded, filepath.Base(fileTeste.Name()), "cloudstorage/")
	lib.FailOnError(err, "Error sending file")

	lib.PrintlnInfo("File sent successfully. File:", cloudfile.Id, "Cloud:", cloudfile.Cloud)

	lib.PrintlnInfo("Fim do client CloudStorage")

}
