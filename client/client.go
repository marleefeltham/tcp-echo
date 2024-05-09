package main

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
)

func main() {
	// connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Error().Err(err).Msg("failed to dial")
		return
	}
	defer conn.Close()

	// write data to the server
	_, err = conn.Write([]byte("hello\n"))
	if err != nil {
		log.Error().Err(err).Msg("failed to write to server")
		return
	}

	// read response from the server
	response := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Error().Err(err).Msg("failed to read")
			return
		}
		response = append(response, buffer[:n]...)
		if buffer[n-1] == '\n' {
			break
		}
	}

	fmt.Printf("server response: %s", string(response))
}
