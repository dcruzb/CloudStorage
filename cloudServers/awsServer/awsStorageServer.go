package main

import (
	"CloudStorage/cloudLib/aws"
	"CloudStorage/shared"
	"github.com/dcbCIn/MidCloud/distribution"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/dcbCIn/MidCloud/services/common"
)

func main() {
	lib.PrintlnInfo("Initializing server AwsStorage")

	lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	err := lp.Bind("awsCloudFunctions", common.ClientProxy{Ip: shared.AWS_SERVER_IP, Port: shared.AWS_SERVER_PORT, ObjectId: 2000})
	lib.FailOnError(err, "Error at lookup.")
	err = lp.Close()
	lib.FailOnError(err, "Error at closing lookup")

	// escuta na porta tcp configurada
	var inv dist.InvokerImpl
	inv.Register(2000, &aws.AwsFunctions{})

	err = inv.Invoke(shared.AWS_SERVER_PORT, shared.CONNECTIONS)
	lib.FailOnError(err, "Error calling invoker.")

	lib.PrintlnInfo("AwsStorage server finished")
}
