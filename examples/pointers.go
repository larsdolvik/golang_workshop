package examples

import "fmt"

// Pointers explained: https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back

func PointerExample() {
	a := 1
	b := &a

	fmt.Print(b)
	*b++
	fmt.Print(a)
}
