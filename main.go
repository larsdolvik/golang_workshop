package main

import (
	"code/golang_workshop/examples"
	"fmt"
)

// Every go program needs a main package with a main function.
// This works as the applications entry point.

// Standard library documentation: https://golang.org/pkg/

// Useful packages:
// strings - https://golang.org/pkg/strings/
// sort - https://golang.org/pkg/sort/
// http - https://golang.org/pkg/net/http/

// Useful tools:
// https://gobyexample.com/
// https://transform.tools/json-to-go

func main() {
	// todo, err := examples.ShortHandRequest()
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }

	// fmt.Printf("userId: %v\n", todo.UserID)
	// fmt.Printf("completed: %v\n", todo.Completed)
	// fmt.Printf("title: %v\n", todo.Title)
	// fmt.Printf("id: %v\n", todo.ID)

	examples.PointerExample()

	ints := []int{1, 2, 3, 4}
	examples.AddOne(&ints, 5)

	for _, val := range ints {
		fmt.Printf("%d\n", val)
	}
}
