package shared

import (
	"github.com/dcbCIn/MidCloud/lib"
	"os"
	"time"
)

func LogEvent(log bool, source string, method string, action string, status string, thread string, dtStart time.Time, duration time.Duration) {
	if !log {
		return
	}

	file, err := os.OpenFile("./temp/logEvent_"+source+"_"+method+"_"+action+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		lib.PrintlnError("Erro ao abrir arquivo do log para inclusão de novo evento. Erro:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(source + ";" + method + ";" + action + ";" + status + ";" + thread + ";" + dtStart.Format(time.RFC3339Nano) + ";" + duration.String() + "\n")
	if err != nil {
		lib.PrintlnError("Erro ao adicionar evento ao log. Erro:", err)
	}
	err = file.Sync()
}
