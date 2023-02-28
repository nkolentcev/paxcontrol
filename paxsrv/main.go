package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatalln(err)
	}

	if err := store.Setup(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8000", store)
	server.Run()
}
