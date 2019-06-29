package main

import (
	"CloudStorage/shared"
	"github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/dcbCIn/MidCloud/services/common"
)

func main() {
	lib.PrintlnInfo("nameServer", "Initializing server CloudStorage(NameServer)")

	// escuta na porta tcp configurada
	var inv dist.InvokerImpl
	inv.Register(0, &common.Lookup{})
	err := inv.Invoke(shared.NAME_SERVER_PORT, 5)
	lib.FailOnError(err, "Error calling invoker.")

	lib.PrintlnInfo("nameServer", "Fim do Servidor CloudStorage(NameServer)")
}
