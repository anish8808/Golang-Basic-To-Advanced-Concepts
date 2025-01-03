package main

import "fmt"

// slice , map , channel --> make to instantiate it
// map is a refrence type
// key of a map can be any thing as and where == operator
// can check nil , if map is not instatiated the map will be nil
// everything of map is also refrence type
func main() {
	var mymap map[string]any
	if mymap == nil {
		fmt.Println("mymap is nil")
	}

	mymap = make(map[string]any)

	mymap["banglore-1"] = 123
	mymap["hydrabad-1"] = 789
	mymap["Gorakhpur-1"] = 101112
	mymap["bangalore-2"] = 456

	//print the map
	for key, val := range mymap {
		fmt.Println("Key:", key, "Value:", val)
	}

	val, ok := mymap["bangalore-2"] // ok is a bool tells that key is present in the map or not
	if ok {
		fmt.Println("The key is present in the map", val)
	} else {
		fmt.Println("Key is not present in the map")
	}

	val, ok = mymap["gorakhour-2"] // ok is a bool tells that key is present in the map or not
	if ok {
		fmt.Println("The key is present in the map", val)
	} else {
		fmt.Println("Key is not present in the map")
	}

	//delete a map with a key without return status
	delete(mymap, "bangalore-1")

	val, ok = mymap["bangalore-1"] // ok is a bool tells that key is present in the map or not
	if ok {
		fmt.Println("The key is present in the map not deleted")
	} else {
		fmt.Println("Deleted from map")
	}

	//copy function
	//perform deep copy of a map
	mymap1 := make(map[string]any)
	//mymap1 = mymap by this way only shallow copy will happen
	//for deepcopy we have to run the loop
	for key, val := range mymap {
		mymap1[key] = val
	}
	delete(mymap, "hydrabad-1")
	for key, val := range mymap1 {
		fmt.Println("mympa1", "Key:", key, "value:", val)
	}
	for key, val := range mymap {
		fmt.Println("mympa", "Key:", key, "value:", val)
	}
}
