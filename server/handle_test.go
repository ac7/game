package server

import (
	"testing"
)

func TestHandleConn(t *testing.T) {
	if HandleConn(nil, 0) == nil {
		t.Errorf("Attempt to handle nil connection did not result in error.")
	}

	// TODO more testing
}
