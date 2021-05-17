package main

import "fmt"

type persone struct {
	firstName string
	lastName  string
}

func main() {
	alex := persone{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex)
}
