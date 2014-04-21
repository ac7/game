package main

import (
	"fmt"
	"log"

	"github.com/ac7/game/server"
)

func main() {
	port := 1030
	location := fmt.Sprintf("localhost:%d", port)

	log.Printf("Starting server on localhost:%d\n", port)

	server := server.NewServer()
	err := server.Listen(location)
	if err != nil {
		log.Fatal(err)
	}
}
