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
	anton := persone{
		firstName: "Anton",
		lastName:  "Lytvynov",
		contactInfo: contactInfo{
			email:   "anton.lytvynov@gmail.com",
			zipCode: 49000,
		},
	}

	anton.updateName("Anton New Value")
	anton.print()
}

func (pointerToPerson *persone) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p persone) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
