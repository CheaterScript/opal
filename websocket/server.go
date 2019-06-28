package websocket

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
)

const (
	MAX_READ_BUFFER_SIZE = 8192
)

type Socket interface {
	Recv()
	Send() error
	Ping() (int, error)
}

type Websocket struct {
	conn net.Conn
}

func New(conn net.Conn) *Websocket {
	return &Websocket{conn }
}

func (socket *Websocket) Recv() {
	buf := make([]byte, MAX_READ_BUFFER_SIZE)
	for {
		n, err := socket.conn.Read(buf)
		if err == nil {
			fmt.Println("It's a data.")
			fmt.Println(buf, n)
			socket.Send(buf[:n])
		}else{
			fmt.Println(err)
		}
	}
}

func (socket *Websocket) Ping() (int, error) {
	conn := socket.conn
	return conn.Write([]byte{})
}

func (socket *Websocket) Send(data []byte) (int, error){
	return socket.conn.Write(data)
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
