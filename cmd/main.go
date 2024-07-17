package main

import (
	"log"
	"untitled-project/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
