package websocket

import (
	"crypto/sha1"
	"encoding/base64"
)

type Socket struct {
}

func SecWebsocketAccept(key string) string {
	const MAGIC_STRING = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	h := sha1.New()
	h.Write([]byte(key + MAGIC_STRING))
	value := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(value)
}
