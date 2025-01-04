package main

import "fmt"

// methods in go by using this we are achiving ploymorphosim similar
func main() {
	p1 := Person{
		ID:      101,
		Name:    "Anish",
		Email:   "anishkashyap11111@gmail.com",
		Contact: "9387373",
	}

	Display(p1)  // this is normal function
	p1.Display() // method can only call with the strture
	p1.toString()

	a1 := Address{
		City:  "GOP",
		State: "UP",
	}
	a1.Display()

	var T1 T1
	T1.dispaly()

	var t2 T2
	t2.dispaly()

}

type Person struct {
	ID                   int
	Name, Email, Contact string
}

func Display(p Person) { //this is normal function cannot call with a strcture
	fmt.Println("ID:", p.ID)
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Contact:", p.Contact)

}

func (p Person) Display() { ////This is method beacuse methods contains recivers here (p Person) is reciver , front of the function is reciver (p Person)
	fmt.Println("ID:", p.ID) //reciver can be only one
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Contact:", p.Contact)
}

func (p Person) toString() string {
	return fmt.Sprint(p)
}

type Address struct {
	City, State string
}

// this is the same method having different type (but with normail function this can be happed like 2 time dispaly function)
func (a Address) Display() {
	fmt.Println("City", a.City)
	fmt.Println("State", a.State)
}

type T1 struct {
}

func (e T1) dispaly() {
	fmt.Println("Hello World from T1")
}

type T2 struct {
}

func (e T2) dispaly() {
	fmt.Println("Hello World from T2")
}
