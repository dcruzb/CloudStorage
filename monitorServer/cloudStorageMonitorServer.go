package main

import (
	"github.com/dcbCIn/CloudStorage/shared"
	"github.com/dcbCIn/MidCloud/lib"
	"github.com/dcbCIn/MidCloudMAPEK"
	"sync"
)

func main() {
	lib.PrintlnInfo("Initializing server CloudStorageMonitor")
	var wg = sync.WaitGroup{}
	wg.Add(1)

	chanAnalyzer := make(chan []MidCloudMAPEK.CloudService)
	chanPlanner := make(chan MidCloudMAPEK.CloudService)
	chanExecutor := make(chan MidCloudMAPEK.CloudService)

	monitor := MidCloudMAPEK.Monitor{}
	go monitor.Start(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT, "CloudFunctions", chanAnalyzer)

	go MidCloudMAPEK.Analyze(chanAnalyzer, chanPlanner)

	go MidCloudMAPEK.Plan(chanPlanner, chanExecutor)

	go MidCloudMAPEK.Execute(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT, "cloudFunctions", chanExecutor)

	wg.Wait()
	lib.PrintlnInfo("Fim do Servidor CloudStorageMonitor")
}
