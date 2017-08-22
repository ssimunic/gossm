package gossm

import (
	"encoding/json"
	"github.com/ssimunic/gossm/logger"
	"net/http"
)

func RunHttp(address string, monitor *Monitor) {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		out, err := json.Marshal(monitor.Data)
		if err != nil {
			logger.Logln(err)
		}
		rw.Write(out)
	})
	http.ListenAndServe(address, nil)
}
