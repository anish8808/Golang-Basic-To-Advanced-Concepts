package main

import "fmt"

// Promoted feilds -->
func main() {
	var p1 Person
	p1.ID = 100
	p1.Name = "Anish"
	p1.Email = "anisenownwe"
	p1.Contact = "23765464"       // this will be Person field not promoted field (so always take care while creating promoted field , that field should not be same)
	p1.Address.Contact = "134442" // this way we can access
	p1.Line1 = "1st lane"
	p1.City = "Gop"
	p1.State = "UP"
	p1.Street = "taramandal"
	p1.Contact = "India"
	p1.PIN = "837364"

	fmt.Println(p1)
}

type Person struct {
	ID                   int
	Name, Email, Contact string
	Address              //--> this is a permoted feild(directly give Address is promotred field) , if a field is promoted feild it can call the member of struct can directly we called
	//in this Example P1.city can be accessed directly beacue of promoted
}

type Address struct {
	Line1, City, Country, Street, State, PIN, Contact string
}
