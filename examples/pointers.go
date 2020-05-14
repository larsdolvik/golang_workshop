package examples

import "fmt"

// Pointers explained: https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back

// PointerExample shows a basic pointer
func PointerExample() {
	a := 1  // Imagine that a represent a physical location in memory where we store 1: [1] - Address - ie: 0x1230ad2
	b := &a // Tell go that b now points to the memory location of a - ie: 0x1230ad2

	fmt.Printf("Address: %v\n ", b)
	*b++ // * in this context means to follow the address and retrieve the value stored in physical memory
	fmt.Printf("Value of a: %d\n", a)
}

// AddOne adds on to ints
func AddOne(ints *[]int, num int) {
	*ints = append(*ints, num)
}
