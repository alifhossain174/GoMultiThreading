package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

// Define a struct to represent a person
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	// Create an instance of Person
	person := Person{
		Name:    "Alice",
		Age:     30,
		Address: "123 Main St, Springfield",
	}

	// Use a WaitGroup to wait for the serialization goroutine to finish
	var wg sync.WaitGroup

	// Measure the total execution time
	start := time.Now()

	// Serialization
	wg.Add(1)
	var personJSON []byte
	go func() {
		defer wg.Done()
		var err error
		personJSON, err = json.Marshal(person)
		if err != nil {
			log.Fatalf("Error during serialization: %v", err)
		}
	}()

	// Wait for the serialization to complete
	wg.Wait() // Ensure serialization completes before proceeding

	// Now we can safely perform deserialization
	var decodedPerson Person
	err := json.Unmarshal(personJSON, &decodedPerson)
	if err != nil {
		log.Fatalf("Error during deserialization: %v", err)
	}

	// Total execution time
	elapsed := time.Since(start)

	// Output results
	fmt.Println("Serialized JSON:", string(personJSON))
	fmt.Println("Deserialized Person:", decodedPerson)
	fmt.Printf("Total Execution Time: %v\n", elapsed)
}
