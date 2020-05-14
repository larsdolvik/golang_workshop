package main

import (
	"log"
	"modules/packages/hello"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {

	// Initate new router
	router := httprouter.New()

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
}
