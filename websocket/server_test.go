package websocket

import (
	"testing"
)

func TestSecWebsocketAccept(t *testing.T) {
	key := "dGhlIHNhbXBsZSBub25jZQ=="
	expect := "s3pPLMBiTxaQ9kYGzzhZRbK+xOo="

	t.Log("Given the string of Sec-Websocket-Accept.")
	{
		result := SecWebsocketAccept(key)
		if result != expect {
			t.Fatalf("返回值不正确, 期望:%s, 结果:%s", expect, result)
		} else {
			t.Log("正确")
		}
	}
}