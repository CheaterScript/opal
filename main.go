package main

import (
	"fmt"
	"github.com/luozan/opal/websocket"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		header := req.Header
		fmt.Println(header)
		if _, ok := header["Sec-Websocket-Key"]; !ok {
			io.WriteString(w, "xxxxxxxxxxxxxx")
			w.WriteHeader(200)
			return
		}
		key := websocket.SecWebsocketAccept(header["Sec-Websocket-Key"][0])
		w.Header().Set("Upgrade", "websocket")
		w.Header().Set("Connection", "Upgrade")
		w.Header().Set("Sec-WebSocket-Accept", key)
		w.WriteHeader(101)
	})
	http.ListenAndServe(":80", nil)
}
