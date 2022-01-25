package main

import "fmt"

type person struct {
	name  string
	age   uint8
	email string
}

type student struct {
	person
	course     string
	university string
}

func main() {
	var person1 = person{"Neo", 28, "neo@matrix.com"}
	fmt.Println(person1)

	var student1 = student{person1, "Computer Engineering", "MIT"}
	fmt.Println(student1)

	fmt.Println(student1.name)
}
