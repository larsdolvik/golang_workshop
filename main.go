package main

import (
	"log"
	"modules/packages/challenge"
	"modules/packages/hello"
	"net/http"

	// Every go program needs a main package with a main function.
	// This works as the applications entry point.

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

// Every go program needs a main package with a main function.
// This works as the applications entry point.

// Standard library documentation: https://golang.org/pkg/
// Useful packages:
// strings - https://golang.org/pkg/strings/
// http - https://golang.org/pkg/net/http/

// Useful tools:
// https://gobyexample.com/
// https://transform.tools/json-to-go

func main() {

	// Initate new router
	router := httprouter.New()

	// Set up your different API routes
	router.GET("/helloworld", hello.Hello)
	router.GET("/hello/:name", hello.SayHello)
	router.GET("/solution", challenge.Solution)
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
}
