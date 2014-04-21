package server

import (
	"bytes"
	"testing"
)

type mockConn struct {
	t *testing.T
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
	for i, b := range []byte(HANDSHAKE_CLIENT) {
		output[i] = b
	}
	return len(HANDSHAKE_CLIENT), nil
}

func (mw *mockConn) Close() error {
	return nil
}

func TestHandleConn(t *testing.T) {
	server := NewServer()
	if server.HandleConn(nil, 0) == nil {
		t.Errorf("Attempt to handle nil connection did not result in error.")
	}

	mc := &mockConn{t}
	err := server.HandleConn(mc, 0)
	if err != nil {
		t.Fatal(err)
	}
}
