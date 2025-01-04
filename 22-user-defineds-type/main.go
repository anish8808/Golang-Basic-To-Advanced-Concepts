package main

import "fmt"

func main() {
	var p1 Person
	var a1 Address

	a1.Line1 = "1st Lane"
	a1.Street = "NGR LAYOUT"
	a1.City = "Bangalore"
	a1.State = "Karatataka"
	a1.Country = "India"
	a1.PIN = "560068"
	p1.ID = 100
	p1.Name = "Anish"
	p1.Address = a1
	p1.Email = "anishkashyap11111@gmail.com"
	p1.Contact = "9305992260"

	fmt.Println(p1)
	p2 := Person{
		ID:      101,
		Name:    "Anish",
		Address: Address{Line1: "Bommanhalli", Street: "NGR", City: "Bangalore", State: "Karanataka", Country: "India", PIN: "560068"},
		Email:   "anishasy111",
		Contact: "9305992260"} //sort hand decleration
	fmt.Println(p2)
}

//1- struct
//2- user defined type from any existing type
//3- methods are called on types , Methods can be  created on types

type Person struct {
	ID      int
	Name    string
	Address Address //can give same name as a type feild
	Email   string
	Contact string
}
type Address struct {
	Line1, Street, City, State, Country, PIN string
}
