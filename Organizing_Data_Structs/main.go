package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	alex := person{firstname: "Alex", lastname: "Anderson"}

	var benex person

	benex.firstname = "Benex"
	benex.lastname = "Drake"

	fmt.Println(alex)
	fmt.Println(benex)
}
