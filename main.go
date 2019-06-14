package main

import (
	// "io"
	"log"
	// "net/http"
	// "fmt"
	"github.com/luozan/opal/websocket"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	header := req.Header
	// 	log.Println(req)
	// 	log.Println(header)
	// 	io.WriteString(w, "xxxxxxxxxxxxxx")
	// })
	// http.ListenAndServe(":80", nil)
	log.Println(websocket.SecWebsocketAccept("dGhlIHNhbXBsZSBub25jZQ=="))
}
