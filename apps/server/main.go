package main

import (
	"log"
	"chess-game/server/cmd/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}