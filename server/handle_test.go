package server

import (
	"bytes"
	"testing"

	"github.com/vmihailenco/msgpack"
)

type mockConn struct {
	t             *testing.T
	handshakeSent bool
}

func (mw *mockConn) Write(input []byte) (int, error) {
	t := mw.t
	if bytes.Equal(input, []byte(HANDSHAKE_SERVER)) {
		t.Logf("Valid handshake '%s' from server.", string(input))
	} else {
		t.Logf("Invalid handshake '%s' from server (expected '%s')", string(input), HANDSHAKE_SERVER)
	}
	return len(input), nil
}

func (mw *mockConn) Read(output []byte) (int, error) {
	if !mw.handshakeSent {
		mw.handshakeSent = true
		for i, b := range []byte(HANDSHAKE_CLIENT) {
			output[i] = b
		}
		return len(HANDSHAKE_CLIENT), nil
	} else {
		quitPacket, err := msgpack.Marshal(map[string]interface{}{
			"type": "quit",
		})
		for i, b := range quitPacket {
			output[i] = b
		}
		if err != nil {
			mw.t.Fatal(err)
		}

		return len(quitPacket), nil
	}
}

func (mw *mockConn) Close() error {
	return nil
}

func TestHandleConn(t *testing.T) {
	server := New()
	if server.HandleConn(nil, 0) == nil {
		t.Errorf("Attempt to handle nil connection did not result in error.")
	}

	mc := &mockConn{t, false}
	err := server.HandleConn(mc, 0)
	if err != nil {
		t.Fatal(err)
	}
}
