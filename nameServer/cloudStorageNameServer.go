package main

import (
	"CloudStorage/shared"
	"github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/dcbCIn/MidCloud/services/common"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	lib.PrintlnInfo("nameServer", "Initializing server CloudStorage(NameServer)")

	// escuta na porta tcp configurada
	var inv dist.InvokerImpl
	lookup := common.Lookup{}
	inv.Register(0, lookup)
	go inv.Invoke(shared.NAME_SERVER_PORT)
	wg.Add(1)

	wg.Wait()
	lib.PrintlnInfo("nameServer", "Fim do Servidor CloudStorage(NameServer)")
}
