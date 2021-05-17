package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type persone struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := persone{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex.anderson@gmail.com",
			zipCode: 12345,
		},
	}

	pointer := &alex
	pointer.updateName("Alex222")
	alex.print()
}

func (pointerToPerson *persone) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p persone) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
