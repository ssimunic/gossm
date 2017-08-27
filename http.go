package gossm

import (
	"io"
	"net/http"
	"strconv"
)

func RunHttp(address string, monitor *Monitor) {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		for server, serverStatus := range monitor.serverStatusData.GetServerStatus() {
			io.WriteString(rw, server.String()+"\n")
			for _, statusAtTime := range serverStatus {
				io.WriteString(rw, " "+strconv.FormatBool(statusAtTime.Status)+"\n")
			}
		}
	})

	http.ListenAndServe(address, nil)
}
