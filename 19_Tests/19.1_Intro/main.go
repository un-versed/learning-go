package main

import (
	"fmt"
	"tests-intro/addresses"
)

func main() {
	addressType := addresses.AddressType("Rua Antônio Costa Rosendo")

	fmt.Println(addressType)
}
