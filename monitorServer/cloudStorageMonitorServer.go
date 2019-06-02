package main

import (
	"CloudStorage/shared"
	"github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/dcbCIn/MidCloud/services/common"
)

func main() {
	lib.PrintlnInfo("Initializing server CloudStorageMonitor")

	lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	err := lp.Bind("cloudMonitor", common.ClientProxy{Ip: shared.MONITOR_IP, Port: shared.MONITOR_PORT, ObjectId: 2000})
	lib.FailOnError(err, "Error at lookup.")
	err = lp.Close()
	lib.FailOnError(err, "Error at closing lookup")

	monitor := common.Monitor{}
	go monitor.Start(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT, "cloudFunctions", "CloudFunctions")

	// escuta na porta tcp configurada
	inv := dist.InvokerImpl{}
	inv.Register(2000, monitor)
	err = inv.Invoke(shared.MONITOR_PORT)
	lib.FailOnError(err, "Error calling invoker.")

	lib.PrintlnInfo("Fim do Servidor CloudStorageMonitor")
}
