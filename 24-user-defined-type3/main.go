package main

import (
	"fmt"
)

func main() {
	var u1 Unique

	u1.int = 101
	u1.string = "Anish"
	u1.bool = false

	fmt.Println(u1)

	u2 := Unique{1, "Anish", true}
	fmt.Println(u2)

	//name annonamouys structure
	a1 := struct {
		ID   int
		Name string
	}{
		ID:   101,
		Name: "Anish",
	}

	fmt.Println(a1)

	//another way unamed annonamouys structure

	a2 := struct {
		int
		string
	}{
		int:    101,
		string: "unnamed",
	}
	fmt.Println(a2)

}

// struct with not name but with type
type Unique struct {
	int
	string
	bool
}
