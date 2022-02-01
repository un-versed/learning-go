package main

import (
	"fmt"
)

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("What")
	generic("This is js?")
	generic(666)
	generic(false)
}
