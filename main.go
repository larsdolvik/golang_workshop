package main

import (
	"modules/examples"
)

// Every go program needs a main package with a main function.
// This works as the applications entry point.

// Useful packages:
// strings - https://golang.org/pkg/strings/
// http - https://golang.org/pkg/net/http/

// Useful tools:
// https://gobyexample.com/
// https://transform.tools/json-to-go

func main() {

	// Initate new router
	// router := httprouter.New()

	// Set up your different API routes
	// router.GET("/helloworld", hello.Hello)
	// router.GET("/hello/:name", hello.SayHello)
	// The GetData func must med assigned a route here

	// Example of an easy cors setup

	// handler := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000", ""},
	// 	AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
	// 	AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	// 	AllowCredentials: true,
	// }).Handler(router)

	// log.Println("Starting api...")
	// http.ListenAndServe(":8080", handler)

	client := examples.NewClient()
	client.GetPageInfo()

	// ints := []int{1, 2, 3, 4}
	// examples.AddOne(&ints, 5)

	// for _, val := range ints {
	// 	fmt.Printf("%d\n", val)
	// }
}
