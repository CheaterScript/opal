package websocket

import (
	"bufio"
	"net"
	"net/http"
	"testing"
)

type fakeHijacker struct {
}

func (hijacker *fakeHijacker) Header() http.Header{
	return nil
}
func (hijacker *fakeHijacker) Write([]byte) (int, error){
	return 0, nil
}
func (hijacker *fakeHijacker) WriteHeader(statusCode int) {
	return
}
func (hijacker *fakeHijacker) Hijack()(net.Conn, *bufio.ReadWriter, error){
	return  nil,nil,nil
}

func TestSecWebsocketAccept(t *testing.T) {
	key := "dGhlIHNhbXBsZSBub25jZQ=="
	expect := "s3pPLMBiTxaQ9kYGzzhZRbK+xOo="

	t.Log("Given the string of Sec-Websocket-Accept.")
	{
		result := SecWebsocketAccept(key)
		if result != expect {
			t.Fatalf("Wrong result, expectation:%s, result:%s", expect, result)
		}
	}
}

func TestHijack(t *testing.T) {
	t.Log("Hijacked the connection of HTTP.")
	{
		hijacker := &fakeHijacker{}
		_, _, err := Hijack(hijacker)
		if err != nil {
			t.Fatalf("unknown error.")
		}
	}
}

func TestGenResponseHeader(t *testing.T) {
	key := "dGhlIHNhbXBsZSBub25jZQ=="
	t.Log("Given the string of response header.")
	{
		header := GenResponseHeader(key)
		if len(header) <= 0{
			t.Fatal("The length of header less than zero.")
		}
	}
}