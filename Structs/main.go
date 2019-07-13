package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

func main() {
	alex := person{
		firstName: "Partha",
		lastName:  "Bhattacharya",
		contactInfo: contactInfo{
			email:   "pb25193@gmail.com",
			zipcode: 213149,
		},
	}

	// alex.firstName = "Gabresh"
	// alex.lastName = "Pagnees"
	// alex.contact = contactInfo{email: "gabru.pagga@gmail.com", zipcode: 12312}
	alex.updateName("makichu")
	alex.print()
}

func (p person) print() {
	fmt.Println("Name is ", p.firstName, p.lastName)
	fmt.Println("contact info: ")
	fmt.Println("Email is ", p.contactInfo.email)
	fmt.Println("zipcode is ", p.contactInfo.zipcode)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
