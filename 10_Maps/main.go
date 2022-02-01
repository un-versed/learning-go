package main

import "fmt"

func main() {
	// Map is a key-value structure to store data with the same type
	user := map[string]string{
		"name":  "Neo",
		"movie": "Matrix",
	}

	fmt.Println(user["name"])

	// A map can be a map :p
	user2 := map[string]map[string]string{
		"name": {
			"first": "Mr.",
			"last":  "Anderson",
		},
		"movie": {
			"name": "Matrix",
			"year": "1999",
		},
	}

	fmt.Println(user2["name"]["first"], user2["name"]["last"])
	fmt.Println(user2)
	fmt.Println("--- delete key ---")
	delete(user2, "name")
	fmt.Println(user2)

	// Create a map key
	user2["whatever"] = map[string]string{
		"whatever": "stuff",
	}
	fmt.Println(user2)
}
