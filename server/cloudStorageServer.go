package main

import (
	"CloudStorage/server/lib/google"
	"CloudStorage/shared"
	"github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/dcbCIn/MidCloud/services/common"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	lib.PrintlnInfo("Initializing server CloudStorage")

	lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	err := lp.Bind("cloudFunctions", common.ClientProxy{"127.0.0.1", shared.MID_PORT, 2000})
	if err != nil {
		lib.PrintlnError("Error at lookup: ", err)
	}

	err = lp.Close()
	if err != nil {
		lib.PrintlnError("Error at closing lookup")
	}

	// escuta na porta tcp configurada
	var inv dist.InvokerImpl
	//inv.StartServer("", strconv.Itoa(shared.RPC_PORT))
	//defer inv.StopServer()
	var gf google.GoogleFunctions
	inv.Register(2000, gf)

	go inv.Invoke(shared.MID_PORT)
	wg.Add(1)
	/*for idx := 0; idx < shared.CONECTIONS; idx++ {
		wg.Add(1)
		go func(i int) {
			waitForConection(inv, i)

			wg.Done()
		}(idx)
	}*/
	wg.Wait()
	lib.PrintlnInfo("Fim do Servidor CloudStorage")
}
