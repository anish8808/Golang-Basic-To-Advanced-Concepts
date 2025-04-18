package main

import "fmt"

func main() {
	mobile := NewMob(&wiredCharger{})
	mobile2 := NewMob(&wirelessCharger{})

	mobile.checkcharging()
	mobile2.checkcharging()
}

//dependencies enjection

type IsCharger interface {
	charge()
}

// dependecies in mobile as mobile is depenedent on interface not concrete strcuture
type Mobile struct {
	isCharger IsCharger
}

// enject
func NewMob(c IsCharger) *Mobile {
	return &Mobile{isCharger: c}
}

// use charges
func (m *Mobile) checkcharging() {
	m.isCharger.charge()
}

// now we will create two type wiht ischarge interfface

type wiredCharger struct{}
type wirelessCharger struct{}

func (ch wiredCharger) charge() {
	fmt.Println("the charages is wired")
}

func (ch wirelessCharger) charge() {
	fmt.Println("The charger is wireless")
}
