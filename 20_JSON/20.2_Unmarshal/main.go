package main

import (
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
	dogJSON := `{"name":"John","breed":"Dachshund","age":4}`

	var dog dog

	if error := json.Unmarshal([]byte(dogJSON), &dog); error != nil {
		log.Fatal(error)
	}

	fmt.Println(dog)

	dogJSON2 := `{"name":"John","breed":"Dachshund"}`
	dog2 := make(map[string]string)

	if error := json.Unmarshal([]byte(dogJSON2), &dog2); error != nil {
		log.Fatal(error)
	}

	fmt.Println(dog2)
}
