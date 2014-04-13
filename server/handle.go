package server

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

const HANDSHAKE_CLIENT = "handshake"
const HANDSHAKE_SERVER = "handshake_part_two"

func HandleConn(conn io.ReadWriter, id int) error {
	if conn == nil {
		return fmt.Errorf("Invalid connection!")
	}

	// first we send the HANDSHAKE_SERVER string to the client
	var helloMessage []byte = make([]byte, 9)
	_, err := conn.Read(helloMessage)
	if err != nil {
		return fmt.Errorf("Could not read from connection: %s", err.Error())
	}
	if !bytes.Equal(helloMessage, []byte(HANDSHAKE_CLIENT)) {
		return fmt.Errorf(`Recieved non-protocol handshake "%s" (expected "%s"), abort.`,
			string(helloMessage), HANDSHAKE_CLIENT)
	}
	log.Printf("Connection %d sent correct handshake \"%s\"\n", id, HANDSHAKE_CLIENT)

	// then we check to make sure the client returns the HANDSHAKE_CLIENT string
	n, err := conn.Write([]byte(HANDSHAKE_SERVER))
	if err != nil {
		return fmt.Errorf(`Cannot write to connection %d: "%s"`, id, err.Error())
	} else if n != len(HANDSHAKE_SERVER) {
		return fmt.Errorf("Only wrote %d bytes to connection (expected %d).", n, len(HANDSHAKE_SERVER))
	} else {
		log.Printf(`Sent return handshake "%s" to connection number %d`, HANDSHAKE_SERVER, id)
	}

	return nil
}
