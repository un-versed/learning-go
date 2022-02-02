package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int32  `json:"age"`
}

func main() {
	dog := dog{
		Name:  "John",
		Breed: "Dachshund",
		Age:   4,
	}

	dogJSON, error := json.Marshal(dog)
	if error != nil {
		log.Fatal("Error parsing dogJSON")
	}

	fmt.Println(dogJSON)                  // JSON bytes
	fmt.Println(bytes.NewBuffer(dogJSON)) // JSON
}
