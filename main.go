package main

import (
	"log"
	"modules/packages/hello"
	"net/http"
	"code/golang_workshop/examples"
	"fmt"
)

// Every go program needs a main package with a main function.
// This works as the applications entry point.

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {

	// Initate new router
	router := httprouter.New()
	// Useful tools:
	// https://gobyexample.com/
	// https://transform.tools/json-to-go

	// Set up your different API routes
	router.GET("/helloworld", hello.Hello)
	router.GET("/hello/:name", hello.SayHello)
	// The GetData func must med assigned a route here

	// Example of an easy cors setup
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", ""},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	log.Println("Starting api...")
	http.ListenAndServe(":8080", handler)


	examples.PointerExample()

	ints := []int{1, 2, 3, 4}
	examples.AddOne(&ints, 5)

	for _, val := range ints {
		fmt.Printf("%d\n", val)
	}
}
