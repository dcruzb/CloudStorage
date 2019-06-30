package main

import (
	"CloudStorage/cloudLib/aws"
	"CloudStorage/cloudLib/google"
	"CloudStorage/shared"
	"fmt"
	"github.com/dcbCIn/MidCloud/lib"
	"os"
	"time"
)

func main() {
	// Todo criar proxy para StorageFunctions

	aws := aws.AwsFunctions{}
	//aws.Price(14.0)

	fileTeste, err := os.Open("C:/Users/CASA/Desktop/mid-cloud.zip");
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileTeste.Close()

	aws.SendFile(fileTeste, "cloudstorage/")

	return

	google := google.GoogleFunctions{}
	fileGoogle, err2 := os.Open("C:/Users/CASA/Desktop/mid-cloud.zip");
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer fileGoogle.Close()

	google.SendFile(fileGoogle, "cloudstorage/")

	//var jp dist.JankenpoProxy
	// connect to server
	//jp = *dist.NewJankenpoProxy(cp.Ip, cp.Port, cp.ObjectId)

	lib.PrintlnInfo("Connected successfully")
	lib.PrintlnInfo()

	//var player1Move, player2Move string
	// loop
	//start := time.Now()
	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		lib.PrintlnMessage("Game", i)

		//	player1Move, player2Move = shared.GetMoves(auto)

		// send request to server and receive reply at the same time
		//	result, err := jp.Play(player1Move, player2Move)
		if err != nil {
			lib.FailOnError(err, "Erro ao obter resultado do jogo no servidor. Erro:")
		}

		lib.PrintlnMessage()
		switch 1 {
		case 1, 2:
			lib.PrintlnMessage("The winner is Player", 1)
		case 0:
			lib.PrintlnMessage("Draw")
		default:
			lib.PrintlnMessage("Invalid move")
		}
		lib.PrintlnMessage("------------------------------------------------------------------")
		lib.PrintlnMessage()
		time.Sleep(shared.WAIT * time.Millisecond)
	}
	//elapsed = time.Since(start)
	//return elapsed*/
}
