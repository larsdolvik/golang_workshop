package main

import (
	"code/golang_workshop/examples"
)

// Every go program needs a main package with a main function.
// This works as the applications entry point.

// Standard library documentation: https://golang.org/pkg/

// Useful packages:
// strings - https://golang.org/pkg/strings/
// sort - https://golang.org/pkg/sort/
// http - https://golang.org/pkg/net/http/

// Useful examples:
// https://gobyexample.com/

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

	examples.Loop()
}
