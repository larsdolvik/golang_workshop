package main

import "fmt"

func main() {
	fmt.Printf("Test \n")

	// Variables
	var num int
	var float float32
	var str string
	var boolean bool

	fmt.Printf("%v \n", num)
	fmt.Printf("%v \n", float)
	fmt.Printf("%v \n", str)
	fmt.Printf("%v \n", boolean)

	// Short hand variable declaration
	standard := "Hello world"
	fmt.Printf("%s\n", standard)

	// Looping - while loop
	count := 0
	for {
		if count == 10 {
			fmt.Print("Exit")
			break
		}
		count++
	}

	// Looping - for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}

	arr := []int{10, 20, 30, 40, 50, 60}
	// Looping - arrays
	for index, val := range arr {
		fmt.Printf("%v\n%v\n", index, val)
	}

}
