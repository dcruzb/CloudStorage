package main

func main() {
	// Todo finalizar proxy

	/*lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	cp, err := lp.Lookup("cloudFunctions")
	if err != nil {
		cloudLib.PrintlnError("Error at lookup")
	}
	err = lp.Close()
	if err != nil {
		cloudLib.PrintlnError("Error at closing lookup")
	}

	//var jp dist.JankenpoProxy
	// connect to server
	//jp = *dist.NewJankenpoProxy(cp.Ip, cp.Port, cp.ObjectId)

	cloudLib.PrintlnInfo("Connected successfully")
	cloudLib.PrintlnInfo()

	//var player1Move, player2Move string
	// loop
	//start := time.Now()
	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		cloudLib.PrintlnMessage("Game", i)

	//	player1Move, player2Move = shared.GetMoves(auto)

		// send request to server and receive reply at the same time
	//	result, err := jp.Play(player1Move, player2Move)
		if err != nil {
			cloudLib.FailOnError(err, "Erro ao obter resultado do jogo no servidor. Erro:")
		}

		cloudLib.PrintlnMessage()
		switch result {
		case 1, 2:
			cloudLib.PrintlnMessage( "The winner is Player", result)
		case 0:
			cloudLib.PrintlnMessage( "Draw")
		default:
			cloudLib.PrintlnMessage("Invalid move")
		}
		cloudLib.PrintlnMessage( "------------------------------------------------------------------")
		cloudLib.PrintlnMessage()
		time.Sleep(shared.WAIT * time.Millisecond)
	}
	//elapsed = time.Since(start)
	//return elapsed*/
}
