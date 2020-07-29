package main

import (
	"github.com/vsivarajah/RiotStatistics/cmd"
	"log"
)

func main() {
	if message, err := cmd.Start(); err != nil {
		log.Printf("error starting server :: message: %s, err: %v", message, err)
	}
	log.Println("Starting application server...")

}
