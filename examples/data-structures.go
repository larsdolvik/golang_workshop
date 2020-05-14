package examples

import "fmt"

// Datastructures contains datastructure examples
func Datastructures() {
	// Map
	basicMap := map[string]string{}

	basicMap["hello"] = "world"
	basicMap["friend"] = "best"

	for key, val := range basicMap {
		fmt.Printf("Key: %v\n Val: %v\n", key, val)
	}

	// Struct
	type person struct {
		Firstname string
		Lastname  string
	}

	personOne := person{
		Firstname: "Ole",
		Lastname:  "Haugen",
	}

	personTwo := person{
		Firstname: "Kjell",
		Lastname:  "Trosten",
	}

	people := []person{personOne, personTwo}

	for index, val := range people {
		fmt.Printf("Index: %v\n Firstname: %v\n Lastname: %v\n", index, val.Firstname, val.Lastname)
	}

	type customInt int

	var myCustomVar customInt
	fmt.Printf("%v", myCustomVar)
}
