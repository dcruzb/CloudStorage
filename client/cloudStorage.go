package main

import (
	"CloudStorage/shared"
	dist "github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/dcbCIn/MidCloud/services/common"
	"time"
)

func main() {
	// Todo criar proxy para StorageFunctions

	lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	defer func() { err := lp.Close(); lib.FailOnError(err, "Error at closing lookup") }()

	err := lp.Bind("googleCloudFunctions", common.ClientProxy{Ip: shared.AWS_SERVER_IP, Port: shared.AWS_SERVER_PORT, ObjectId: 2000}) // Todo tirar daqui depois, somente para testes
	lib.FailOnError(err, "Error at lookup.")

	cp, err := lp.Lookup("googleCloudFunctions")
	lib.FailOnError(err, "Error at lookup.")

	lib.PrintlnInfo("CP:", cp, "ip:", cp.Ip, "Port:", cp.Port, "ObjectId:", cp.ObjectId)

	return

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
