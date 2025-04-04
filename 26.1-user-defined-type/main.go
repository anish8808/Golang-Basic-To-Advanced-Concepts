package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var num1 myInt
	var num2 int = 100
	num1 = myInt(num2)

	fmt.Println("Value of num1", num1, "type of num1", reflect.TypeOf(num1))
	fmt.Println("Value of num2", num2, "type of num2", reflect.TypeOf(num2))

	var mymap1 myMap

	//mymap1 = make(myMap)
	mymap1["BANGALORE"] = 123544
	mymap1["HYDRABAND"] = 12345

	var map1 map[string]any

	map1 = mymap1

	fmt.Println(map1)
	fmt.Println("Value of num1", mymap1, "type of num1", reflect.TypeOf(mymap1))
	fmt.Println("Value of num2", map1, "type of num2", reflect.TypeOf(map1))

	str := num1.ToString()
	fmt.Println("value of str1", str, "type of str", reflect.TypeOf(str))

	if ok, err := mymap1.Delete("BANGALORE"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("deleted", ok)
	}

	if ok, err := myMap(map1).Delete("BANGALORE"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted", ok)
	}

}

type myInt int //-->
type myMap map[string]any

func (mi myInt) ToString() string {
	return fmt.Sprint(mi)
}

func (mm myMap) Delete(key string) (bool, error) {
	if mm == nil {
		return false, errors.New("nil Map")
	}
	if _, ok := mm[key]; !ok {
		return false, errors.New("nil Map")
	}
	delete(mm, key)
	return true, nil
}
