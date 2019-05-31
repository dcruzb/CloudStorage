package main

import (
/*"CloudStorage/shared"
"github.com/dcbCIn/MidCloud/distribution"
"github.com/dcbCIn/MidCloud/lib"
"time"*/
)

func main() {
	// Todo finalizar proxy

	/*lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	cp, err := lp.Lookup("cloudFunctions")
	if err != nil {
		lib.PrintlnError("Error at lookup")
	}
	err = lp.Close()
	if err != nil {
		lib.PrintlnError("Error at closing lookup")
	}

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
		switch result {
		case 1, 2:
			lib.PrintlnMessage( "The winner is Player", result)
		case 0:
			lib.PrintlnMessage( "Draw")
		default:
			lib.PrintlnMessage("Invalid move")
		}
		lib.PrintlnMessage( "------------------------------------------------------------------")
		lib.PrintlnMessage()
		time.Sleep(shared.WAIT * time.Millisecond)
	}
	//elapsed = time.Since(start)
	//return elapsed*/
}
