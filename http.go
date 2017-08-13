package gossm

import (
	"io"
	"net/http"
)

func RunHttp(address string) {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		io.WriteString(rw, "Hello world!")
	})
	http.ListenAndServe(address, nil)
}
