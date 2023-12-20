package main

import (
	"fmt"
	"hexagonal-structure/pkg/core/storage/json"
	"hexagonal-structure/pkg/domain/adding"
)

func main() {

	var adder adding.Service

	// error handling omitted for simplicity
	s, _ := json.NewStorage()

	adder = adding.NewService(s)

	// add some sample data
	adder.AddSampleUsers(DefaultUsers)

	fmt.Println("Finished adding sample data.")
}
