package shared

import (
	"github.com/dcbCIn/MidCloud/lib"
	"os"
	"time"
)

func LogEvent(log bool, source string, event string, action string, status string, thread string) {
	if !log {
		return
	}

	file, err := os.OpenFile("./temp/logEvent.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		lib.PrintlnError("Erro ao abrir arquivo do log para inclusão de novo evento. Erro:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(source + "," + event + "," + action + "," + status + "," + thread + "," + time.Now().Format(time.RFC3339Nano) + "\n")
	if err != nil {
		lib.PrintlnError("Erro ao adicionar evento ao log. Erro:", err)
	}
	err = file.Sync()
}
