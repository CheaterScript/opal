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
		tcpConn, rBuf, err := websocket.Hijack(w)
		fmt.Println(tcpConn, rBuf, err)
		responseHeader := websocket.GenResponseHeader(header["Sec-Websocket-Key"][0])
		fmt.Println(tcpConn, responseHeader)
		tcpConn.Write(responseHeader);
		socket := websocket.New(tcpConn)
		go socket.Recv()
	})

	http.ListenAndServe(":80", nil)
}

func Recv(socket websocket.Sokcet) {

}