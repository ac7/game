package server

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

type mockWriter struct {
	t *testing.T
}

func (mw *mockWriter) Write(input []byte) (int, error) {
	t := mw.t
	if bytes.Equal(input, []byte(HANDSHAKE_SERVER)) {
		t.Logf("Valid handshake '%s' from server.", string(input))
	} else {
		t.Logf("Invalid handshake '%s' from server (expected '%s')", string(input), HANDSHAKE_SERVER)
	}
	return len(input), nil
}

func TestHandleConn(t *testing.T) {
	if HandleConn(nil, 0) == nil {
		t.Errorf("Attempt to handle nil connection did not result in error.")
	}

	reader := strings.NewReader(HANDSHAKE_CLIENT)
	writer := &mockWriter{t}
	err := HandleConn(bufio.NewReadWriter(bufio.NewReader(reader), bufio.NewWriter(writer)), 0)
	if err != nil {
		t.Fatal(err)
	}
}
