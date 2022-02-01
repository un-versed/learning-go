package main

import "fmt"

type user struct {
	name string
	age  int32
}

func (u user) save() {
	fmt.Printf("Saving user %s in database...\n", u.name)
}

func (u user) isUnderAge() bool {
	return u.age <= 18
}

func (u *user) makeBirthday() {
	u.age++
}

func main() {
	user1 := user{name: "Roger", age: 14}
	user1.save()

	user2 := user{name: "Marcus", age: 30}
	user2.save()

	fmt.Printf("The user %s under age is %t\n", user1.name, user1.isUnderAge())
	fmt.Printf("The user %s under age is %t\n", user2.name, user2.isUnderAge())

	user1.makeBirthday()
	fmt.Printf("The user %s age is %d years", user1.name, user1.age)
}
