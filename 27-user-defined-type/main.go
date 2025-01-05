package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var num1 myInt

	var num2 int = 100

	num1 = myInt(num2) // both type is not same so er have to use type casting
	str := num1.ToString()
	fmt.Println("This is converted String:", str)
	fmt.Println("Value of num1:", num1, reflect.TypeOf(num1), "Value of the num2:", num2, reflect.TypeOf(num1))

	var mymap1 mymap
	mymap1 = make(mymap)
	mymap1["Anish"] = 25
	mymap1["Manish"] = 23
	if ok, err := mymap1.delete("Anish"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted ", ok)
	}
	fmt.Println("My map is :", mymap1, "Type of Map is :", reflect.TypeOf(mymap1))

	var map1 map[string]any
	map1 = mymap1 // refrence is here shallow copy so type and normal cn be same
	fmt.Println("My map is :", map1, "Type of Map is :", reflect.TypeOf(map1))

	//alias

	var newnum MyInt2 = 100
	fmt.Println("Type of newnu is ", reflect.TypeOf(newnum))

	var newInterface object = 100.0
	fmt.Println("Type of mewInterface is :", reflect.TypeOf(newInterface))

}

type myInt int //this is type ans not struct this is my own underlinne type(using this we can our method)

func (mi myInt) ToString() string {
	str := fmt.Sprint(mi)
	return str
}

type mymap map[string]any // but in this case the above we can use beacuse of refrence copy

func (mm mymap) delete(key string) (bool, error) {
	if mm == nil {
		return false, errors.New("nil Map")
	}

	if _, ok := mm[key]; !ok {
		return false, errors.New("No key is present")
	}

	delete(mm, key)
	return true, nil
}

type MyInt2 = int     // this is just an alias instead of int you can give to MyInt2
type object = any     //// this is just an alias instead of any you can give to object
type double = float64 //you can not create method with these alias names
type long = int64

//read or research in builtin standard package
//create user defined map and perform insert update get also delete method , each method should have error as last parameter

type newMap map[string]any

func (mymap newMap) insert(val int, key string) error {
	if mymap == nil {
		return errors.New("map is nil so no data is present")
	} else {
		mymap[key] = val
	}

	return nil
}

func (myMap newMap) delete(key string) error {
	if myMap == nil {
		return errors.New("map is nil so cant delete")
	}

	value, ok := myMap[key]
	if !ok {
		// Return an error if the key does not exist
		return errors.New("key not found")
	}

	// Delete the key from the map
	delete(myMap, key)

	// Optionally, print the deleted value or perform other actions
	fmt.Printf("Deleted key: %s, value: %s\n", key, value)
	return nil
}

func (myMap newMap) get(key string) (string, error) {
	value, ok := myMap[key]

	if !ok {
		return "", errors.New("The key is not present")
	}
	ans := value.(string)
	return ans, nil
}
