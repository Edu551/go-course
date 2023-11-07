package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type pets struct {
	name string
	race string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	pets
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Harper",
		contact: contactInfo{
			email:   "jim@email.com",
			zipCode: 123456,
		},
		pets: pets{
			name: "Logan",
			race: "Belgian Shepherd",
		},
	}

	jim.updateName("Sam")
	jim.print()
	fmt.Println("")
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
