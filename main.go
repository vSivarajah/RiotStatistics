package main

import (
	"github.com/vsivarajah/RiotStatistics/cmd"
	"log"
)

func main() {
	if err := cmd.Start(); err != nil {
		log.Printf("error starting server %v", err)
	}
	log.Println("Starting application server...")

}
