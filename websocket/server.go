package websocket

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"net"
	"net/http"
)

type Socket struct {
	conn net.Conn
}


func SecWebsocketAccept(key string) string {
	const MAGIC_STRING = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	h := sha1.New()
	h.Write([]byte(key + MAGIC_STRING))
	value := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(value)
}

func Hijack(writer http.ResponseWriter) (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := writer.(http.Hijacker)
	if !ok {
		panic("It's a end.")
	}

	return hijacker.Hijack()
}

func GenResponseHeader(key string) []byte {
	acceptKey := SecWebsocketAccept(key)
	header := "HTTP/1.1 101 Switching Protocols\r\n" +
		"Upgrade: websocket\r\n" +
		"Connection: Upgrade\r\n" +
		"Sec-WebSocket-Accept:" + acceptKey + "\r\n"
	return []byte(header + "\r\n")
}
