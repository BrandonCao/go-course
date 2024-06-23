package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	brandon := person{
		firstName: "Brandon",
		lastName:  "Cao",
		contact: contactInfo{
			email:   "brandon.cao@gametime.co",
			zipCode: 94085,
		},
	}

	brandon.print()
	brandon.updateName("Brandon2")
	brandon.print()

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName // * gives the value at that memory address. different than * at a type than a variable
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
