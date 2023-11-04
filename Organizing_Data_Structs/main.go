package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstname: "Alex",
		lastname:  "Anderson",
		contact: contactInfo{
			email:   "alex@gmail.de",
			zipCode: "12345",
		},
	}

	var benex person

	benex.firstname = "Benex"
	benex.lastname = "Drake"
	benex.contact.email = "sagichnicht@nope.de"
	benex.contact.zipCode = "40000"

	benex.print()
	fmt.Println()
	alex.print()

	alex.updateName("Wilhelm")

	fmt.Println()
	alex.print()

}

func (p *person) updateName(newFirstName string) {
	(*p).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
