package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/dcbCIn/CloudStorage/cloudLib"
	"github.com/dcbCIn/CloudStorage/shared"
	dist "github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	dtBegin := time.Now()
	for i := 1; i <= shared.SAMPLE_SIZE; i++ {
		dtStart := time.Now()
		sendFile()
		shared.LogEvent(shared.LOG, "cloudStorage", "main", "sendFile", "finished", strconv.Itoa(i), dtStart, time.Since(dtStart))
		time.Sleep(shared.WAIT)
	}
	shared.LogEvent(shared.LOG, "cloudStorage", "main", "everything", "finished", "none", dtBegin, time.Since(dtBegin))
}

func sendFile() {
	lib.PrintlnInfo("Initializing client CloudStorage")

	lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	cp, err := lp.Lookup("cloudFunctions") //"googleCloudFunctions") //
	lib.FailOnError(err, "Error at lookup")
	err = lp.Close()
	lib.FailOnError(err, "Error at closing lookup")

	var sp cloudLib.StorageFunctionsProxy
	sp = *cloudLib.NewStorageFunctionsProxy(cp.Ip, cp.Port, cp.ObjectId)
	defer sp.Close()

	file, err := os.Open("C:/Users/dcruz/OneDrive/Documents/Mestrado/Download artigos para Fagner/preview.mini.jpg") //p426-hilton.pdf") //
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	dtStart := time.Now()
	cloudFile, err := sp.SendFile(encoded, filepath.Base(file.Name()), "cloudstorage/")
	shared.LogEvent(shared.LOG, "cloudStorage", "sendFile", "sp.SendFile", "finished", "none", dtStart, time.Since(dtStart))
	lib.FailOnError(err, "Error sending file")

	lib.PrintlnInfo("File sent successfully. File:", cloudFile.Id, "Cloud:", cloudFile.Cloud)

	lib.PrintlnInfo("Fim do client CloudStorage")

}
