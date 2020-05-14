package examples

import "fmt"

// Initialized variables in golang get initialized with a "zero-value"
// See more: https://tour.golang.org/basics/12
// Zero value is 0 for numeric types, false for boolean types and "" for empty strings.

var num int
var float float32
var str string
var boolean bool

const neverChanging = "uniqueId-123"

// PrintVars prints variables and is publicly exported
func PrintVars() {
	fmt.Printf("%v \n", num)
	fmt.Printf("%v \n", float)
	fmt.Printf("%v \n", str)
	fmt.Printf("%v \n", boolean)
	fmt.Printf("%v \n", neverChanging)

	// Short hand variable declaration
	standard := "Hello world"
	fmt.Printf("%s \n", standard)
}

// this function is not publicly exported, because it's function name is not
// capitalized.
func printVars() {
	fmt.Print("I'm private")
}
