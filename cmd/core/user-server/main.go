package main

import (
	"fmt"
	"hexagonal-structure/pkg/core/http/rest"
	"hexagonal-structure/pkg/core/storage/json"
	"hexagonal-structure/pkg/core/storage/memory"
	"hexagonal-structure/pkg/domain/adding"
	"hexagonal-structure/pkg/domain/deleting"
	"hexagonal-structure/pkg/domain/listing"
	"log"
	"net/http"
)

// Type StorageType defines available storage types
type Type int

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory
)

func main() {

	// set up storage
	storageType := JSON // this could be a flag; hardcoded here for simplicity

	var adder adding.Service
	var lister listing.Service
	var deleter deleting.Service

	switch storageType {
	case Memory:
		s := new(memory.Storage)

		adder = adding.NewService(s)
		lister = listing.NewService(s)
		deleter = deleting.NewService(s)

	case JSON:
		// error handling omitted for simplicity
		s, _ := json.NewStorage()

		adder = adding.NewService(s)
		lister = listing.NewService(s)
		deleter = deleting.NewService(s)
	}

	// set up the HTTP server
	router := rest.Handler(adder, lister, deleter)

	fmt.Println("The user server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
