/*
* A struct (short for "structure") is a collection of data fields with declared data types.
* Golang has the ability to declare and create own data types by combining one or more types,
* including both built-in and user-defined types.
* Struct types are declared by composing a fixed set of unique fields.
**/
package main

import "fmt"

type user struct {
	name    string
	age     uint
	address address
}

type address struct {
	street string
	number int8
}

func main() {
	var user1 user
	user1.name = "Luke Skywalker"
	user1.age = 34
	user1.address.street = "Street Random"
	user1.address.number = 11
	fmt.Println(user1)

	var exampleAddress = address{"Starway", 22}
	var user2 = user{"Obi-wan Kenobi", 65, exampleAddress}
	fmt.Println(user2)

	// Implicit struct declaration
	user3 := user{"Anakin Skywalker", 50, address{"Star Gate", 112}}
	fmt.Println(user3)

	user4 := user{age: 20, address: address{street: "Whoops"}}
	fmt.Println(user4)

}
