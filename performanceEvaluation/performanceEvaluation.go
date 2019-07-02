package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/dcbCIn/MidCloud/lib"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	path := "./temp/1000google/" //1000google-monitor/"
	//totalDuration, avarage := performanceEvaluation(path + "logEvent_cloudStorage_main_everything.csv")
	//fmt.Println("cloudStorage_main_everything;1;"+totalDuration.String()+";"+avarage.String()+";0;0")

	totalDuration2, avarage2 := performanceEvaluation(path + "logEvent_cloudStorage_main_sendFile.csv")
	//difference := totalDuration - totalDuration2
	//avarageDiference := difference.Nanoseconds()/1000
	//avgDiffDuration := time.Duration(avarageDiference)
	fmt.Println("cloudStorage_main_sendFile;1000;" + totalDuration2.String() + ";" + avarage2.String() + ";0;0;0") //+difference.String()+";"+avgDiffDuration.String())

	totalDuration3, avarage3 := performanceEvaluation(path + "logEvent_cloudStorage_sendFile_sp.SendFile.csv")
	difference := totalDuration2 - totalDuration3
	avarageDiference := difference.Nanoseconds() / 1000
	avgDiffDuration := time.Duration(avarageDiference)
	percDiff := float64(difference.Nanoseconds()) / float64(totalDuration2.Nanoseconds())
	spercDiff := fmt.Sprintf("%f", percDiff)
	fmt.Println("cloudStorage_sendFile_sp.SendFile;1000;" + totalDuration3.String() + ";" + avarage3.String() + ";" + difference.String() + ";" + spercDiff + ";" + avgDiffDuration.String())

	totalDuration4, avarage4 := performanceEvaluation(path + "logEvent_StorageFunctionsProxy_sendFile_sfp.requestor.Invoke.csv")
	difference = totalDuration3 - totalDuration4
	avarageDiference = difference.Nanoseconds() / 1000
	avgDiffDuration = time.Duration(avarageDiference)
	percDiff = float64(difference.Nanoseconds()) / float64(totalDuration2.Nanoseconds())
	spercDiff = fmt.Sprintf("%f", percDiff)
	fmt.Println("StorageFunctionsProxy_sendFile_sfp.requestor.Invoke;1000;" + totalDuration4.String() + ";" + avarage4.String() + ";" + difference.String() + ";" + spercDiff + ";" + avgDiffDuration.String())

	totalDuration5, avarage5 := performanceEvaluation(path + "logEvent_googleAPI_SendFile_decodeAndWrite.csv")
	difference = totalDuration4 - totalDuration5
	avarageDiference = difference.Nanoseconds() / 1000
	avgDiffDuration = time.Duration(avarageDiference)
	percDiff = float64(difference.Nanoseconds()) / float64(totalDuration2.Nanoseconds())
	spercDiff = fmt.Sprintf("%f", percDiff)
	fmt.Println("googleAPI_SendFile_decodeAndWrite;1000;" + totalDuration5.String() + ";" + avarage5.String() + ";" + difference.String() + ";" + spercDiff + ";" + avgDiffDuration.String())
}

func performanceEvaluation(path string) (totalDuration, avarage time.Duration) {
	file, err := os.Open(path) //"./temp/1000google - Teste analise/logEvent_StorageFunctionsProxy_sendFile_sfp.requestor.Invoke.csv")
	lib.FailOnError(err, "Erro ao abrir arquivo")

	reader := bufio.NewReader(file)
	var i int

	for {
		var buffer bytes.Buffer
		lb, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		i++

		lib.FailOnError(err, "Erro ao ler arquivo")

		buffer.Write(lb)

		event := buffer.String()
		//fmt.Println(event)

		parts := strings.Split(event, ";")
		duration, err := time.ParseDuration(parts[len(parts)-1])
		lib.FailOnError(err, "Erro ao ler duration: "+parts[len(parts)-1]+" linha: "+strconv.Itoa(i))

		totalDuration += duration
	}

	//fmt.Println("Qtd registros:", i)
	//fmt.Println("Total:", totalDuration)
	mediaNano := totalDuration.Nanoseconds() / int64(i)
	avarage = time.Duration(mediaNano)
	//fmt.Println("Média:", avarage)
	return totalDuration, avarage
}
